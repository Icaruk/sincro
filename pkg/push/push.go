package push

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sincro/pkg/utils/config"
	"sincro/pkg/utils/files"
	"sincro/pkg/utils/ui"

	"github.com/fatih/color"
)

func Local() {
	config, err := config.Read()
	if err != nil {
		fmt.Println("Config file 'sincro.json' not found. Please run 'sincro init'")
		return
	}

	bar := ui.InitProgressBar(-1)

	var sourceFilesCount int32 = 0
	var filesSyncedCount int32 = 0
	var bytesSyncedCount int64 = 0

	// Iterate sources
	for _, syncItem := range config.Sync {

		// Clean path
		syncItem.Source = filepath.Clean(syncItem.Source)

		// Walk through the source folder to get all the files
		err := filepath.Walk(syncItem.Source, func(sourceFilePath string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return err
			}
			if syncItem.Source == sourceFilePath {
				return nil
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			filesSynced, bytesSynced := files.CopyFileToDestinations(
				files.WithSourceFolder(syncItem.Source),
				files.WithSourceFilePath(sourceFilePath),
				files.WithDestionationFilePaths(syncItem.Destinations),
				files.WithProgressBar(bar),
			)

			filesSyncedCount += filesSynced
			bytesSyncedCount += bytesSynced
			sourceFilesCount++

			/* 			sourceFile, err := os.OpenFile(sourceFilePath, os.O_RDONLY, 0644)
			   			if err != nil {
			   				fmt.Println(err)
			   			}
			   			defer sourceFile.Close()

			   			for _, destinationPath := range syncItem.Destinations {

			   				// Get filename from source
			   				sourceFilename := files.RemovePathParts(sourceFilePath, partsLen)

			   				// Join with destination path
			   				destinationFilePath := filepath.Join(destinationPath, sourceFilename)

			   				// Write file
			   				writeFileResult := files.WriteFileToPath(sourceFile, destinationFilePath, true)

			   				bar.Add(1)

			   				filesSynced++
			   				bytesSynced += writeFileResult.WrittenBytes

			   			} */

			return nil
		})

		if err != nil {
			log.Println("Error: could not walk")
		}

	}

	humanReadableSize := files.PrettyByteSize(bytesSyncedCount)

	fmt.Println("")
	color.Green("Success!")
	fmt.Println("  ├─ Sources:", len(config.Sync))
	fmt.Println("  ├─ Sources files:", sourceFilesCount)
	fmt.Println("  ├─ Total files synced:", filesSyncedCount)
	fmt.Println("  └─ Total transferred:", humanReadableSize)

}
