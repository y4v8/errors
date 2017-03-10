package xer

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strconv"
)

type xer struct {
	err error
	rpc []uintptr
}

func (e xer) Error() string {
	if e.err == nil {
		return "nil"
	}

	var fileName, lastFileName string
	var buf bytes.Buffer

	buf.WriteRune('[')
	for i := range e.rpc {
		fn := runtime.FuncForPC(e.rpc[i])
		if fn == nil {
			buf.WriteString("unknown")
			break
		}

		file, line := fn.FileLine(e.rpc[i])
		fileName = filepath.Base(file)
		if fileName != lastFileName {
			lastFileName = fileName
			if i > 0 {
				buf.WriteRune(' ')
			}
			buf.WriteString(lastFileName)
			buf.WriteRune(':')
		} else {
			buf.WriteRune(',')
		}

		buf.WriteString(strconv.Itoa(line))
	}
	buf.WriteString("] ")
	buf.WriteString(e.err.Error())

	return buf.String()
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	var e *xer
	var ok bool
	if e, ok = err.(*xer); !ok {
		e = &xer{err: err}
	}
	e.rpc = append(e.rpc, pcs[0])

	return e
}
