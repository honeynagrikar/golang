package main

import "fmt"

func add(num1,num2 int) int{
	return num1 + num2
	
}
func sub(num3,num4 int) int{
	return num3 - num4  
}
func mul(num5,num6 int) int{
	return num5 + num6
	
}
func div(num7,num8 int) int{
	return num7 + num8
	
}
func mod(num9,num10 int) int{
	return num9 + num10
	
}


func main(){
	x,y := 6,3
	fmt.Println("the value of x=",x ,"the value of y=",y)
	fmt.Println("addition",add(x,y))
	fmt.Println("substraction",sub(x,y))
	fmt.Println("multiply",mul(x,y))
	fmt.Println("Division",div(x,y))
	fmt.Println("Mod",mod(x,y))

}













