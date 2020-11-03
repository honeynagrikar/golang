//print numbers

package main

import "fmt"

func main() {

	for i :=1; i<=6; i++{
		for j :=1; j<=i; j++{

			fmt.Printf("%d",i)
		}
		fmt.Println()
	}
}