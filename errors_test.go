package errors

import (
	"testing"
)

func TestTrimSourcePath(t *testing.T) {
	cases := []struct {
		name, path, trim string
	}{
		{"main.main", "D:/Gowork/src/github.com/y4v8/test/main.go", "main.go"},
		{"main.main", "D:/main.go", "main.go"},
		{"github.com/y4v8/test/m.MyFunc", "D:/Gowork/src/github.com/y4v8/test/m/tst.go", "github.com/y4v8/test/m/tst.go"},
	}

	for _, c := range cases {
		trim := trimSourcePath(c.name, c.path)
		if trim != c.trim {
			t.Errorf(`[%v] for path=[%v], want [%v]`, trim, c.path, c.trim)
		}
	}
}
