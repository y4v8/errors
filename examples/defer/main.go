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
