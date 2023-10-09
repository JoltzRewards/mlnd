//go:build nop
// +build nop

package main

import "github.com/urfave/cli"

// chainCommands will return nil for non-chainrpc builds.
func chainCommands() []cli.Command {
	return nil
}
