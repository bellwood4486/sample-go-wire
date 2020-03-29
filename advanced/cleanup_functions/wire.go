//+build wireinject

package main

import (
	"os"

	"github.com/google/wire"
)

func InitializeFile(path Path) (*os.File, func(), error) {
	wire.Build(Set)
	return nil, nil, nil
}
