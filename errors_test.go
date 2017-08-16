package errors

import (
	"testing"
	"strings"
	e "errors"
)

func TestTrimSourcePath(t *testing.T) {
	cases := []struct {
		name, path, result string
	}{
		{
			name:   "main.main",
			path:   "D:/Gowork/src/github.com/y4v8/test/main.go",
			result: "main.go",
		}, {
			name:   "main.main",
			path:   "D:/main.go",
			result: "main.go",
		}, {
			name:   "github.com/y4v8/test/m.MyFunc",
			path:   "D:/Gowork/src/github.com/y4v8/test/m/tst.go",
			result: "github.com/y4v8/test/m/tst.go",
		},
	}

	for _, c := range cases {
		result := trimSourcePath(c.name, c.path)

		if result != c.result {
			t.Errorf(`("%v", "%v") = "%v", expected "%v"`, c.name, c.path, result, c.result)
		}
	}
}

func TestNew(t *testing.T) {
	cases := []struct {
		text   string
		params []interface{}
		suffix string
	}{
		{
			text:   "test error %v %v %v",
			params: []interface{}{1, 2, 3},
			suffix: "test error 1 2 3",
		}, {
			text:   "test error",
			params: nil,
			suffix: "test error",
		},
	}

	for _, c := range cases {
		err := New(c.text, c.params...)

		result := err.Error()
		if !strings.HasSuffix(result, c.suffix) {
			t.Errorf(`("%v", %v) = "%v", expected "%v"`, c.text, c.params, result, c.suffix)
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		err      error
		text     string
		params   []interface{}
		suffixes []string
	}{
		{
			err:      e.New("test error"),
			text:     "append error %v %v %v",
			params:   []interface{}{1, 2, 3},
			suffixes: []string{"test error", "append error 1 2 3"},
		}, {
			err:      nil,
			text:     "append error %v %v %v",
			params:   []interface{}{1, 2, 3},
			suffixes: []string{"append error 1 2 3"},
		},
	}

	for _, c := range cases {
		err := Append(c.err, c.text, c.params...)

		result := errorSuffixes(err, c.suffixes)

		resultString := strings.Join(result, ",")
		testString := strings.Join(c.suffixes, ",")
		if resultString != testString {
			t.Errorf(`("%v", "%v", %v) = %v, expected %v`, c.err, c.text, c.params, result, c.suffixes)
		}
	}
}

func TestWrap(t *testing.T) {
	cases := []struct {
		errors   []error
		suffixes []string
	}{
		{
			errors:   []error{e.New("error1"), e.New("error2"), e.New("error3")},
			suffixes: []string{"error1", "error2", "error3"},
		}, {
			errors:   []error{nil},
			suffixes: []string{""},
		}, {
			errors:   []error{nil, e.New("error2"), e.New("error3")},
			suffixes: []string{"error2", "error3"},
		}, {
			errors:   []error{e.New("error1"), nil, e.New("error3")},
			suffixes: []string{"error1", "error3"},
		},
	}

	for _, c := range cases {
		err := Wrap(c.errors...)

		result := errorSuffixes(err, c.suffixes)

		resultString := strings.Join(result, ",")
		testString := strings.Join(c.suffixes, ",")
		if resultString != testString {
			t.Errorf(`(%v) = %v, expected %v`, c.errors, result, c.suffixes)
		}
	}
}

func errorSuffixes(err error, suffixes []string) []string {
	if err == nil {
		return []string{}
	}
	split := strings.Split(err.Error(), "\n")
	result := make([]string, len(split))
	for i := range split {
		if len(split[i]) > len(suffixes[i]) {
			result[i] = split[i][len(split[i])-len(suffixes[i]):]
		} else {
			result[i] = split[i]
		}
	}
	return result
}
