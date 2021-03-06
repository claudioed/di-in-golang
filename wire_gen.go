// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"di-in-golang/pkg"
)

// Injectors from wire.go:

func SetupApplication() (*pkg.WordGenerator, error) {
	dictionary := pkg.NewDictionary()
	wordGenerator := pkg.NewWordGenerator(dictionary)
	return wordGenerator, nil
}
