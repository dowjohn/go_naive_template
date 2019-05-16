package types



type Example struct {
	FirstName string
	LastName  string
	NextNode  *Example
}

func ExampleLinkedList() *Example {
	firstNode := NewExample("john", "snow", nil)
	AppendNode(NewExample("jennifer", "mccollum", nil), &firstNode)
	AppendNode(NewExample("billy", "bob", nil), &firstNode)
	AppendNode(NewExample("courtney", "turner", nil), &firstNode)
	return &firstNode
}

func NewExample(first string, last string, nextNode *Example) Example {
	return Example{
		FirstName: first,
		LastName:  last,
		NextNode:  nextNode,
	}
}

func AppendNode(example Example, exampleList *Example) *Example {
	for currentExample := exampleList; currentExample != nil; currentExample = currentExample.NextNode {
		if currentExample.NextNode == nil {
			currentExample.NextNode = &example
			return exampleList
		}
	}

	return exampleList
}