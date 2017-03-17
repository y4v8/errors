package main

import (
	e "errors"
	"fmt"
	"github.com/y4v8/errors"
)

func main() {
	var err error

	err = e.New("New golang error")

	err = errors.Wrap(err)

	err = errors.Append(err, "Append text error")

	er2 := errors.New("New extended error")

	er3 := e.New("New golang error")

	err = errors.Wrap(err, er2, er3)

	fmt.Println(err.Error())
}
