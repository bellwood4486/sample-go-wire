package main

import (
	"fmt"
	"sync"

	"github.com/google/wire"
)

type Foo int
type Bar int

func ProvideFoo() Foo { return Foo(0) }

func ProvideBar() Bar { return Bar(0) }

type FooBar struct {
	mu    sync.Mutex `wire:"-"` //nolint:unused
	MyFoo Foo
	MyBar Bar
}

//nolint:unused // This is used by Wire.
var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "*"),
	// or
	// wire.Struct(new(FooBar), "MyFoo", "MyBar"),
)

func main() {
	fb := InitializeFooBar()
	fmt.Println(fb)
}
