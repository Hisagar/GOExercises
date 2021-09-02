package main

import "fmt"
func main()  {

	// For loop
	fmt.Println("Simple Loop")
	for i:=1; i<=10;i++{
		fmt.Println(i)
	}

	// Go contain only for loop. but we can convert it into while and do while loop.
	// While loop
	fmt.Println("While Loop")
	j:=1;
	for ;j<=10;{
		fmt.Println(j)
		j++
	}

	// Do while loop
	fmt.Println("Do While Loop")
	k:=1
	for{
		fmt.Println(k)
		k++
		if k>10{
			break
		}
	}

	// function call
	fmt.Println("Fullname : "+concatName("Sagar","Wankhade"))



}

func concatName(fname ,  lname string) string {
	return fname+" "+lname
}
