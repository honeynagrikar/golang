package main

import "fmt"

func main() {

	for i :=1; i<=5; i++{
		for j :=i; j<i; j++{
			fmt.Printf("  ")
				}

	for k :=5; k>i; k--{
		fmt.Printf("*")
		
	}
	fmt.Println()
	}
}