package main

import "fmt"

func main() {

	x := 10
    
    myfunc(&x)
	fmt.Println(x)
}

func myfunc(x *int ) {
	*x = 7;
}