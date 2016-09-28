package mondas

import "os"

var CommandLine = NewApp(os.Args[0])

func Run() {
	CommandLine.Run(os.Args[1:])
}
