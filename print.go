package cmdutils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func PrintErrorf(errMsgFmt string, v ...any) {
	c := color.New(color.FgHiRed)
	f := fmt.Sprintf("ERROR: %s\n", errMsgFmt)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func FailOnError(errMsg string) {
	PrintError(errMsg)
	os.Exit(1)
}

func FailOnErrorf(format string, v ...any) {
	c := color.New(color.FgHiRed)
	f := fmt.Sprintf("ERROR: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
	os.Exit(1)
}

func CheckErrorfOnFailf(err error, errMsgFmt string, v ...any) {
	if err != nil {
		FailOnErrorf(errMsgFmt, v...)
	}
}

func CheckErrorfOnFail(err error) {
	if err != nil {
		FailOnErrorf("%v", err)
	}
}

func PrintInfo(msg string) {
	c := color.New(color.FgHiGreen)
	_, _ = c.Fprintf(os.Stderr, "INFO: %s\n", msg)
}

func PrintInfof(format string, v ...any) {
	c := color.New(color.FgHiGreen)
	f := fmt.Sprintf("INFO: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func PrintWarn(msg string) {
	c := color.New(color.FgHiYellow)
	_, _ = c.Fprintf(os.Stderr, "WARN: %s\n", msg)
}

func PrintWarnf(format string, v ...any) {
	c := color.New(color.FgHiYellow)
	f := fmt.Sprintf("WARN: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func PrintError(errMsg string) {
	c := color.New(color.FgHiRed)
	_, _ = c.Fprintf(os.Stderr, "ERROR: %s\n", errMsg)
}
