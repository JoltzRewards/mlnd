//go:build nop
// +build nop

package main

import "github.com/urfave/cli"

// watchtowerCommands will return nil for non-watchtowerrpc builds.
func watchtowerCommands() []cli.Command {
	return nil
}
