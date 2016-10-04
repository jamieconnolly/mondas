package mondas

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
)

type Command struct {
	Name string
	Path string
	Summary string
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
	fmt.Printf("Name: %v - %v\n", c.Name, c.Summary)
	return nil
}
