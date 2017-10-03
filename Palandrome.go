package main

import (
	"fmt"
)
//Global Variable//
var userWord string
///////////////////

func isPalindrome(input string) string {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
		return "This word is not a palandrome, try another word!"
		}
	}
	return "This word is a palandrome"
}

func userInput (){

	fmt.Println("Please enter a word, I will tell you if its a palandrome or not.")
	fmt.Scanln(&userWord)
	
}


func main(){

	userInput()// Calls user input function
	fmt.Println(isPalindrome(userWord))//Passes word to function which checks if word is palandrome
}