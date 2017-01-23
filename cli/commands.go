package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/jamieconnolly/mondas/utils"
	"github.com/renstrom/dedent"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// Command represents a single command within an application.
type Command struct {
	// Action is the function called when the command is run
	Action func(*Context) error
	// Description is the description of the command
	Description string
	// Hidden determines if the command is hidden from the help list of commands
	Hidden bool
	// Name is the name of the command
	Name string
	// Path is the path to the program executable
	Path string
	// Summary is the overview of the command
	Summary string
	// Usage is the one-line usage message
	Usage string

	parsed bool
}

// Parse parses the contents of the program executable for metadata.
func (c *Command) Parse() error {
	c.parsed = true

	if _, err := exec.LookPath(c.Path); err != nil {
		return err
	}

	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(utils.ScanMetadata)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ":", 2)
		key, value := parts[0], strings.TrimLeft(dedent.Dedent(parts[1]), "\r\n")

		field := reflect.ValueOf(c).Elem().FieldByName(key)
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.Bool:
				if boolValue, err := strconv.ParseBool(string(value)); err == nil {
					field.SetBool(boolValue)
				}
			case reflect.Slice:
				field.Set(reflect.ValueOf(strings.Split(value, "\n")))
			case reflect.String:
				field.SetString(string(value))
			}
		}
	}

	return scanner.Err()
}

// Parsed returns true if the program executable has been parsed for metadata.
func (c *Command) Parsed() bool {
	return c.parsed
}

// Run executes the command.
func (c *Command) Run(ctx *Context) error {
	if c.Action != nil {
		return c.Action(ctx)
	}

	if _, err := exec.LookPath(c.Path); err != nil {
		return fmt.Errorf("'%s' appears to be a valid command, but we were not\n"+
			"able to execute it. Maybe %s is broken?", c.Name, filepath.Base(c.Path))
	}

	args := append([]string{c.Path}, ctx.Args...)
	env := ctx.Env
	env.Unset("BASH_ENV")

	return syscall.Exec(c.Path, args, env.Environ())
}

// Visible returns true if the command is visible in the help list of commands.
func (c *Command) Visible() bool {
	return !c.Hidden
}

// MaxSuggestionDistance is the maximum levenshtein distance allowed
// between command names for it to be displayed as a suggestion.
const MaxSuggestionDistance = 3

// Commands represents a list of commands.
type Commands []*Command

// Add appends a command to the list (no-op if the command name is taken).
func (c *Commands) Add(cmd *Command) {
	if c.Lookup(cmd.Name) == nil {
		*c = append(*c, cmd)
	}
}

// Len returns the number of commands in the slice.
// It is part of the sort.Interface implementation.
func (c Commands) Len() int {
	return len(c)
}

// Less returns true if the command with index i should sort before the
// command with index j. It is part of the sort.Interface implementation.
func (c Commands) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

// Lookup returns the command with the matching name, or nil if not found.
func (c *Commands) Lookup(name string) *Command {
	for _, cmd := range *c {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

// SuggestionsFor returns a list of commands with similar names.
func (c *Commands) SuggestionsFor(typedName string) Commands {
	suggestions := Commands{}
	for _, cmd := range c.Visible() {
		stringDistance := levenshtein.DistanceForStrings(
			[]rune(typedName),
			[]rune(cmd.Name),
			levenshtein.DefaultOptions,
		)
		suggestForDistance := stringDistance <= MaxSuggestionDistance
		suggestForPrefix := strings.HasPrefix(strings.ToLower(cmd.Name), strings.ToLower(typedName))
		if suggestForDistance || suggestForPrefix {
			suggestions = append(suggestions, cmd)
		}
	}
	sort.Sort(suggestions)
	return suggestions
}

// Swap swaps the commands with indexes i and j.
// It is part of the sort.Interface implementation.
func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Visible returns a list of visible commands.
func (c *Commands) Visible() Commands {
	visible := Commands{}
	for _, cmd := range *c {
		if !cmd.Parsed() {
			cmd.Parse()
		}

		if cmd.Visible() {
			visible = append(visible, cmd)
		}
	}
	sort.Sort(visible)
	return visible
}
