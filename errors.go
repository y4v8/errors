package errors

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var prefix string

type erx struct {
	pc     uintptr
	error  error
	parent *erx
}

// Implementation of the error interface.
func (e erx) Error() string {
	x, flat := &e, make([]*erx, 0, 16)
	for {
		flat = append(flat, x)
		if x.parent == nil {
			break
		}
		x = x.parent
	}
	max := len(flat) - 1

	var buf bytes.Buffer
	for i := 0; i <= max; i++ {
		x = flat[max-i]

		if i == 0 {
			buf.WriteString("E ")
		} else {
			buf.WriteString("\n")
			buf.WriteString(prefix)
			buf.WriteString("+ ")
		}

		fn := runtime.FuncForPC(x.pc)
		if fn == nil {
			buf.WriteString("unknown")
		} else {
			file, line := fn.FileLine(x.pc)
			buf.WriteString(trimSourcePath(fn.Name(), file))
			buf.WriteRune(':')
			buf.WriteString(strconv.Itoa(line))
		}

		if x.error != nil {
			buf.WriteRune(' ')
			buf.WriteString(x.error.Error())
		}
	}

	return buf.String()
}

// New returns a new error with the given printf-formatted error message.
func New(text string, a ...interface{}) error {
	x := &erx{
		pc:    getPC(3),
		error: errors.New(text),
	}
	if len(a) == 0 {
		x.error = errors.New(text)
	} else {
		x.error = fmt.Errorf(text, a...)
	}

	return x
}

// Append returns a new error with the parent error
// and given printf-formatted error message.
func Append(parent error, text string, a ...interface{}) error {
	pc := getPC(3)
	p := extend(parent, pc)

	var err error
	if len(a) == 0 {
		err = errors.New(text)
	} else {
		err = fmt.Errorf(text, a...)
	}

	if parent == nil {
		p.error = err
		return p
	}

	e := &erx{
		pc:     pc,
		error:  err,
		parent: p,
	}

	return e
}

// Wrap returns wrapped one or more errors.
func Wrap(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	pc := getPC(3)

	var x, parent *erx

loop:
	for _, err := range errs {
		if err == nil {
			continue
		}

		parent = x
		x = extend(err, pc)

		if parent == nil {
			continue
		}

		e := x
		for e.parent != nil {
			if e == parent {
				x = parent
				continue loop
			}
			e = e.parent
		}
		e.parent = parent
	}

	if x == nil {
		return nil
	}

	return x
}

// Sets the prefix for child errors.
func SetPrefix(s string) {
	prefix = s
}

func getPC(calldepth int) uintptr {
	var pcs [1]uintptr
	runtime.Callers(calldepth, pcs[:])

	return pcs[0]
}

func extend(err error, pc uintptr) *erx {
	e, ok := err.(*erx)
	if !ok {
		e = &erx{
			pc:    pc,
			error: err,
		}
	}

	return e
}

func trimSourcePath(name, path string) string {
	const sep = '/'

	indexName := strings.LastIndexFunc(name, func(r rune) bool {
		return r == sep
	})

	n := 2
	if indexName == -1 {
		n = 1
	}

	indexPath := strings.LastIndexFunc(path, func(r rune) bool {
		if r == sep {
			n--
		}
		return n == 0
	})

	if indexName == -1 {
		return path[indexPath+1:]
	}

	return name[:indexName] + path[indexPath:]
}
