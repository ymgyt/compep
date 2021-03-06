package compep

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Lines -
type Lines interface {
	Next() Line
	WithSeparator(string) Lines
}

// Line -
type Line interface {
	String() string
	Strings() []string
	Int64() int64
	Int() int
	Ints() []int
}

// NewLines -
func NewLines(r io.Reader, bufSize int) Lines {
	return newLines(r, bufSize)
}

type lines struct {
	br      *bufio.Reader
	bufSize int
	sep     string
}

func newLines(r io.Reader, bufSize int) *lines {
	return &lines{
		br:  bufio.NewReaderSize(r, bufSize),
		sep: " ",
	}
}

// Next -
func (l *lines) Next() Line {
	return &line{b: l.readLine(), sep: l.sep}
}

// WithSeparator -
func (l *lines) WithSeparator(sep string) Lines {
	l.sep = sep
	return l
}

func (l *lines) readLine() []byte {
	var b []byte
	for {
		line, next, err := l.br.ReadLine()
		if err != nil {
			panic(err)
		}
		b = append(b, line...)
		if !next {
			break
		}
	}
	return b
}

type line struct {
	b   []byte
	sep string
}

// String -
func (l *line) String() string {
	return string(l.b)
}

func (l *line) Int64() int64 {
	i, err := strconv.ParseInt(string(l.b), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (l *line) Int() int {
	return int(l.Int64())
}

func (l *line) Ints() []int {
	var ns []int
	for _, f := range l.fields() {
		n, err := strconv.ParseInt(f, 10, 0)
		if err != nil {
			panic(err)
		}
		ns = append(ns, int(n))
	}
	return ns
}

func (l *line) Strings() []string {
	return l.fields()
}

func (l *line) fields() []string {
	return strings.Split(string(l.b), l.sep)
}
