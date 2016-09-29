package mondas

import "os"

func Run() {
	app := &App{
		Name: os.Args[0],
	}
	app.Run(os.Args[1:])
}
