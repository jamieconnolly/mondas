package mondas

import "os"

func Run() {
	app := &App{
		Name: Name,
		LibexecDir: LibexecDir,
	}
	app.Run(os.Args[1:])
}
