// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"fmt"
	"io"
	"os"

	"github.com/issue9/term/ansi"
)

// 前景色对照表
var foreTables = []string{
	Default: ansi.FDefault,
	Black:   ansi.FBlack,
	Red:     ansi.FRed,
	Green:   ansi.FGreen,
	Yellow:  ansi.FYellow,
	Blue:    ansi.FBlue,
	Magenta: ansi.FMagenta,
	Cyan:    ansi.FCyan,
	White:   ansi.FWhite,
}

// 背景色对照表
var backTables = []string{
	Default: ansi.BDefault,
	Black:   ansi.BBlack,
	Red:     ansi.BRed,
	Green:   ansi.BGreen,
	Yellow:  ansi.BYellow,
	Blue:    ansi.BBlue,
	Magenta: ansi.BMagenta,
	Cyan:    ansi.BCyan,
	White:   ansi.BWhite,
}

// Fprint 带色彩输出的 fmt.Fprint。
func Fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprint(foreground, background, v...))
}

// Fprintln 带色彩输出的 fmt.Fprintln。
func Fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, Sprint(foreground, background, v...))
}

// Fprintf 带色彩输出的 fmt.Fprintf。
func Fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprintf(foreground, background, format, v...))
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func Print(foreground, background Color, v ...interface{}) (int, error) {
	return Fprint(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func Println(foreground, background Color, v ...interface{}) (int, error) {
	return Fprintln(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func Printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return Fprintf(os.Stdout, foreground, background, format, v...)
}

// Sprint 带色彩输出的 fmt.Sprint，返回的字符 。
func Sprint(foreground, background Color, v ...interface{}) string {
	buf := fmt.Sprint(foreTables[foreground], backTables[background])
	buf += fmt.Sprint(v...)
	return buf + fmt.Sprint(ansi.Reset)
}

// Sprintln 带色彩输出的 fmt.Sprintln。
func Sprintln(foreground, background Color, v ...interface{}) string {
	buf := fmt.Sprint(foreTables[foreground], backTables[background])
	buf += fmt.Sprint(v...)
	return buf + fmt.Sprintln(ansi.Reset)
}

// Sprintf 带色彩输出的 fmt.Sprintf。
func Sprintf(foreground, background Color, format string, v ...interface{}) string {
	buf := fmt.Sprint(foreTables[foreground], backTables[background])
	buf += fmt.Sprintf(format, v...)
	return buf + fmt.Sprint(ansi.Reset)
}
