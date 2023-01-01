package main

import "github.com/LNKLEO/oh-my-posh/cli"

var (
	Version = "development"
)

func main() {
	cli.Execute(Version)
}
