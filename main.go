package main

import (
	"github.com/vindh1er/heimda11r-web/cmd"
)

var (
	// VERSION is set during build
	version = "dev"
	commit  = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
