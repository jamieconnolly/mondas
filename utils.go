package mondas

import (
	"os"
	"syscall"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func isExecutable(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && !fileInfo.IsDir() && (fileInfo.Mode()&syscall.S_IXUSR) != 0
}

func stringDistance(a, b string) int {
	return levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
}
