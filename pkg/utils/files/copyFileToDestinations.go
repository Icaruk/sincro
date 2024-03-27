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
	sourceFilePathFromRoot string
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

func WithSourceFilePathFromRoot(sourceFilePathFromRoot string) CopyeFileToDestionationsOption {
	return func(opts *CopyFileToDestinationsOptions) {
		opts.sourceFilePathFromRoot = sourceFilePathFromRoot
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

// CopyFileToDestinations copies a file from the source folder to the destination folder.

/*
Copies a file from the source folder to the destination folder.

Usage:

	CopyFileToDestinations(options)

The options are:

	WithSourceFolder("path/to/my/source")
	WithSourceFilePathFromRoot("path/to/my/source/file.go")
	WithDestionationFilePaths(["path/to/my/destination1", "path/to/my/other/destination2"])
	WithProgressBar(*bar)
*/
func CopyFileToDestinations(opts ...CopyeFileToDestionationsOption) (
	filesSynced int32,
	bytesSynced int64,
) {

	options := &CopyFileToDestinationsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	options.sourceFolder = path.Clean(options.sourceFolder)
	options.sourceFilePathFromRoot = path.Clean(options.sourceFilePathFromRoot)

	// Read file
	sourceFile, err := os.OpenFile(options.sourceFilePathFromRoot, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()

	// Iterate over destination paths
	for _, destinationPath := range options.destionationFilePaths {

		sourceFilename, err := filepath.Rel(options.sourceFolder, options.sourceFilePathFromRoot)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Join with destination path
		destinationFilePath := filepath.Join(destinationPath, sourceFilename)

		// Write file
		writeFileResult := WriteFileToPath(sourceFile, destinationFilePath, true)

		if options.bar != nil {
			options.bar.Add(1)
		}

		filesSynced++
		bytesSynced += writeFileResult.WrittenBytes

		// fmt.Printf("   Wrote file from %s to %s\n", options.sourceFilePathFromRoot, destinationFilePath)

	}

	return filesSynced, bytesSynced

}
