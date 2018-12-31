package main

import (
	"flag"
	"log"
	"os"

	"github.com/toshi0607/pse/subcmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
	)
	flag.BoolVar(&showHelp, "h", false, "")
	flag.BoolVar(&showHelp, "help", false, "")
	flag.BoolVar(&showVersion, "v", false, "")
	flag.BoolVar(&showVersion, "version", false, "")

	if len(os.Args) <= 1 {
		return subcmd.NewHelp().Run(nil)
	}

	flag.Parse()
	if showHelp {
		return subcmd.NewHelp().Run(nil)
	}

	c, err := subcmd.Repository().Find(os.Args[1])
	if err != nil {
		log.Fatalf("failed to find cmd: %s err:%v", os.Args[1], err)
	}
	return c.Run(os.Args[1:])
}
