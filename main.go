package main

import (
	//"encoding/json"

	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v3"
)

type StringFilter struct {
	Fragment string `yaml:"fragment"`
	Matcher  string `yaml:"matcher"` // contains, equals, startswith, nequals, ncontains
}

type MaintainerFilter struct {
	NameFilter  StringFilter `yaml:"name_filter"`
	EmailFilter StringFilter `yaml:"email_filter"`
}

type VersionFilter struct {
	Lower        int    `yaml:"lower"`
	Upper        int    `yaml:"upper"`
	RegexPattern string `yaml:"regex_pattern"`
}

type MaintainerInfo struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type AppMetadata struct {
	Title       string           `yaml:"title"`
	Version     string           `yaml:"version"`
	Maintainers []MaintainerInfo `yaml:"maintainer"`
	Company     string           `yaml:"company"`
	Website     string           `yaml:"website"`
	Source      string           `yaml:"source"`
	License     string           `yaml:"license"`
	Description string           `yaml:"description"`
}

var appInfos = []AppMetadata{}

func createSingle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var newMetadata AppMetadata
	err = yaml.Unmarshal(body, &newMetadata)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	appInfos = append(appInfos, newMetadata)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New metadata added successfully"))
}

func getAllNoQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")

	yamlData, err := yaml.Marshal(&appInfos)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(yamlData)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func singleQuery(w http.ResponseWriter, r *http.Request) {
}

func multipleQueries(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/all", getAllNoQuery)
	http.HandleFunc("/add", createSingle)
	http.HandleFunc("/query/single", createSingle)
	http.HandleFunc("/query/multiple", createSingle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
