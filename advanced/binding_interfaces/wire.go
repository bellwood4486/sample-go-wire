//+build wireinject

package main

import "github.com/google/wire"

func InitializeBar() Bar {
	wire.Build(Set)
	return Bar("")
}
