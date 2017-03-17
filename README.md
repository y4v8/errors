### Extended errors

#### Adding a file name and line number to errors.

---------------------------------------

### Overview

```go
func New(text string, a ...interface{}) error
```
New returns a new error with the given printf-formatted error message.

```go
func Append(parent error, text string, a ...interface{}) error
```
Append returns a new error with the parent error and given printf-formatted error message.

```go
func Wrap(errs ...error) error
```
Wrap returns wrapped one or more errors.

---------------------------------------

### Example 1

```go
package main

import (
	"fmt"
	"github.com/y4v8/errors"
)

func main() {
	err := fn1()
	fmt.Println(err.Error())
}

func fn1() (err error) {
	defer func() {
		err = errors.Wrap(err, errors.New("Deferred error"))
	}()

	err = errors.New("New extended error")
	return err
}
```

### Output

```
E main.go:18 New extended error
+ main.go:15 Deferred error
```

---------------------------------------

### Example 2

```go
package main

import (
	e "errors"
	"fmt"
	"github.com/y4v8/errors"
)

func main() {
	var err error

	err = e.New("New Golang error")

	err = errors.Wrap(err)

	err = errors.Append(err, "Append text error")

	er2 := errors.New("New extended error")

	er3 := e.New("New Golang error")

	err = errors.Wrap(err, er2, er3)

	fmt.Println(err.Error())
}
```

### Output

```
E main.go:14 New golang error
+ main.go:16 Append text error
+ main.go:18 New extended error
+ main.go:22 New golang error
```
