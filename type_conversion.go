// we covert the data type in this program


package main

import "fmt"

func main() {

	var a int = 12
	fmt.Printf("%T,%v", a,a)

	var b float32
	b = float32(a)
    fmt.Printf("%T,%v", b,b)
}