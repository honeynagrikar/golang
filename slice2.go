package main

import "fmt"

func main(){

	numsclice := []int{2,9,7,4,1}

	slice2 := make([] int, 5,10)

	copy(slice2, numsclice) // it copy the numslcie into slice2
     
     fmt.Println(numsclice)

     fmt.Println(slice2)
}