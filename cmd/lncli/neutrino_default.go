//go:build nop
// +build nop

package main

import "github.com/urfave/cli"

// neutrinoCommands will return nil for non-neutrinorpc builds.
func neutrinoCommands() []cli.Command {
	return nil
}
