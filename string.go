package main

import "fmt"

/* func main() {   //simple string and it's data type

 s1 := "Hello World"

 fmt.Printf("%v,%T\n", s1,s1)	
} */

/*func main(){ //string concatination

s1 :="hello guys "
s2 := " good maorning"

fmt.Printf("%v\n", s1+s2)

} */


func main() { // string indexing

	s1 :="hello world"

fmt.Printf("%v\n", s1[1]) // it print the bit type

fmt.Printf("%v\n", string(s1[1])) // it's an string indexing 
}
