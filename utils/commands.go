package utils

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func FindAvailableCommands(prefix string, path string) map[string]string {
	cmds := make(map[string]string)
	files, _ := filepath.Glob(filepath.Join(path, prefix+"*"))
	for _, file := range files {
		if isExecutable(file) {
			cmds[strings.TrimPrefix(filepath.Base(file), prefix)] = file
		}
	}
	return cmds
}

func isExecutable(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && !fileInfo.IsDir() && (fileInfo.Mode()&syscall.S_IXUSR) != 0
}
