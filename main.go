package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v3"
)

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

func dummyRetrieve(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that we are returning YAML data
	w.Header().Set("Content-Type", "application/yaml")
	// Marshal the appMetadataArray to YAML
	yamlData, err := yaml.Marshal(&appInfos)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Write the YAML data to the response
	_, err = w.Write(yamlData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func dummyPersist(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	// Parse the YAML request payload
	var newMetadata AppMetadata
	err = yaml.Unmarshal(body, &newMetadata)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Here, you can perform additional validation or business logic, if required.
	// For example, check if the ISBN is unique, etc.

	// Save the book to the database or perform any other necessary actions.
	// For simplicity, we'll just print the book information for now.
	fmt.Printf("New Book: %+v\n", newMetadata)
	appInfos = append(appInfos, newMetadata)

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book created successfully"))
}

func fillWithDummyApps() {
	dummyMaintainer := MaintainerInfo{
		Name:  "ana",
		Email: "anaclara.zoppiserpa@gmail.com",
	}

	dummyApp_1 := AppMetadata{
		Title:       "my app",
		Version:     "1.0",
		Maintainers: []MaintainerInfo{dummyMaintainer},
		Company:     "ana's company",
		Website:     "ana's github",
		Source:      "source",
		License:     "license",
		Description: "this is my app",
	}

	appInfos = append(appInfos, dummyApp_1, dummyApp_1)

	fmt.Println(appInfos)
}

func main() {
	//fillWithDummyApps()

	http.HandleFunc("/metadata", dummyRetrieve)
	http.HandleFunc("/new", dummyPersist)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
