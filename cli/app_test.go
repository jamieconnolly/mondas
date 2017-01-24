package cli_test

import (
	"os"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := cli.NewApp("foo")
	assert.Equal(t, "foo-", app.ExecPrefix)
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "foo <command> [<args>]", app.Usage)
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	app := &cli.App{Commands: cmds}

	app.AddCommand(cmd2)
	assert.Equal(t, 2, len(app.Commands))
	assert.Equal(t, cmd1.Name, app.Commands[0].Name)
	assert.Equal(t, cmd2.Name, app.Commands[1].Name)
}

func TestApp_Initialize(t *testing.T) {
	defer os.Setenv("PATH", os.Getenv("PATH"))
	os.Setenv("PATH", "")

	app := cli.NewApp("foo")
	app.ExecPath = "testdata"
	app.HelpCommand = &cli.Command{Name: "help"}

	err := app.Initialize()
	assert.Nil(t, err)
	assert.Equal(t, os.Getenv("PATH"), app.ExecPath+string(os.PathListSeparator))
	assert.Equal(t, 2, len(app.Commands))
	assert.Equal(t, app.HelpCommand.Name, app.Commands[0].Name)
	assert.Equal(t, "hello", app.Commands[1].Name)
	assert.Equal(t, "testdata/foo-hello", app.Commands[1].Path)
}

func TestApp_InitializeWithNoExecPath(t *testing.T) {
	envPath := os.Getenv("PATH")
	defer os.Setenv("PATH", envPath)

	app := cli.NewApp("foo")
	app.HelpCommand = &cli.Command{Name: "help"}

	err := app.Initialize()
	assert.Nil(t, err)
	assert.Equal(t, envPath, os.Getenv("PATH"))
}

func TestApp_InitializeWithNoHelpCommand(t *testing.T) {
	defer os.Setenv("PATH", os.Getenv("PATH"))

	app := cli.NewApp("foo")

	err := app.Initialize()
	assert.Equal(t, "No help command has been set", err.Error())
}

func TestApp_Initialized(t *testing.T) {
	app := cli.NewApp("foo")
	assert.False(t, app.Initialized())

	app.Initialize()
	assert.True(t, app.Initialized())
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1, cmd2}

	app := &cli.App{Commands: cmds}
	assert.Equal(t, cmd1.Name, app.LookupCommand("one").Name)
	assert.Equal(t, cmd2.Name, app.LookupCommand("two").Name)
	assert.Nil(t, app.LookupCommand("three"))
}

func TestApp_Run(t *testing.T) {
	var s string

	app := cli.NewApp("foo")
	app.AddCommand(&cli.Command{
		Action: func(ctx *cli.Context) error {
			s = "foo"
			return nil
		},
		Name: "foo",
	})
	app.HelpCommand = &cli.Command{Name: "help"}

	err := app.Run([]string{"foo"})
	assert.Nil(t, err)
	assert.Equal(t, "foo", s)
}

func TestApp_RunWithErrorInitializing(t *testing.T) {
	app := cli.NewApp("foo")

	err := app.Run([]string{"foo"})
	assert.Equal(t, "No help command has been set", err.Error())
}

func TestApp_RunWithHelpFlagInArguments(t *testing.T) {
	var s string

	app := cli.NewApp("foo")
	app.HelpCommand = &cli.Command{
		Action: func(ctx *cli.Context) error {
			s = ctx.Args.First()
			return nil
		},
		Name: "help",
	}

	err := app.Run([]string{"foo", "bar", "--help"})
	assert.Nil(t, err)
	assert.Equal(t, "foo", s)
}

func TestApp_RunWithNoArguments(t *testing.T) {
	var s string

	app := cli.NewApp("foo")
	app.HelpCommand = &cli.Command{
		Action: func(ctx *cli.Context) error {
			s = "foo"
			return nil
		},
		Name: "help",
	}

	err := app.Run([]string{})
	assert.Nil(t, err)
	assert.Equal(t, "foo", s)
}

func TestApp_RunWithUnknownCommand(t *testing.T) {
	app := cli.NewApp("foo")
	app.HelpCommand = &cli.Command{Name: "help"}

	err := app.Run([]string{"foo"})
	assert.Equal(t, "'foo' is not a valid command.\n", err.Error())
}

func TestApp_ShowUnknownCommandErrorWithMultipleSuggestion(t *testing.T) {
	app := cli.NewApp("foo")
	app.AddCommand(&cli.Command{Name: "bar"})
	app.AddCommand(&cli.Command{Name: "baz"})

	err := app.ShowUnknownCommandError("bat")
	assert.Equal(t, "'bat' is not a valid command.\n\nDid you mean one of these?\n\tbar\n\tbaz\n", err.Error())
}

func TestApp_ShowUnknownCommandErrorWithNoSuggestions(t *testing.T) {
	app := cli.NewApp("foo")

	err := app.ShowUnknownCommandError("foo")
	assert.Equal(t, "'foo' is not a valid command.\n", err.Error())
}

func TestApp_ShowUnknownCommandErrorWithSingleSuggestion(t *testing.T) {
	app := cli.NewApp("foo")
	app.AddCommand(&cli.Command{Name: "baz"})

	err := app.ShowUnknownCommandError("bar")
	assert.Equal(t, "'bar' is not a valid command.\n\nDid you mean this?\n\tbaz\n", err.Error())
}
