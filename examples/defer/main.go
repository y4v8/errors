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
