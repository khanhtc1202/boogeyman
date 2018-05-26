package io

import (
	"github.com/fatih/color"
	"github.com/khanhtc1202/boogeyman/infrastructure/io"
	"github.com/mattn/go-colorable"
)

var (
	Stdout     = colorable.NewColorableStdout()
	Stderr     = colorable.NewColorableStderr()
	Default UI = io.Console{Stdout: Stdout, Stderr: Stderr}
)

func Infof(format string, a ...interface{}) (n int, err error) {
	return Default.Printf(color.CyanString("INFO: ")+format, a...)
}

func Warnf(format string, a ...interface{}) (n int, err error) {
	return Default.Errorf(color.YellowString("WARN: ")+format, a...)
}

func Errorf(format string, a ...interface{}) (n int, err error) {
	return Default.Errorf(color.RedString("ERROR: ")+format, a...)
}

func Errorln(a ...interface{}) (n int, err error) {
	return Errorf("%s\n", a...)
}
