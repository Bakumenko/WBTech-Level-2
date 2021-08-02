package main

import (
	"fmt"
)

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	fmt.Println(err == nil)
	fmt.Println(err == (*customError)(nil))

	if err != nil {
		println("error")
		return
	}
	println("ok")
}
