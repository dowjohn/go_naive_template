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
	refFunc := ref(&example)
	refFunc()
	log.Print(example.LastName)
}

func ref(example *types.Example) func() {
	positiveName := "smellsgoodelton"
	return func() {
		example.LastName = positiveName
	}
}