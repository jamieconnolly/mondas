package mondas

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"syscall"
)

type Command struct {
	Name string
	Path string
	Summary string
	Usage []string
}

func NewCommand(name string, path string) *Command {
	return &Command{
		Name: name,
		Path: path,
	}
}

func (c *Command) Parse() error {
	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		summaryRegexp := regexp.MustCompile("^# Summary: (.*)$")
		summaryMatch := summaryRegexp.FindStringSubmatch(scanner.Text())
		if summaryMatch != nil {
			c.Summary = strings.TrimSpace(summaryMatch[1])
		}

		usageRegexp := regexp.MustCompile("^# Usage: (.*)$")
		usageMatch := usageRegexp.FindStringSubmatch(scanner.Text())
		if usageMatch != nil {
			c.Usage = append(c.Usage, strings.TrimSpace(usageMatch[1]))
		}
	}
	return scanner.Err()
}

func (c *Command) Run(arguments []string) error {
	for _, arg := range arguments {
		switch(arg) {
		case "--help", "-h":
			return c.ShowHelp()
		}
	}

	args := append([]string{c.Path}, arguments...)
	return syscall.Exec(c.Path, args, os.Environ())
}

func (c *Command) ShowHelp() error {
	c.Parse()

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name, c.Summary)

	if len(c.Usage) > 0 {
		fmt.Println("\nUsage:")
		for _, use := range c.Usage {
			fmt.Printf("   %s\n", use)
		}
	}

	return nil
}

const MaxSuggestionDistance = 3

type Commands []*Command

func (c Commands) Len() int {
	return len(c)
}

func (c Commands) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Commands) Sort() Commands {
	sort.Sort(c)
	return c
}

func (c Commands) SuggestionsFor(typedName string) Commands {
	suggestions := Commands{}
	for _, cmd := range c {
		suggestForDistance := stringDistance(typedName, cmd.Name) <= MaxSuggestionDistance
		suggestForPrefix := strings.HasPrefix(strings.ToLower(cmd.Name), strings.ToLower(typedName))
		if suggestForDistance || suggestForPrefix {
			suggestions = append(suggestions, cmd)
		}
	}
	return suggestions.Sort()
}

func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
