package main

import "fmt"

func main() {

	arr := [7]string{"a", "b", "c", "d",
		"e", "f", "g"}

	fmt.Println("Array:", arr)
	slice1 := arr[1:6]
	slice2 := arr[2:4]
	slice3 := arr[0:7]
	fmt.Println("Slice:", slice1)
	fmt.Println("Slice:", slice2)
	fmt.Println("Slice:", slice3)
}