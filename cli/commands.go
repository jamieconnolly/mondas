package cli

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
	"syscall"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// Command represents a single command within an application.
type Command struct {
	Action  func(*Context) error
	Name    string
	Path    string
	Summary string
	Usage   string
}

var (
	summaryRegexp = regexp.MustCompile("^# Summary: (.*)$")
	usageRegexp   = regexp.MustCompile("^# Usage: (.*)$")
)

func (c *Command) LoadMetadata() error {
	if c.Path == "" {
		return nil
	}

	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		summaryMatch := summaryRegexp.FindStringSubmatch(scanner.Text())
		if summaryMatch != nil {
			c.Summary = strings.TrimSpace(summaryMatch[1])
		}

		usageMatch := usageRegexp.FindStringSubmatch(scanner.Text())
		if usageMatch != nil {
			c.Usage = strings.TrimSpace(usageMatch[1])
		}
	}
	return scanner.Err()
}

// Run executes the command. It calls the Action method if present,
// otherwise it invokes the executable located at Path.
func (c *Command) Run(ctx *Context) error {
	if !c.Runnable() {
		return ctx.App.ShowInvalidCommandError(c.Name)
	}

	for _, arg := range ctx.Args {
		switch arg {
		case "--help", "-h":
			return ctx.App.HelpCommand.Run(NewContext(ctx.App, []string{c.Name}, os.Environ()))
		}
	}

	if c.Action != nil {
		return c.Action(ctx)
	}

	args := append([]string{c.Path}, ctx.Args...)

	env := ctx.Env
	env.Set("PATH", strings.Join(
		[]string{ctx.App.LibexecDir, env.Get("PATH")},
		string(os.PathListSeparator),
	))
	env.Unset("BASH_ENV")

	return syscall.Exec(c.Path, args, env.Environ())
}

// Runnable determines if the command can be run.
func (c *Command) Runnable() bool {
	return c.Action != nil || c.Path != ""
}

// MaxSuggestionDistance is the maximum Levenshtein distance allowed
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

func (c Commands) LoadMetadata() Commands {
	for _, cmd := range c {
		cmd.LoadMetadata()
	}
	return c
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
