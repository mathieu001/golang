package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	rootPath   string    = "E:/Go/Development/src"
	nodirs     [2]string = [...]string{"/github.com", "/goplayer"}
	suffixname string    = ".go"
)

var (
	linesum int
	mutex   *sync.Mutex = new(sync.Mutex)
)

func main() {
	argsLen := len(os.Args)
	if argsLen == 2 {
		rootPath = os.Args[1]
	} else if argsLen == 3 {
		rootPath = os.Args[1]
		suffixname = os.Args[2]
	}
	done := make(chan bool)
	fmt.Println("rootPath=", rootPath)
	go codeLineSum(rootPath, done)
	fmt.Println("Analysing...")
	<-done
	fmt.Println("total line:", linesum)
}

func codeLineSum(root string, done chan bool) {
	var goes int
	// fmt.Println("codeLineSum root=", root)
	godone := make(chan bool)
	isDstDir := checkDir(root)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("root: %s, panic: %#v\n", root, pan)
		}

		for i := 0; i < goes; i++ {
			<-godone
		}
		done <- true
	}()

	if !isDstDir {
		fmt.Printf("%s is not DstDir, skip\n", root)
		return
	}
	rootfi, err := os.Lstat(root)
	checkerr(err)

	//open root dir
	rootdir, err := os.Open(root)
	checkerr(err)
	defer rootdir.Close()

	if rootfi.IsDir() {
		//parse the dir
		fis, err := rootdir.Readdir(0)
		checkerr(err)
		for _, fi := range fis {
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			// fmt.Println("fi.name=", fi.Name())
			//create one goroutine for each file
			goes++
			// fmt.Println("goes=", goes)
			if fi.IsDir() {
				// go codeLineSum(root+"/"+fi.Name(), done)
				go codeLineSum(root+"/"+fi.Name(), godone)
			} else {
				go readfile(root+"/"+fi.Name(), godone)
			}
		}
	} else {
		//rootfi is a file, current goroutine has only one child
		goes = 1
		go readfile(root, godone)
	}

}

func readfile(filename string, done chan bool) {
	var line int
	isDstFile := strings.HasSuffix(filename, suffixname)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("filename: %s, panic: %#v\n", filename, pan)
		}

		//if it's the target file, add linesum
		if isDstFile {
			addLineNum(line)
			fmt.Printf("file %s complete, line=%d\n", filename, line)
		}

		done <- true
	}()

	if !isDstFile {
		return
	}

	file, err := os.Open(filename)
	checkerr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			// fmt.Println("breaking------")
			break
		}

		if !isPrefix {
			line++
		}
	}

}

func addLineNum(num int) {
	mutex.Lock()
	defer mutex.Unlock()
	linesum += num
}

func checkDir(dirPath string) bool {
	for _, dir := range nodirs {
		// fmt.Println("dirPath=", dirPath)
		// fmt.Println("checkDir=", rootPath+dir)
		if rootPath+dir == dirPath {
			return false
		}
	}
	return true
}

func checkerr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
