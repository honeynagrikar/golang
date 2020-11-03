package main

import "fmt"

func main() {

	numslice := []int{9,7,5,4,2}

	slice3 := append(numslice, 8,1,3) // it append the value in slice

    fmt.Println("normal value", numslice)
	fmt.Println("append value",slice3)
}