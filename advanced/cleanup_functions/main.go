package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/wire"
)

type Path string

func provideLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Lshortfile)
}

func provideFile(log *log.Logger, path Path) (*os.File, func(), error) {
	f, err := os.Open(string(path))
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}
	return f, cleanup, nil
}

//unlint:unused // This is used by Wire.
var Set = wire.NewSet(
	provideLogger,
	provideFile,
)

func main() {
	p := Path("/temp/log.txt")
	f, _, _ := InitializeFile(p)
	fmt.Println(f)
}
