package mondas

import (
	"os"
	"path/filepath"
)

var (
	BinDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	Name = filepath.Base(os.Args[0])
	LibexecDir = filepath.Join(BinDir, "../libexec")
)
