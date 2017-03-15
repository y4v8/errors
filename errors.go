package errors

import (
	"bytes"
	"runtime"
	"strconv"
	"errors"
	"sync/atomic"
)

type xer struct {
	number uint32
	parent *xer
	rpc []uintptr
	ext []ext
}

type ext struct {
	err error
	index int
}

var number uint32

func New(text string) *xer {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	x := &xer{
		number: atomic.AddUint32(&number, 1),
	}
	x.ext = append(x.ext, ext{ err: errors.New(text) } )
	x.rpc = append(x.rpc, pcs[0])

	return x
}

func Wrap(err error) *xer {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	x, ok := err.(*xer)
	if !ok {
		x = &xer{
			number: atomic.AddUint32(&number, 1),
		}
		x.ext = append(x.ext, ext{ err: err } )
	}
	x.rpc = append(x.rpc, pcs[0])

	return x
}

func (e *xer) Append(err error) *xer {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	x, ok := err.(*xer)
	if ok {
		x.parent = e
	}

	e.ext = append(e.ext, ext{ err: err, index: len(e.rpc) } )
	e.rpc = append(e.rpc, pcs[0])

	return e
}

func (e *xer) AppendText(text string) *xer {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	e.ext = append(e.ext, ext{ err: errors.New(text), index: len(e.rpc) } )
	e.rpc = append(e.rpc, pcs[0])

	return e
}

func (e xer) Error() string {
	var buf bytes.Buffer

	prefix := ' '
	if e.parent != nil {
		prefix = '>'
		buf.WriteString("\n")
	}
	buf.WriteRune(prefix)

	extIndex := 0
	for i := range e.rpc {
		if i > 0 {
			buf.WriteString("\n")
			buf.WriteRune(prefix)
		}
		buf.WriteRune('E')
		buf.WriteString(strconv.FormatUint( uint64(e.number), 10))
		buf.WriteRune(' ')

		fn := runtime.FuncForPC(e.rpc[i])
		if fn == nil {
			buf.WriteString("unknown")
		} else {
			file, line := fn.FileLine(e.rpc[i])
			buf.WriteString( trimSourcePath(file) )
			buf.WriteRune(':')
			buf.WriteString(strconv.Itoa(line))
		}

		if extIndex < len(e.ext) && e.ext[extIndex].index == i {
			buf.WriteRune(' ')
			buf.WriteString(e.ext[extIndex].err.Error())
			extIndex++
		}
	}

	return buf.String()
}

func trimSourcePath(path string) string {
	const (
		srcLength = 5
		srcLowerCase = "/src/"
		srcUpperCase = "/SRC/"

		vendorLength = 8
		vendorLowerCase = "/vendor/"
		vendorUpperCase = "/VENDOR/"
	)

	var si, vi int
	for i := range path {
		if si == srcLength {
			return path[i:]
		}

		if vi == vendorLength {
			return path[i:]
		}

		if path[i] == srcLowerCase[si] || path[i] == srcUpperCase[si] {
			si++
			continue
		}
		si = 0

		if path[i] == vendorLowerCase[vi] || path[i] == vendorUpperCase[vi] {
			vi++
			continue
		}
		vi = 0
	}

	return path
}
