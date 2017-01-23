package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
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
	// Name is the name of the command
	Name string
	// Path is the path to the program executable
	Path string
	// Summary is the overview of the command
	Summary string
	// Usage is the one-line usage message
	Usage string
}

// Parse parses the contents of the program executable for metadata.
func (c *Command) Parse() error {
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
		key, val := parts[0], strings.TrimLeft(dedent.Dedent(parts[1]), "\r\n")

		field := reflect.ValueOf(c).Elem().FieldByName(key)
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.Slice:
				field.Set(reflect.ValueOf(strings.Split(val, "\n")))
			default:
				field.Set(reflect.ValueOf(val))
			}
		}
	}

	return scanner.Err()
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

// Less reports whether the command with index i should sort before the
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

// Parse parses all the commands. It returns itself for function chaining.
func (c Commands) Parse() Commands {
	for _, cmd := range c {
		cmd.Parse()
	}
	return c
}

// Sort sorts the list of commands. It returns itself for function chaining.
func (c Commands) Sort() Commands {
	sort.Sort(c)
	return c
}

// SuggestionsFor returns a list of commands with similar names.
func (c *Commands) SuggestionsFor(typedName string) Commands {
	suggestions := Commands{}
	for _, cmd := range *c {
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
	return suggestions
}

// Swap swaps the commands with indexes i and j.
// It is part of the sort.Interface implementation.
func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
