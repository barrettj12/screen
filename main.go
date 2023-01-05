/*
Package screen implements buffered terminal printing.

To print:

	screen.Print(...)
	screen.Printf(...)
	screen.Println(...)

To clear the terminal screen:

	screen.Clear()

The screen will not actually be updated until a call to screen.Update().
*/
package screen

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var buf *bytes.Buffer

func init() {
	buf = &bytes.Buffer{}
}

// Print is an analogue for fmt.Print.
func Print(a ...any) (n int, err error) {
	return fmt.Fprint(buf, a...)
}

// Printf is an analogue for fmt.Printf.
func Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(buf, format, a...)
}

// Println is an analogue for fmt.Println.
func Println(a ...any) (n int, err error) {
	return fmt.Fprintln(buf, a...)
}

// Clear clears the terminal screen by printing ANSI escape codes.
func Clear() {
	buf.WriteString("\033[H\033[2J")
}

// Update flushes the buffer, "committing" the printed output to os.Stdout.
func Update() error {
	_, err := io.Copy(os.Stdout, buf)
	if err != nil {
		return err
	}

	buf = &bytes.Buffer{}
	return nil
}

// Writer returns an io.Writer writing to the buffer.
func Writer() io.Writer {
	return buf
}
