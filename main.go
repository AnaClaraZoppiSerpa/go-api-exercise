package main

import (
	//"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
)

type MaintainerInfo struct {
	name  string `yaml:"name"`
	email string `yaml:"email"`
}

type AppMetadata struct {
	title       string           `yaml:"title"`
	version     string           `yaml:"version"`
	maintainers []MaintainerInfo `yaml:"maintainer"`
	company     string           `yaml:"company"`
	website     string           `yaml:"website"`
	source      string           `yaml:"source"`
	license     string           `yaml:"license"`
	description string           `yaml:"description"`
}

var appInfos = []AppMetadata{}

func handleGetAppMetadata(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that we are returning YAML data
	w.Header().Set("Content-Type", "application/x-yaml")

	// Marshal the appMetadataArray to YAML
	yamlData, err := yaml.Marshal(appInfos)
	if err != nil {
		log.Printf("Error marshaling YAML: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the YAML data to the response
	_, err = w.Write(yamlData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}


func fillWithDummyApps() {
	dummyMaintainer := MaintainerInfo{
		name:  "ana",
		email: "anaclara.zoppiserpa@gmail.com",
	}

	dummyApp_1 := AppMetadata{
		title:       "my app",
		version:     "1.0",
		maintainers: []MaintainerInfo{dummyMaintainer},
		company:     "ana's company",
		website:     "ana's github",
		source:      "source",
		license:     "license",
		description: "this is my app",
	}

	dummyApp_2 := AppMetadata{
		title:       "my app",
		version:     "1.0",
		maintainers: []MaintainerInfo{dummyMaintainer},
		company:     "ana's company",
		website:     "ana's github",
		source:      "source",
		license:     "license",
		description: "this is my app",
	}

	appInfos = append(appInfos, dummyApp_1, dummyApp_2)

	fmt.Println(appInfos)
}

func main() {
	fmt.Println("main")
	fmt.Println("dummy apps")
	fillWithDummyApps()

	http.HandleFunc("/metadata", handleGetAppMetadata)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// define yaml struct of metadata
// define endpoints

// check if yaml is working properly on the requests
// create new app metadata
// return all the metadatas
// edit a metadata
// delete a metadata
// more sophisticated queries
