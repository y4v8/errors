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

```go
func SetPrefix(s string)
```
Sets the prefix for child errors.

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
	fmt.Println(err)
}

func fn1() (err error) {
	defer func() {
		err = errors.Append(err, "Deferred error")
	}()

	return errors.New("New extended error")
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
	"fmt"
	"github.com/y4v8/errors"
	"log"
)

func main() {
	errors.SetPrefix("                    ")

	err := errors.New("NEW")
	err = errors.Append(err, "APPEND")
	err = errors.Wrap(err, fmt.Errorf("%v", "WRAP"))

	log.Println(err)
}
```

### Output

```
2017/03/20 18:25:49 E main.go:12 NEW
                    + main.go:13 APPEND
                    + main.go:14 WRAP
```
