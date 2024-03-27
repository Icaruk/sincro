package watch

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sincro/pkg/utils/config"
	"sincro/pkg/utils/files"

	"github.com/fatih/color"
	"github.com/rjeczalik/notify"
)

const DEBUG = false

func Start() {

	config, err := config.Read()
	if err != nil {
		fmt.Println("Config file 'sincro.json' not found. Please run 'sincro init'")
		return
	}

	watchedFolders := []string{}

	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)

	for _, syncItem := range config.Sync {
		// Clean path
		syncItem.Source = filepath.Clean(syncItem.Source)

		// Set up a watchpoint listening for events within a directory tree rooted
		// at current working directory. Dispatch remove events to c.
		if err := notify.Watch(
			path.Join(syncItem.Source, "..."),
			c,
			notify.Write,
			notify.Remove,
		); err != nil {
			log.Fatal(err)
		}
		defer notify.Stop(c)

		watchedFolders = append(watchedFolders, syncItem.Source)

	}

	// Print watcher folders tree, knowing index
	color.Green("Watching")
	for i, folder := range watchedFolders {
		isLastIndex := i == len(watchedFolders)-1

		if isLastIndex {
			color.Blue(" └─ " + folder)
		} else {
			color.Blue(" ├─ " + folder)
		}
	}

	// Listen for events
	for {
		event, ok := <-c
		if !ok {
			return
		}
		// fmt.Println("Got event:", event)

		fmt.Println("Event: ", event.Event())

		if event.Event() == notify.Write {
			fmt.Println("Path: ", event.Path())

			syncItemParent := files.GetPathSyncItemParent(event.Path(), config.Sync)

			fmt.Println("Source parent: ", syncItemParent.Source)
			fmt.Println("Filepath: ", event.Path())
			fmt.Println("Destinations: ", syncItemParent.Destinations)

			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}

			sourceFilePathFromRoot, err := filepath.Rel(pwd, event.Path())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(sourceFilePathFromRoot)

			filesSynced, bytesSynced := files.CopyFileToDestinations(
				files.WithSourceFolder(syncItemParent.Source),
				files.WithSourceFilePath(sourceFilePathFromRoot),
				files.WithDestionationFilePaths(syncItemParent.Destinations),
			)

			fmt.Printf("Files synced: %d, bytes synced: %d\n", filesSynced, bytesSynced)

		}

		if event.Event() == notify.Remove {
			fmt.Println("[WIP] Removed path: ", event.Path())
		}

	}

	// Block main goroutine forever.
	// <-make(chan struct{})

}
