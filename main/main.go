package main

import (
	"log"
	"template/types"
)

func main() {
	example := types.Example{
		FirstName: "joseph",
		LastName:  "stinkleton",
	}
	closeIt := closure(&example)
	closeIt()
	log.Print(example.LastName)
}

func closure(example *types.Example) func() {
	positiveName := "smellsgoodelton"
	return func() {
		example.LastName = positiveName
	}
}