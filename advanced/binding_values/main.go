package main

import "fmt"

type Foo struct {
	X int
}

func main() {
	f := InitializeFoo()
	fmt.Println(f)
}
