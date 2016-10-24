package mondas

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
)

type ExecCommand struct {
	name string
	path string
	summary string
	usage []string
}

func NewExecCommand(name string, path string) *ExecCommand {
	return &ExecCommand{
		name: name,
		path: path,
	}
}

func (c *ExecCommand) LoadHelp() error {
	file, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		summaryRegexp := regexp.MustCompile("^# Summary: (.*)$")
		summaryMatch := summaryRegexp.FindStringSubmatch(scanner.Text())
		if summaryMatch != nil {
			c.summary = strings.TrimSpace(summaryMatch[1])
		}

		usageRegexp := regexp.MustCompile("^# Usage: (.*)$")
		usageMatch := usageRegexp.FindStringSubmatch(scanner.Text())
		if usageMatch != nil {
			c.usage = append(c.usage, strings.TrimSpace(usageMatch[1]))
		}
	}
	return scanner.Err()
}

func (c *ExecCommand) Name() string {
	return c.name
}

func (c *ExecCommand) Run(ctx *Context) error {
	for _, arg := range ctx.Args {
		switch(arg) {
		case "--help", "-h":
			return c.ShowHelp()
		}
	}

	args := append([]string{c.path}, ctx.Args...)
	env := os.Environ()
	return syscall.Exec(c.path, args, env)
}

func (c *ExecCommand) ShowHelp() error {
	c.LoadHelp()

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name(), c.Summary())

	if len(c.Usage()) > 0 {
		fmt.Println("\nUsage:")
		for _, use := range c.Usage() {
			fmt.Printf("   %s\n", use)
		}
	}

	return nil
}

func (c *ExecCommand) Summary() string {
	return c.summary
}

func (c *ExecCommand) Usage() []string {
	return c.usage
}