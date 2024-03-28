package files

import (
	"path/filepath"
	"sincro/pkg/utils/config"
)

type SourceWithCount struct {
	Source        string
	RequiredCount int
	Count         int
	SyncItem      config.SyncItem
}

func GetPathSyncItemParent(path string, sync []config.SyncItem) config.SyncItem {

	// Create struct
	sourcesWithCount := []SourceWithCount{}

	// Iterate over sync
	for _, syncItem := range sync {

		// Clean path
		source := filepath.Clean(syncItem.Source)

		// Split into parts
		sourceParts, sourcePartsCount := GetPathParts(source)
		pathParts, _ := GetPathParts(path)

		// Count how many parts are the same
		count := 0
		for i := 0; i < len(sourceParts); i++ {

			// Check if pathParts[i] is present in any index of sourceParts
			for j := 0; j < len(pathParts); j++ {
				if sourceParts[i] == pathParts[j] {
					count++
					break
				}
			}

		}

		// Save that number
		sourcesWithCount = append(sourcesWithCount, SourceWithCount{
			Source:        syncItem.Source,
			Count:         count,
			RequiredCount: sourcePartsCount,
			SyncItem:      syncItem,
		})

	}

	// Find the sync.Source with more parts
	maxCount := 0
	var returningSyncItem config.SyncItem

	for _, sourcesWithCountItem := range sourcesWithCount {

		// Skip if the minimum number of parts is not reached
		if sourcesWithCountItem.Count < sourcesWithCountItem.RequiredCount {
			continue
		}

		if sourcesWithCountItem.Count > maxCount {
			maxCount = sourcesWithCountItem.Count
			returningSyncItem = sourcesWithCountItem.SyncItem
		}
	}

	return returningSyncItem

}
