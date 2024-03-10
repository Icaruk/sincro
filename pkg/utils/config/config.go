package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sincro/pkg/id"
	"sincro/pkg/utils/ui"
	validation "sincro/pkg/utils/validations"
	"strings"
)

type SyncItem struct {
	Source       string   `json:"source"`
	Destinations []string `json:"destinations"`
}

type JSONConfig struct {
	Version int64      `json:"version"`
	ID      string     `json:"id"`
	Type    string     `json:"type"`
	Sync    []SyncItem `json:"sync"`
}

const CONFIG_FILENAME = "sincro.json"

// Return JSONConfig and err
func Read() (config JSONConfig, err error) {
	// Open
	file, err := os.Open(CONFIG_FILENAME)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Decode
	var data JSONConfig
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return config, err
	}

	// Return
	return data, err
}

func validateProjectName(name string) error {
	_, error := validation.ValidateProjectName(name)
	return error
}

func Init() (reason string, success bool) {
	// Check if file exists
	_, err := os.Stat(CONFIG_FILENAME)

	var projectName string

	// If it exists, exit
	if err == nil {
		return "File 'sincro.json' already exists", false
	} else {
		projectName = ui.PromptText("Insert the name of the project:", "? ", validateProjectName)
		if projectName == "" {
			return "Invalid project name", false
		}
	}

	// Create json file
	file, err := os.Create(CONFIG_FILENAME)
	if err != nil {
		return fmt.Sprintf("Could not create '%s'", CONFIG_FILENAME), false
	}
	defer file.Close()

	// Generate id
	id := id.Generate(projectName)

	jsonString := strings.TrimSpace(fmt.Sprintf(`
{
	"$schema": "https://github.com/Icaruk/sincro/blob/main/json-schema.json",
	"version": 1,
	"id": "%s",
	"type": "source",
	"sync": [
		{
			"source": "",
			"destinations": []
		}
	]
}`,
		id))

	// Write content
	_, err = file.WriteString(jsonString)
	if err != nil {
		return fmt.Sprintf("Could not write '%s'", CONFIG_FILENAME), false
	}

	return fmt.Sprintf("File 'sincro.json' created successfully for project %s", projectName), true

}
