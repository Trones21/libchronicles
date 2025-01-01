package combine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type VersionData struct {
	Version  string `json:"version"`
	Packages []struct {
		PackageName string   `json:"package"`
		Types       []string `json:"types"`
		Functions   []string `json:"functions"`
	} `json:"packages"`
}

func Run() {
	fmt.Println("Combining JSON files...")

	versionsDataDir := "../out/versions_data"
	combinedDir := "../out/combined"
	outputFile := filepath.Join(combinedDir, "combined_versions_data.json")

	// Ensure combined folder exists
	if _, err := os.Stat(combinedDir); os.IsNotExist(err) {
		os.Mkdir(combinedDir, os.ModePerm)
		fmt.Println("Created combined folder.")
	}

	// Read all JSON files in the versions_data directory
	files, err := ioutil.ReadDir(versionsDataDir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", versionsDataDir, err)
		return
	}

	var allData []VersionData

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(versionsDataDir, file.Name())

			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filePath, err)
				continue
			}

			var versionData VersionData
			if err := json.Unmarshal(data, &versionData); err != nil {
				fmt.Printf("Error parsing JSON file %s: %v\n", filePath, err)
				continue
			}

			allData = append(allData, versionData)
		}
	}

	// Save combined JSON
	combinedData, err := json.MarshalIndent(allData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling combined JSON: %v\n", err)
		return
	}

	if err := ioutil.WriteFile(outputFile, combinedData, 0644); err != nil {
		fmt.Printf("Error writing combined JSON to %s: %v\n", outputFile, err)
		return
	}

	fmt.Printf("Combined JSON saved to %s\n", outputFile)
}
