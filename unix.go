// +build linux darwin

package main
import (
    "code.google.com/p/go.crypto/ssh/terminal"
	"syscall"
)    

func isTerminal() bool{
    return terminal.IsTerminal(syscall.Stdin)
}