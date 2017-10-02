package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Global Variables//

var GuessInt int//
var tries int

///////////////////

func Random (){

	rand.Seed(time.Now().UnixNano())//initalize the global Source used by rand.Intn() , Otherwise random number would keep repeating
	GuessInt = rand.Intn(100)

}


func CheckInt (userGuess int)int{

	
	if userGuess == GuessInt{
		fmt.Println("Correct")
		fmt.Println("Well done! You completed the game in ", tries + 1 , " tries!")
	}else if userGuess > GuessInt {
		fmt.Println("Incorrect Try a lower number")
		tries++//Adds a try because user is wrong
		userInput()
	}else if userGuess < GuessInt{
		fmt.Println("Incorrect Try a higher number")
		tries++//adds a try because user is wrong
		userInput()
	}
	
	
	return userGuess
}

func userInput (){

	var userGuess int //Variable to store user input
	fmt.Println("Random Number has been generated! This is guess number " ,tries+1, " :" )
	fmt.Scanln(&userGuess)
	CheckInt(userGuess)//Passes guess to function which checks if guess is right
	
}

func main(){
	
	Random()//Calls random function
	userInput()//Calls user input
	
}