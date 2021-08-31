package main

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//
//  2. Assign the result to the `perimeter` variable
//
//  3. Print the `perimeter` variable
//
// HINT
//  Formula = 2 times the width and height
//
// EXPECTED OUTPUT
//  22
//

import "fmt"

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter        int
		width, height = 5, 6
	)

	perimeter = 2*(width+height)
	fmt.Println(perimeter)
}