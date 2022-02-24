package main

import "fmt"

type IntElement struct {
	value int
	pointer *IntElement
}

var pointer *IntElement

func main() {
	selfMadeArray()
}

func selfMadeArray() {
	// Build an array
	elementOne := IntElement{3, nil}
	elementTwo := IntElement{2, nil}
	elementThree := IntElement{1, nil}

	elementOne.pointer = &elementTwo;
	elementTwo.pointer = &elementThree;

	fmt.Print("First element address is: ")
	fmt.Print(&elementOne)
	fmt.Print("First element value is: ")
	fmt.Print(elementOne.value)
	fmt.Print(" pointer is: ")
	fmt.Println(elementOne.pointer)
	fmt.Println()

	fmt.Print("Second element address is: ")
	fmt.Print(&elementTwo)
	fmt.Print(" value is: ")
	fmt.Print(elementTwo.value)
	fmt.Print(" pointer is: ")
	fmt.Println(elementTwo.pointer)
	fmt.Println()

	fmt.Print("Third element address is: ")
	fmt.Print(&elementThree)
	fmt.Print(" value is: ")
	fmt.Print(elementThree.value)
	fmt.Print(" pointer is: ")
	fmt.Println(elementThree.pointer)

	// Print self made array (list)
	pointer = &elementOne
	for pointer != nil {
		fmt.Println(pointer.value)
		pointer = pointer.pointer
	}
}