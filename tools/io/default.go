package io

import (
	"github.com/fatih/color"
)

var (
	Default UI = ColorfulConsole()
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
