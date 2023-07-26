package main

import (
	"fmt"
	"strings"
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
	Lower        string    `yaml:"lower"`
	Upper        string    `yaml:"upper"`
}

type Query struct {
	Title StringFilter
	Version VersionFilter
	Maintainers MaintainerFilter
	Company StringFilter
	Website StringFilter
	Source StringFilter
	License StringFilter
	Description StringFilter
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

var metadataArray []AppMetadata

func addSingleElement(newElement AppMetadata) {
	// validations
	metadataArray = append(metadataArray, newElement)
}

func addMultipleElements(newElements []AppMetadata) {
	// validations
	metadataArray = append(metadataArray, newElements...)
}

func getAll() []AppMetadata {
	return metadataArray
}

func checkStringFilter(field string, filter StringFilter) bool {
	if filter.Matcher == "contains" {
		return strings.Contains(field, filter.Fragment)
	}
	if filter.Matcher == "equals" {
		return field == filter.Fragment
	}
	if filter.Matcher == "startswith" {
		return strings.HasPrefix(field, filter.Fragment)
	}
	if filter.Matcher == "ncontains" {
		return !strings.Contains(field, filter.Fragment)
	}
	if filter.Matcher == "nequals" {
		return field != filter.Fragment
	}
	return true
}

func checkMaintainerFilter(field []MaintainerInfo, filter MaintainerFilter) bool {
	return true
}

func checkVersionFilter(field string, filter VersionFilter) bool {
	return true
}

func (filter StringFilter) empty() bool {
	return filter.Fragment == "" && filter.Matcher == ""
}

func (filter VersionFilter) empty() bool {
	return filter.Lower == "" && filter.Upper == ""
}

func (filter MaintainerFilter) empty() bool {
	return filter.NameFilter.empty() && filter.EmailFilter.empty()
}

func satisfiesQuery(q Query, app AppMetadata) bool {
	if !q.Title.empty() && !checkStringFilter(app.Title, q.Title) { return false }
	if !q.Company.empty() && !checkStringFilter(app.Company, q.Company) { return false }
	if !q.Website.empty() && !checkStringFilter(app.Website, q.Website) { return false }
	if !q.Source.empty() && !checkStringFilter(app.Source, q.Source) { return false }
	if !q.Description.empty() && !checkStringFilter(app.Description, q.Description) { return false }

	if !q.Version.empty() && !checkVersionFilter(app.Version, q.Version) { return false }
	if !q.Maintainers.empty() && !checkMaintainerFilter(app.Maintainers, q.Maintainers) { return false }

	return true
}

func fillArrayWithDummyApps() {
	app1 := AppMetadata{
		Title: "My Awesome App",
		Version: "1.0.0",
		Maintainers: []MaintainerInfo{
			{Name: "John Doe", Email: "john.doe@example.com"},
			{Name: "Jane Smith", Email: "jane.smith@example.com"},
		},
		Company:     "XYZ Inc.",
		Website:     "https://www.example.com",
		Source:      "https://github.com/example/myapp",
		License:     "MIT",
		Description: "This is a dummy app for testing purposes.",
	}

	app2 := AppMetadata{
		Title: "GameMaster 9000",
		Version: "1.8.3",
		Maintainers: []MaintainerInfo{
			{Name: "Sarah Williams", Email: "sarah.williams@example.com"},
			{Name: "James Anderson", Email: "james.anderson@example.com"},
		},
		Company:     "GameTech Co.",
		Website:     "https://www.gamemaster9000.com",
		Source:      "https://github.com/gametech/gamemaster",
		License:     "MIT",
		Description: "The ultimate game development toolkit.",
	}

	app3 := AppMetadata{
		Title: "Data Cruncher Pro",
		Version: "3.2.0",
		Maintainers: []MaintainerInfo{
			{Name: "David Lee", Email: "david.lee@example.com"},
		},
		Company:     "Data Cruncher Inc.",
		Website:     "https://www.datacruncherpro.com",
		Source:      "https://github.com/datacruncherpro/source",
		License:     "GPLv3",
		Description: "Your go-to tool for data analysis and visualization.",
	}

	app4 := AppMetadata{
		Title: "Cool App 2000",
		Version: "2.5.1",
		Maintainers: []MaintainerInfo{
			{Name: "Michael Johnson", Email: "michael.johnson@example.com"},
			{Name: "Emily Brown", Email: "emily.brown@example.com"},
		},
		Company:     "ABC Corp.",
		Website:     "https://www.coolapp2000.com",
		Source:      "https://github.com/coolapp2000/source",
		License:     "Apache 2.0",
		Description: "An amazing app that does cool things!",
	}

	metadataArray = append(metadataArray, app1, app2, app3, app4)
}

func applySingleQuery(q Query) []AppMetadata {
	var results []AppMetadata

	for _, app := range metadataArray {
		// check if app satisfies the query
		if satisfiesQuery(q, app) {
			results = append(results, app)
		}
	}

	return results
}

func queryExample1() {
	shouldReturnApp4 := Query{
		Title: StringFilter{
			Fragment: "Cool",
			Matcher:  "startswith",
		},
		Company: StringFilter{
			Fragment: "ABC Corp.",
			Matcher:  "equals",
		},
	}

	shouldReturnEmpty := Query{
		Title: StringFilter{
			Fragment: "Cool",
			Matcher:  "equals",
		},
		Company: StringFilter{
			Fragment: "ABC Corp.",
			Matcher:  "equals",
		},
	}

	results1 := applySingleQuery(shouldReturnApp4)
	results2 := applySingleQuery(shouldReturnEmpty)

	fmt.Println("results1")
	fmt.Println(results1)
	fmt.Println("results2")
	fmt.Println(results2)
}

/*func testingQueries() {
	query := Query{
		Title: StringFilter{
			Fragment: "cool",
			Matcher:  "contains",
		},
		Version: VersionFilter{
			Lower: 1,
			Upper: 5,
		},
		Maintainers: MaintainerFilter{
			NameFilter: StringFilter{
				Fragment: "John",
				Matcher:  "startswith",
			},
			EmailFilter: StringFilter{
				Fragment: "example.com",
				Matcher:  "ncontains",
			},
		},
		Company: StringFilter{
			Fragment: "Tech",
			Matcher:  "equals",
		},
		Website: StringFilter{
			Fragment: "https://",
			Matcher:  "startswith",
		},
		Source: StringFilter{
			Fragment: "github.com",
			Matcher:  "contains",
		},
		License: StringFilter{
			Fragment: "MIT",
			Matcher:  "equals",
		},
		Description: StringFilter{
			Fragment: "awesome",
			Matcher:  "contains",
		},
	}
}*/

func main() {
	fillArrayWithDummyApps()
	queryExample1()
}
