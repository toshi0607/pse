package main

import (
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
	c, err := subcmd.Repository().Find(os.Args[1])
	if err != nil {
		log.Fatalf("failed to find cmd: %s err:%v", os.Args[1], err)
	}
	return c.Run(os.Args[1:])
}
