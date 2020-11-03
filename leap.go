package main

import "fmt"

func main(){

	year := 2007


	if year %400 ==0 && year %100==0 || year %4==0{
		fmt.Println(year,"is a leap year")
}else{
	fmt.Println(year," is a non leap year")
}


/*	if year %4==0{
		fmt.Println("the given year is a leap year",year)
			}else{
				fmt.Println("the given year is not a leap year",year)
			}
		if year %100==0{
			fmt.Println("the given year is a leap year",year)
			}else{
				fmt.Println("the given year is not a leap year",year)
			}
			if year %400==0{
				fmt.Println("the given year is a leap year",year)
			}else{
				fmt.Println("the given year is not a leap year",year)
			} 
*/

/* if year %4==0{
 	if year %100==0{
 		if year%400==0{
 			fmt.Println("the given year is a leap year",year)
 		}else{
 			fmt.Println("the given year is a non leap year", year)
 		}
 	      fmt.Println("the given year is a leap year",year)
 		}else{
 			fmt.Println("the given year is a non leap year", year)
 	}
 
       fmt.Println("the given year is a leap year",year)
 		}else{
 			fmt.Println("the given year is a non leap year", year)
 }	*/		

				
			
}