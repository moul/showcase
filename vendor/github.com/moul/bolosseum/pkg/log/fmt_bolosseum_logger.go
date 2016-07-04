package log

import (
	"fmt"
	"io"
	"os"
)

type FmtBolosseumLogger struct {
	outWriter io.Writer
	errWriter io.Writer
	debug     bool
}

func NewFmtBolosseumLogger() FmtBolosseumLogger {
	return FmtBolosseumLogger{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
		debug:     false,
	}
}

func (bl *FmtBolosseumLogger) SetDebug(debug bool)        { bl.debug = debug }
func (bl *FmtBolosseumLogger) SetOutWriter(out io.Writer) { bl.outWriter = out }
func (bl *FmtBolosseumLogger) SetErrWriter(err io.Writer) { bl.errWriter = err }

func (bl *FmtBolosseumLogger) Debug(args ...interface{}) {
	if bl.debug {
		fmt.Fprintln(bl.errWriter, args...)
	}
}

func (bl *FmtBolosseumLogger) Debugf(fmtString string, args ...interface{}) {
	if bl.debug {
		fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
	}
}

func (bl *FmtBolosseumLogger) Error(args ...interface{}) {
	fmt.Fprintln(bl.errWriter, args...)
}

func (bl *FmtBolosseumLogger) Errorf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
}

func (bl *FmtBolosseumLogger) Fatal(args ...interface{}) {
	fmt.Fprintln(bl.errWriter, args...)
	os.Exit(-1)
}

func (bl *FmtBolosseumLogger) Fatalf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.errWriter, fmtString+"\n", args...)
	os.Exit(-1)
}

func (bl *FmtBolosseumLogger) Info(args ...interface{}) {
	fmt.Fprintln(bl.outWriter, args...)
}

func (bl *FmtBolosseumLogger) Infof(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.outWriter, fmtString+"\n", args...)
}

func (bl *FmtBolosseumLogger) Warn(args ...interface{}) {
	fmt.Fprintln(bl.outWriter, args...)
}

func (bl *FmtBolosseumLogger) Warnf(fmtString string, args ...interface{}) {
	fmt.Fprintf(bl.outWriter, fmtString+"\n", args...)
}
