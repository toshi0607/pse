package subcmd

import (
	"fmt"
)

// Help command
type Help struct{}

// NewHelp returns Help
func NewHelp() *Help {
	return &Help{}
}

// Name returns command name
func (c *Help) Name() string {
	return "help"
}

// Summary returns command summary
func (c *Help) Summary() string {
	return "Show commands"
}

// Usage shows command usage
func (c *Help) Usage() {
	fmt.Printf(`Usage: pse COMMAND [OPTIONS]
  -h --help
    Show commands
  -v --version
    Show version
Commands:
`)
	for _, c := range Repository().Commands() {
		if c.Name() == "help" {
			continue
		}
		fmt.Printf("  %-9s  %s\n", c.Name(), c.Summary())
	}
	fmt.Printf(`
See more details:
  pse COMMAND -h
`)
}

// Run runs help command
func (c *Help) Run(args []string) error {
	c.Usage()
	return nil
}
