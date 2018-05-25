package io

import (
	"fmt"
	"io"
)

type Console struct {
	Stdout io.Writer
	Stderr io.Writer
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
