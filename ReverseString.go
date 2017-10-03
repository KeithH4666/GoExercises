package main

import (
	"fmt"
)

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

func userInput (){

	var userIn string //Variable to store user input
	fmt.Println("Enter a string to be reversed: " )
	fmt.Scanln(&userIn)
	fmt.Println(userIn , " is " ,Reverse(userIn), " reversed")
	
}

func main(){

	userInput()

}