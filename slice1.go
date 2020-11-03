package main

import "fmt"

func main(){
	numSlice := []int{5,4,6,9,3}

	//slice := numSlice [1:3]  // it print the value bwtn 1st postion to 3rd position
    
    //slice := numSlice [:5]  // it print the all values
    
    slice := numSlice [0:]

	fmt.Println(slice)
}