package main

import (
	"os"
)

/*

In this file, we should import only github.com/stevegt/grid-cli/v2,
and keep the rest of the code in v2/*.go.  Keep main.go minimal, and
have it call grid-cli.Main(), passing os.Args.

*/

func main() {
	grid_cli.Main(os.Args)
}
