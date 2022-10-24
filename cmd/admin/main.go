package main

import (
	_ "embed"
	"github.com/gotti/meshover/cmd/admin/cmd"
)

func main() {
	cmd.Execute()
}
