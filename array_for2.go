package main

import "fmt"

func main(){

	EvenNum := [5]int{2,54,63,32,95}

	for i, value := range EvenNum{  
                                   // to print the values in column with indexing
		fmt.Println(value, i)
	}
}