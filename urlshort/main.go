package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/urlshort"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yaml := `
		    - path: /urlshort
		      url: https://github.com/gophercises/urlshort
		    - path: /urlshort-final
		      url: https://github.com/gophercises/urlshort/tree/solution
		    `

	//jsonByte := `{"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort", "/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2"}`

	yamlFlagPtr := flag.String("f", "test.yaml", "specify the yaml file name")
	flag.Parse()
	fmt.Println(*yamlFlagPtr)

	yamlByte := readYAMLFile(*yamlFlagPtr)

	//yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	yamlHandler, err := urlshort.YAMLHandler(yamlByte, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func readYAMLFile(f string) []byte {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return dat

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/abc", abc)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func abc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "abc!")
}
