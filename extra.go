package main

import "github.com/gookit/color"

var fc = color.New(color.FgYellow)
var fr = color.New(color.FgRed)
var fg = color.New(color.FgGreen)
var fw = color.New(color.BgWhite, color.FgBlack)

func sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
