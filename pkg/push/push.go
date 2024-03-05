package push

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sincro/pkg/utils/config"
	"sincro/pkg/utils/files"
	"sincro/pkg/utils/ui"
	"strings"
	"time"

	"github.com/fatih/color"
)

func Local() {
	config, err := config.Read()
	if err != nil {
		fmt.Println("Config file 'sincro.json' not found. Please run 'sincro init'")
		return
	}

	bar := ui.InitProgressBar(-1)

	var filesSynced int32 = 0
	var bytesSynced int64 = 0

	// Iterate sources
	for _, syncItem := range config.Sync {
		// Clean path
		syncItem.Source = filepath.Clean(syncItem.Source)

		// Split into parts
		rgxPathSplit := regexp.MustCompile(`\\{1,2}|\/`)
		parts := rgxPathSplit.Split(syncItem.Source, -1)
		partsOffset := len(parts)

		// Walk through the source folder to get all the files
		err := filepath.Walk(syncItem.Source, func(sourcePath string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return err
			}
			if syncItem.Source == sourcePath {
				return nil
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			sourceFile, err := os.OpenFile(sourcePath, os.O_RDONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}
			defer sourceFile.Close()

			for _, destinationPath := range syncItem.Destinations {

				// Split sourcePath path into parts and remove the first partsOffset elements
				parts = rgxPathSplit.Split(sourcePath, -1)

				// Remove partsOffset elements from begginning
				parts = parts[partsOffset:]

				// Join parts
				sourcePath = strings.Join(parts, "/")

				destinationFilePath := filepath.Join(destinationPath, sourcePath)

				// Create folders until the file
				destinationFolderPath := filepath.Dir(destinationFilePath)

				err := os.MkdirAll(destinationFolderPath, 0770)
				if err != nil {
					log.Println("Error: could not create directory", filepath.Dir(destinationFolderPath))
				}

				destinationFile, err := os.Create(destinationFilePath)
				if err != nil {
					log.Println("Error: could not create file", destinationFilePath)
				}
				defer destinationFile.Close()

				// Write file
				nBytes, err := io.Copy(destinationFile, sourceFile)
				if err != nil {
					log.Println("Error: could not copy file", destinationFilePath)
				}

				time.Sleep(time.Millisecond * 1500) //! debugging

				bar.Add(1)
				filesSynced++
				bytesSynced += nBytes

			}

			return nil
		})

		if err != nil {
			log.Println("Error: could not walk")
		}

	}

	humanReadableSize := files.PrettyByteSize(bytesSynced)

	color.Green("Success!")
	fmt.Println("  ├─ Files synced:", filesSynced)
	fmt.Println("  └─ Total transferred:", humanReadableSize)

}
