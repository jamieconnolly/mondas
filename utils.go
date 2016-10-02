package mondas

import (
	"os"
	"syscall"
)

func isExecutable(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && !fileInfo.IsDir() && (fileInfo.Mode()&syscall.S_IXUSR) != 0
}
