package main

import "fmt"

func main(){

	num := 5
	fmt.Println(fact(num))
}
func fact(num int) int{

	if num ==0{
		return 1
	}
	return num * fact (num-1)
}