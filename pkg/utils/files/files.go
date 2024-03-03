package files

import (
	"fmt"
	"math"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Checks if two files have the same size and mode
// This does not check the contents
func CompareFiles(path1 string, path2 string) bool {

	stat1, err1 := os.Stat(path1)
	stat2, err2 := os.Stat(path2)

	if err1 != nil || err2 != nil {
		return false
	}

	if stat1.Size() != stat2.Size() {
		return false
	}
	if stat1.Mode() != stat2.Mode() {
		return false
	}

	return true

}

// Converts bytes to human readable format
// from: https://gist.github.com/anikitenko/b41206a49727b83a530142c76b1cb82d?permalink_comment_id=4467913#gistcomment-4467913
func PrettyByteSize(b int64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.d %sB", int(bf), unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%d YiB", int(bf))
}
