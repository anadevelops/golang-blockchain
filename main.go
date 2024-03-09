package main

import (
	"os"

	"github.com/anadevelops/golang-blockchain/blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.run()
}
