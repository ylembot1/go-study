package main

import "fmt"

type AnimalType int

const (
	AnimalTypeDog AnimalType = iota
	AnimalTypeCat
	AnimalTypeBird
)

var animalName = map[AnimalType]string{
	AnimalTypeDog:  "dog",
	AnimalTypeCat:  "cat",
	AnimalTypeBird: "bird",
}

func (a AnimalType) String() string {
	return animalName[a]
}

func main() {
	fmt.Println(AnimalTypeDog)
	fmt.Println(AnimalTypeCat)
	fmt.Println(AnimalTypeBird)
}
