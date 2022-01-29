package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix on each line to identify the logger (but see Lmsgprefix)
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}

func (l *Logger) Print(v ...interface{}) {
	l.Output(2, fmt.Sprint(v...))
}
func (l *Logger) Output(calldepth int, s string) error {
	// now := time.Now() // get this early.
	// var file string
	// var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = []byte{}
	l.buf = l.buf[:0]
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.out = os.Stdout
	_, err := l.out.Write(l.buf)
	return err
}
func main() {
	l := Logger{}
	l.Print(1)
}
