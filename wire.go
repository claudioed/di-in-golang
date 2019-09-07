//+build wireinject

package main

import (
	"di-in-golang/pkg"
	"github.com/google/wire"
)

func SetupApplication() (*pkg.WordGenerator, error) {
	wire.Build(pkg.NewWordGenerator, pkg.NewDictionary)
	return nil, nil
}
