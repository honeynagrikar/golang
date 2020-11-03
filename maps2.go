// map inside map

package main

import "fmt"

func main(){

	superhero := map[string]map[string]string{

		"Ironman" : map[string]string{
		"realname" : "Tony Stark",
		"city" : "America" ,
		},

		"Captain America" : map[string]string{
        "realname" : "Steve Rodgers",
        "city" : "Los Angeles",

		},
	}
	if temp, hero := superhero["Ironman"]; hero{

		fmt.Println(temp["realname"], temp["city"])
	}
}