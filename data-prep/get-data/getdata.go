package getdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type PackageMetadata struct {
	PackageName string   `json:"package"`
	Level       int      `json:"level"`
	Types       []string `json:"types"`
	Functions   []string `json:"functions"`
}

type VersionData struct {
	Version  string            `json:"version"`
	Packages []PackageMetadata `json:"packages"`
}

func Run() {
	// Path to the Go repository
	goRepoPath := "/home/trones/Documents/not_my_repos/go"
	srcPath := filepath.Join(goRepoPath, "src")
	outputDir := "../out/versions_data"

	// Create the directory if it doesn't exist
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Get all tags
	tags, err := getGitTags(goRepoPath)
	if err != nil {
		fmt.Printf("Error getting tags: %v\n", err)
		return
	}

	// Process each tag
	//tags = tags[:5] // Limit to first 5 tags for testing
	for _, tag := range tags {
		fmt.Printf("Processing version: %s\n", tag)

		// Checkout the tag
		if err := checkoutGitTag(goRepoPath, tag); err != nil {
			fmt.Printf("Error checking out tag %s: %v\n", tag, err)
			continue
		}

		// Scrape package data
		versionData, err := scrapePackageData(srcPath, tag)
		if err != nil {
			fmt.Printf("Error scraping data for version %s: %v\n", tag, err)
			continue
		}

		// Save to JSON
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.json", tag))
		if err := saveToJSON(versionData, outputFile); err != nil {
			fmt.Printf("Error saving JSON for version %s: %v\n", tag, err)
			continue
		}
	}

	fmt.Println("Data collection completed.")
}

func getGitTags(repoPath string) ([]string, error) {
	cmd := exec.Command("git", "tag")
	cmd.Dir = repoPath
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	tags := strings.Split(out.String(), "\n")
	return tags, nil
}

func checkoutGitTag(repoPath, tag string) error {
	cmd := exec.Command("git", "checkout", tag)
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func scrapePackageData(srcPath, version string) (VersionData, error) {
	var versionData VersionData
	versionData.Version = version

	packages, err := getPackagesRecursive(srcPath)
	if err != nil {
		return versionData, err
	}

	for _, pkg := range packages {
		level := parseLevels(pkg)
		packageMetadata := PackageMetadata{
			PackageName: pkg,
			Level:       level,
			Types:       []string{},
			Functions:   []string{},
		}
		versionData.Packages = append(versionData.Packages, packageMetadata)
	}

	return versionData, nil
}

func getPackagesRecursive(basePath string) ([]string, error) {
	var packages []string
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.Contains(path, "vendor") {
			relativePath, _ := filepath.Rel(basePath, path)
			packages = append(packages, relativePath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return packages, nil
}

func saveToJSON(data VersionData, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func parseLevels(packagePath string) int {
	return len(strings.Split(packagePath, string(filepath.Separator))) - 1
}
