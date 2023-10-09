//go:build nop
// +build nop

package main

import "github.com/urfave/cli"

// autopilotCommands will return nil for non-autopilotrpc builds.
func autopilotCommands() []cli.Command {
	return nil
}
