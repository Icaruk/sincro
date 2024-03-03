package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sincro/pkg/id"
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
	"$schema": "https://raw.githubusercontent.com/Icaruk/sincro/main/json-schema.json?token=GHSAT0AAAAAAB4QV72CEXAXMNBWCJT2KIRKZPCLZIA",
	"version": 1,
	"id": "%s",
	"type": "source",
	"sources": [],
	"destinations": []
}`,
		id))

	// Write content
	_, err = file.WriteString(jsonString)
	if err != nil {
		return "Could not write 'sincro.json'", false
	}

	return "File 'sincro.json' created successfully", true

}
