package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type pathToUrl struct {
	Path string `json:"path" yaml:"path"`
	Url  string `json:"url" yaml:"url"`
}

func createHandlerFromYaml(mux *http.ServeMux) {
	var yamlPaths = []pathToUrl{}
	yamlFile, err := os.ReadFile("paths.yaml")
	if err != nil {
		fmt.Println("can't read yaml file", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &yamlPaths)
	if err != nil {
		fmt.Println("can't Unmarshal yaml", err)
		return
	}
	for _, path := range yamlPaths {
		mux.HandleFunc(path.Path, mappedToHandler(path.Url))
	}
}

func createHandlerFromJson(mux *http.ServeMux) http.Handler {
	var jsonPaths = []pathToUrl{}
	jsonFile, err := os.ReadFile("paths.json")
	if err != nil {
		fmt.Println("can't read json file", err)
		return mux
	}
	err = yaml.Unmarshal(jsonFile, &jsonPaths)
	if err != nil {
		fmt.Println("can't Unmarshal yaml", err)
		return mux
	}
	for _, path := range jsonPaths {
		mux.HandleFunc(path.Path, mappedToHandler(path.Url))
	}
	return mux
}

func mappedToHandler(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
