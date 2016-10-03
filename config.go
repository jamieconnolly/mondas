package mondas

import (
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
)

const MaxSuggestionDistance = 3

var (
	BinDir, _ = osext.ExecutableFolder()
	Name = filepath.Base(os.Args[0])
	LibexecDir = filepath.Join(BinDir, "../libexec")
)
