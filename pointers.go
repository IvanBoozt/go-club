package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	i := 42
	j := 2000

	pointer := &i // point to i
	
	fmt.Print("Address in memory for variable value is: ")
	fmt.Println(pointer) // the pointer address
	fmt.Print("Variable value is: ")
	fmt.Println(*pointer) // read i through the pointer

	*pointer = 21   // set i through the pointer
	fmt.Println(i)  // see the new value of i

	fmt.Print("Address in memory for NEW variable value is: ")
	fmt.Println(pointer) // the pointer address
	fmt.Print("NEW variable value is: ")
	fmt.Println(*pointer) // read i through the pointer

	pointer = &j         // point to j
	*pointer = *pointer / 20   // divide j through the pointer

	fmt.Print("Address in memory for \"pointer\" variable value is: ")
	fmt.Println(pointer) // the pointer address
	fmt.Print("variable \"j\" value is: ")
	fmt.Println(j) // see the new value of j
}
