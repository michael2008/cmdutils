package cmdutils

import (
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
)

var _mu = sync.Mutex{}

func FailOnError(errMsg string) {
	PrintError(errMsg)
	os.Exit(1)
}

func FailOnErrorf(format string, v ...any) {
	PrintErrorf(format, v...)
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
	_mu.Lock()
	c := color.New(color.FgHiGreen)
	_, _ = c.Fprintf(os.Stderr, "INFO: %s\n", msg)
	_mu.Unlock()
}

func PrintInfof(format string, v ...any) {
	_mu.Lock()
	c := color.New(color.FgHiGreen)
	f := fmt.Sprintf("INFO: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
	_mu.Unlock()
}

func PrintWarn(msg string) {
	_mu.Lock()
	c := color.New(color.FgHiYellow)
	_, _ = c.Fprintf(os.Stderr, "WARN: %s\n", msg)
	_mu.Unlock()
}

func PrintWarnf(format string, v ...any) {
	_mu.Lock()
	c := color.New(color.FgHiYellow)
	f := fmt.Sprintf("WARN: %s\n", format)
	_, _ = c.Fprintf(os.Stderr, f, v...)
	_mu.Unlock()
}

func PrintError(errMsg string) {
	_mu.Lock()
	c := color.New(color.FgHiRed)
	_, _ = c.Fprintf(os.Stderr, "ERROR: %s\n", errMsg)
	_mu.Unlock()
}

func PrintErrorf(errMsgFmt string, v ...any) {
	_mu.Lock()
	c := color.New(color.FgHiRed)
	f := fmt.Sprintf("ERROR: %s\n", errMsgFmt)
	_, _ = c.Fprintf(os.Stderr, f, v...)
	_mu.Unlock()
}

func PrintHeader(header string) {
	_mu.Lock()
	c := color.New(color.FgHiMagenta)
	_, _ = c.Printf("\n%s\n", header)
	_mu.Unlock()
}
