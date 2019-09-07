package main

import "fmt"

func main()  {
	wordGenerator, _ := SetupApplication()
	println(fmt.Sprintf("The world is %v",wordGenerator.GenerateWord()))
}
