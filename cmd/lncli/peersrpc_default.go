//go:build nop
// +build nop

package main

import "github.com/urfave/cli"

// peersCommands will return nil for non-peersrpc builds.
func peersCommands() []cli.Command {
	return nil
}
