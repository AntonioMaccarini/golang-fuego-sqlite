package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "Woof!" }
func (Dog) Walk() string  { return "Woof!" }

type Cat struct{}

func (Cat) Speak() string { return "Meow!" }

func main() {
	var animalCat Animal = Dog{}

	var cat Cat
	fmt.Println(cat.Speak())

	animalCat.Speak()
}
