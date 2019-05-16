package main

import (
	"log"
	"template/types"
)

type ExampleLister struct {
	examples []types.Example
}

func main() {
	example := types.ExampleLinkedList()
	for ex := example; ex != nil; ex = ex.NextNode {
		log.Print(ex.FirstName + " " + ex.LastName)
	}
}