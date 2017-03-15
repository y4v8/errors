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
