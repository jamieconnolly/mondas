package mondas

import (
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
)

var (
	BinDir, _ = osext.ExecutableFolder()
	Name = filepath.Base(os.Args[0])
	LibexecDir = filepath.Join(BinDir, "../libexec")
)
