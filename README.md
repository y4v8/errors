## Extended errors

### Adds the filename and line numbers to an error

Example
```go
package main

import (
	goerror "errors"
	"fmt"
	"github.com/y4v8/errors"
)

func main() {
	err := f1()
	fmt.Println(err)
}

func f1() error {
	eg := g1()
	ew := errors.Wrap(eg)

	ea := errors.New("error_A1")
	ew = ew.Append(ea)

	eg = goerror.New("error_Go2")
	ew = ew.Append(eg)

	ew = ew.AppendText("error_B2")

	return ew
}

func g1() error {
	return goerror.New("error_Go1")
}

```  

Output

```
 E1 github.com/y4v8/errors/example/main.go:16 error_Go1
 E1 github.com/y4v8/errors/example/main.go:19
>E2 github.com/y4v8/errors/example/main.go:18 error_A1
 E1 github.com/y4v8/errors/example/main.go:22 error_Go2
 E1 github.com/y4v8/errors/example/main.go:24 error_B2

```
