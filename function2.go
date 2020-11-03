package main

import "fmt"

func even_odd(num1 int) int{

     if num1 %2==0{
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
	return num1
}

func main(){
 a := 7

 fmt.Println(even_odd(a))

}