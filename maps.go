package main

import "fmt"

func main(){

	studentAge := make(map[string] int)

	studentAge["Dhawal"] = 22
	studentAge["Buda"] = 23
	studentAge["ketan"]= 21
	studentAge["Nikesh"]= 27
	studentAge["Rohit"]= 26
	studentAge["Hobo"]= 25
	studentAge["goma"]= 20

    fmt.Println(studentAge) // to print all key : values

	fmt.Println(studentAge["Nikesh"]) // to print a particular employee age

	fmt.Println("the total value are :",len(studentAge)) // tp print the length (means total values)
}