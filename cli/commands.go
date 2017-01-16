package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"syscall"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

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

func (c *Command) Run(ctx *Context) error {
	if !c.Runnable() {
		return ctx.App.ShowInvalidCommandError(c.Name)
	}

	for _, arg := range ctx.Args {
		switch arg {
		case "--help", "-h":
			return c.ShowHelp()
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

func (c *Command) Runnable() bool {
	return c.Action != nil || c.Path != ""
}

func (c *Command) ShowHelp() error {
	c.LoadMetadata()

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name, c.Summary)

	if c.Usage != "" {
		fmt.Println("\nUsage:")
		fmt.Printf("   %s\n", c.Usage)
	}

	return nil
}

const MaxSuggestionDistance = 3

type Commands []*Command

func (c *Commands) Add(cmd *Command) {
	if c.Lookup(cmd.Name) == nil {
		*c = append(*c, cmd)
	}
}

func (c Commands) Len() int {
	return len(c)
}

func (c Commands) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Commands) LoadMetadata() Commands {
	for _, cmd := range c {
		cmd.LoadMetadata()
	}
	return c
}

func (c *Commands) Lookup(name string) *Command {
	for _, cmd := range *c {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

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

func (c Commands) Sort() Commands {
	sort.Sort(c)
	return c
}

func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
