package files

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

type CopyFileToDestinationsOptions struct {
	sourceFolder           string
	sourceFilePath         string
	sourceFilePathLeftTrim int
	destionationFilePaths  []string
	bar                    *progressbar.ProgressBar
}

type CopyeFileToDestionationsOption func(*CopyFileToDestinationsOptions)

func WithSourceFolder(sourceFolder string) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.sourceFolder = sourceFolder
	}
}

func WithSourceFilePath(sourceFilePath string) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.sourceFilePath = sourceFilePath
	}
}

func WithSourceFilePathLeftTrim(sourceFilePathLeftTrim int) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.sourceFilePathLeftTrim = sourceFilePathLeftTrim
	}
}

func WithDestionationFilePaths(destionationFilePaths []string) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.destionationFilePaths = destionationFilePaths
	}
}

func WithProgressBar(bar *progressbar.ProgressBar) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.bar = bar
	}
}

func CopyFileToDestinations(opts ...CopyeFileToDestionationsOption) (
	filesSynced int32,
	bytesSynced int64,
) {

	options := &CopyFileToDestinationsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	options.sourceFolder = path.Clean(options.sourceFolder)
	options.sourceFilePath = path.Clean(options.sourceFilePath)

	// Split into parts
	_, sourceFilePathLeftTrim := GetPathParts(options.sourceFolder)

	// Read file
	sourceFile, err := os.OpenFile(options.sourceFilePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()

	// Iterate over destination paths
	for _, destinationPath := range options.destionationFilePaths {

		relPath, err := filepath.Rel(options.sourceFolder, options.sourceFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relPath)

		// Get filename from source
		sourceFilename := RemovePathParts(options.sourceFilePath, sourceFilePathLeftTrim)

		// Join with destination path
		destinationFilePath := filepath.Join(destinationPath, sourceFilename)

		// Write file
		writeFileResult := WriteFileToPath(sourceFile, destinationFilePath, true)

		if options.bar != nil {
			options.bar.Add(1)
		}

		filesSynced++
		bytesSynced += writeFileResult.WrittenBytes

		fmt.Printf("   Wrote file from %s to %s\n", options.sourceFilePath, destinationFilePath)

	}

	return filesSynced, bytesSynced

}
