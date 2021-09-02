package main

import "fmt"

// Main function
func main() {


	map1 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	var map2 = make(map[int]string)
	map2[1] = "Sagar"
	map2[2] = "Wankhade"
	fmt.Println(map2)


	for id, pet := range map1 {
		fmt.Println(id, pet)
	}
}