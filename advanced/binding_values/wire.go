//+build wireinject

package main

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeFoo() Foo {
	wire.Build(wire.Value(Foo{X: 42}))
	return Foo{}
}

// For interface values
func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
