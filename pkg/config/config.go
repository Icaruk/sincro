package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sincro/pkg/id"
	"strings"
)

type JSONConfig struct {
	Version   int      `json:"version"`
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Childrens []string `json:"childrens"`
	Include   []string `json:"include"`
}

// Return JSONConfig and err
func Read() (config JSONConfig, err error) {
	// Open
	file, err := os.Open("sincro.json")
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

func Init() (reason string, success bool) {
	// Check if file exists
	_, err := os.Stat("sincro.json")

	// If it exists, exit
	if err == nil {
		return "File 'sincro.json' already exists", false
	}

	// Create json file
	file, err := os.Create("sincro.json")
	if err != nil {
		return "Could not create 'sincro.json'", false
	}
	defer file.Close()

	// Generate id
	id := id.Generate("repo1")

	jsonString := strings.TrimSpace(fmt.Sprintf(`
{
	"version": 1,
	"id": "%s",
	"type": "source",
	"childrens": [],
	"sources": []
}`,
		id))

	// Write content
	_, err = file.WriteString(jsonString)
	if err != nil {
		return "Could not write 'sincro.json'", false
	}

	return "File 'sincro.json' created successfully", true

}
