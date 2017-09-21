/*This program is the game FizzBuzz, This prints the numbers 1-100 but where a number is divisable by 3 fizz is printed,
Where a number is divisable by 5 buzz is printed and where a number is divisable by 3 and 5 fizzbuzz is printed*/

package main

import (
	"fmt"
)

func main() {

	
	
	for i:=0;i<=100;i++{
	
		if i%5==0 && i%3==0{
			fmt.Println("FizzBuzz")
		}else if i%3==0{
			fmt.Println("Fizz")
		}else if i%5==0 {
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
		
	}


}