// ther are two types of variable declaration

package main

import "fmt"

var a int = 12  // package level variable declaration

func main (){

/*	
   var a int=21   // function level variable declaration
	fmt.Println(a)

// the output of program is 21 because we redeclare the value of a

*/

fmt.Println(a)
var a int=21
fmt.Println(a)

// it will print the both value of a

}

