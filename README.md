## Extended errors

### Adds the filename and line numbers to an error

Example
```go
package main

import (
	"fmt"
	"errors"
	"github.com/y4v8/xer"
)

func main() {
	err := fn1()
	fmt.Println(err)
}

func fn1() error {
	err := fn2()
	return xer.Wrap(err)
}

func fn2() error {
	err := fn3()
	return xer.Wrap(err)
}

func fn3() error {
	return errors.New("TEST")
}

```  
Output

```
[main.go:21,16] TEST
```
