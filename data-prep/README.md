# LibChronicles

This project explores Go's standard library evolution by extracting and analyzing metadata across Go versions.

## Structure
- `combine-json`: Script to combine JSON files in `versions_data` into a single file in `combined`.
- `get-data`: Script to extract metadata from the Go repository and save it in `versions_data`.
- `versions_data`: Folder for raw JSON data (ignored by Git).
- `combined`: Folder for the combined JSON file.

## Usage

### Prerequisites
1. Install [Go](https://golang.org/doc/install).
2. Clone this repository:
   ```bash
   git clone <repository-url>
   cd libchronicles
   ```
3. Initialize the Go module (if not already done):
   ```bash
   go mod tidy
   ```
4. Nav to the data-prep folder `cd ./data-prep` (all example commands assume your pwd is the data-prep folder)

### Setup
You **must** change this before running:
- `getdata.go:27`Ensure you have the Go source repository cloned locally and update the `goRepoPath` variable to the correct path.

You **may** want to change this before running:
- `getdata.go:46` is currently commented out, you may choose to uncomment... `//tags = tags[:5] // Limit to first 5 tags for testing`


### Commands
#### Extract Metadata (`get-data`)
This command extracts metadata from the Go repository and saves it as JSON files in the `versions_data` folder.

```bash
go run main/main.go get-data
```

- The `versions_data` folder will be created automatically if it does not exist.

#### Combine JSON Files (`combine-json`)
This command combines all JSON files in the `versions_data` folder into a single JSON file in the `combined` folder.

```bash
go run main/main.go combine-json
```

- The `combined` folder will be created automatically if it does not exist.

### Notes
- The `out` folder is ignored by Git (`.gitignore`) to prevent large JSON files from being committed.

## Example Workflow
1. Run the `get-data` command to extract metadata:
   ```bash
   go run main/main.go get-data
   ```
2. Verify that JSON files have been created in the `out/versions_data` folder.
3. Run the `combine-json` command to merge these files:
   ```bash
   go run main/main.go combine-json
   ```
4. Verify the combined JSON file in the `out/combined` folder.

## Future Improvements
- TBD


