package mondas

import (
	"log"
	"os"
)

func Run() {
	app := &App{
		Name: Name,
		LibexecDir: LibexecDir,
	}

	log.SetFlags(0)

	if err := app.Run(os.Args[1:]); err != nil {
		log.Fatalf("%s: %v", Name, err)
	}
}
