package main

import "fmt"

func main(){

rect1 := Rectangle{10,5}

fmt.Println("height =",rect1.height)
fmt.Println("width =",rect1.width)
fmt.Println("Area of Rectangle", rect1.area())
}

type Rectangle struct{
	height float64
	width float64
}

func (rect *Rectangle) area() float64{
	return rect.height * rect.width
}