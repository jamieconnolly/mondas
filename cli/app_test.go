package cli_test

import (
	"os"
	"strings"
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
	envPath := os.Getenv("PATH")
	defer os.Setenv("PATH", envPath)

	os.Setenv("PATH", string(os.PathListSeparator)+envPath)

	app := &cli.App{
		ExecPath:    "testdata",
		ExecPrefix:  "foo-",
		HelpCommand: &cli.Command{Name: "help"},
		Name:        "foo",
	}
	assert.False(t, app.Initialized())

	err := app.Initialize()
	if assert.NoError(t, err) {
		assert.Equal(t, os.Getenv("PATH"), strings.Join(
			[]string{app.ExecPath, "", envPath},
			string(os.PathListSeparator),
		))
		assert.Equal(t, 2, len(app.Commands))
		assert.Equal(t, app.HelpCommand.Name, app.Commands[0].Name)
		assert.Equal(t, "hello", app.Commands[1].Name)
		assert.Equal(t, "testdata/foo-hello", app.Commands[1].Path)
		assert.True(t, app.Initialized())
	}
}

func TestApp_Initialize_WithNoExecPath(t *testing.T) {
	envPath := os.Getenv("PATH")
	defer os.Setenv("PATH", envPath)

	app := &cli.App{
		ExecPrefix:  "foo-",
		HelpCommand: &cli.Command{Name: "help"},
		Name:        "foo",
	}
	assert.False(t, app.Initialized())

	err := app.Initialize()
	if assert.NoError(t, err) {
		assert.Equal(t, envPath, os.Getenv("PATH"))
		assert.True(t, app.Initialized())
	}
}

func TestApp_Initialize_WithNoHelpCommand(t *testing.T) {
	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}
	assert.False(t, app.Initialized())

	err := app.Initialize()
	if assert.Error(t, err, "An error was expected") {
		assert.EqualError(t, err, "the help command has not been set")
		assert.True(t, app.Initialized())
	}
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

	app := &cli.App{
		Commands: cli.Commands{
			{
				Action: func(ctx *cli.Context) error {
					s = "bar"
					return nil
				},
				Name: "bar",
			},
		},
		ExecPath:    "testdata",
		ExecPrefix:  "foo-",
		HelpCommand: &cli.Command{Name: "help"},
		Name:        "foo",
	}

	err := app.Run([]string{"bar"})
	if assert.NoError(t, err) {
		assert.Equal(t, "bar", s)
	}
}

func TestApp_Run_WithEmptyArguments(t *testing.T) {
	var s string

	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		HelpCommand: &cli.Command{
			Action: func(ctx *cli.Context) error {
				s = "bar"
				return nil
			},
			Name: "help",
		},
		Name: "foo",
	}

	err := app.Run([]string{})
	if assert.NoError(t, err) {
		assert.Equal(t, "bar", s)
	}
}

func TestApp_Run_WithErrorInitializing(t *testing.T) {
	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}

	err := app.Run([]string{"bar"})
	assert.EqualError(t, err, "the help command has not been set")
}

func TestApp_Run_WithHelpFlagInArguments(t *testing.T) {
	var s string

	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		HelpCommand: &cli.Command{
			Action: func(ctx *cli.Context) error {
				s = ctx.Args.First()
				return nil
			},
			Name: "help",
		},
		Name: "foo",
	}

	err := app.Run([]string{"bar", "baz", "--help"})
	if assert.NoError(t, err) {
		assert.Equal(t, "bar", s)
	}
}

func TestApp_Run_WithUnknownCommand(t *testing.T) {
	app := &cli.App{
		ExecPath:    "testdata",
		ExecPrefix:  "foo-",
		HelpCommand: &cli.Command{Name: "help"},
		Name:        "foo",
	}

	err := app.Run([]string{"bar"})
	assert.EqualError(t, err, "'bar' is not a valid command.\n")
}

func TestApp_ShowUnknownCommandError_WithMultipleSuggestions(t *testing.T) {
	app := &cli.App{
		Commands: cli.Commands{
			{Name: "bar"},
			{Name: "baz"},
		},
		Name: "foo",
	}

	err := app.ShowUnknownCommandError("bat")
	assert.EqualError(t, err, "'bat' is not a valid command.\n"+
		"\nDid you mean one of these?\n\tbar\n\tbaz\n")
}

func TestApp_ShowUnknownCommandError_WithNoSuggestions(t *testing.T) {
	app := &cli.App{Name: "foo"}

	err := app.ShowUnknownCommandError("bar")
	assert.EqualError(t, err, "'bar' is not a valid command.\n")
}

func TestApp_ShowUnknownCommandError_WithSingleSuggestion(t *testing.T) {
	app := &cli.App{
		Commands: cli.Commands{
			{Name: "bar"},
		},
		Name: "foo",
	}

	err := app.ShowUnknownCommandError("baz")
	assert.EqualError(t, err, "'baz' is not a valid command.\n"+
		"\nDid you mean this?\n\tbar\n")
}
