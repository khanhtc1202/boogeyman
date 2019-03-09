package io

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
)

type UI interface {
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	Errorf(format string, a ...interface{}) (n int, err error)
	Errorln(a ...interface{}) (n int, err error)
}

type Console struct {
	Stdout io.Writer
	Stderr io.Writer
}

func ColorfulConsole() *Console {
	Stdout := colorable.NewColorableStdout()
	Stderr := colorable.NewColorableStderr()
	return &Console{Stdout: Stdout, Stderr: Stderr}
}

func (c Console) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Stdout, format, a...)
}

func (c Console) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.Stdout, a...)
}

func (c Console) Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Stderr, format, a...)
}

func (c Console) Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.Stderr, a...)
}
