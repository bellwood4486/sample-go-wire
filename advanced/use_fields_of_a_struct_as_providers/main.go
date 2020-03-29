package main

import (
	"fmt"

	"github.com/google/wire"
)

type Foo struct {
	S string
	N int
	F float64
}

//func getS(foo Foo) string {
//	return foo.S
//}

func provideFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}

var Set = wire.NewSet(
	provideFoo,
	wire.FieldsOf(new(Foo), "S"),
)

func main() {
	f := InitializeMessage()
	fmt.Println(f)
}
