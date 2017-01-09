package cli

import (
	"sort"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

type Command interface {
	LoadMetadata() error
	Name() string
	Run(ctx *Context) error
	ShowHelp() error
	Summary() string
	Usage() []string
}

const MaxSuggestionDistance = 3

type Commands []Command

func (c *Commands) Add(cmd Command) {
	if c.Lookup(cmd.Name()) == nil {
		*c = append(*c, cmd)
	}
}

func (c Commands) Len() int {
	return len(c)
}

func (c Commands) Less(i, j int) bool {
	return c[i].Name() < c[j].Name()
}

func (c *Commands) LoadMetadata() Commands {
	for _, cmd := range *c {
		cmd.LoadMetadata()
	}
	return *c
}

func (c *Commands) Lookup(name string) Command {
	for _, cmd := range *c {
		if cmd.Name() == name {
			return cmd
		}
	}
	return nil
}

func (c *Commands) Sort() Commands {
	sort.Sort(*c)
	return *c
}

func (c *Commands) SuggestionsFor(typedName string) Commands {
	suggestions := Commands{}
	for _, cmd := range *c {
		stringDistance := levenshtein.DistanceForStrings(
			[]rune(typedName),
			[]rune(cmd.Name()),
			levenshtein.DefaultOptions,
		)
		suggestForDistance := stringDistance <= MaxSuggestionDistance
		suggestForPrefix := strings.HasPrefix(strings.ToLower(cmd.Name()), strings.ToLower(typedName))
		if suggestForDistance || suggestForPrefix {
			suggestions = append(suggestions, cmd)
		}
	}
	return suggestions.Sort()
}

func (c Commands) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
