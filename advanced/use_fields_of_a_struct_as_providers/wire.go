//+build wireinject

package main

import "github.com/google/wire"

func InitializeMessage() string {
	wire.Build(Set)
	return ""
}
