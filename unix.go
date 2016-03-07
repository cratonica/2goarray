// +build linux darwin

package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func isTerminal() bool {
	return terminal.IsTerminal(syscall.Stdin)
}
