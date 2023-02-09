package cmdutils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func printErrorf(errMsgFmt string, v ...any) {
	c := color.New(color.FgHiRed)
	f := fmt.Sprintf("ERROR: %s\n", errMsgFmt)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func failOnError(errMsg string) {
	printError(errMsg)
	os.Exit(1)
}

func failOnErrorf(format string, v ...any) {
	c := color.New(color.FgHiRed)
	f := fmt.Sprintf("ERROR: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
	os.Exit(1)
}

func checkErrorfOnFailf(err error, errMsgFmt string, v ...any) {
	if err != nil {
		failOnErrorf(errMsgFmt, v...)
	}
}

func checkErrorfOnFail(err error) {
	if err != nil {
		failOnErrorf("%v", err)
	}
}

func printInfo(msg string) {
	c := color.New(color.FgHiGreen)
	_, _ = c.Fprintf(os.Stderr, "INFO: %s\n", msg)
}

func printInfof(format string, v ...any) {
	c := color.New(color.FgHiGreen)
	f := fmt.Sprintf("INFO: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func printWarn(msg string) {
	c := color.New(color.FgHiYellow)
	_, _ = c.Fprintf(os.Stderr, "WARN: %s\n", msg)
}

func printWarnf(format string, v ...any) {
	c := color.New(color.FgHiYellow)
	f := fmt.Sprintf("WARN: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
}

func printError(errMsg string) {
	c := color.New(color.FgHiRed)
	_, _ = c.Fprintf(os.Stderr, "ERROR: %s\n", errMsg)
}
