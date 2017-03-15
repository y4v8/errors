package errors

import (
	"testing"
)

func TestTrimSourcePath(t *testing.T) {
	cases := []struct {
		full, trim string
	}{
		{"D:/Gowork/src/github.com/y4v8/test/main.go", "github.com/y4v8/test/main.go" },
		{"D:/Gowork/sRc/github.com/y4v8/test/main.go", "github.com/y4v8/test/main.go" },
		{"/维基百科/src/github.com/y4v8/test/main.go", "github.com/y4v8/test/main.go" },
	}

	for _, c := range cases {
		trim := trimSourcePath(c.full)
		if trim != c.trim {
			t.Errorf(`[%v] for [%v], want [%v]`, trim, c.full, c.trim)
		}
	}
}
