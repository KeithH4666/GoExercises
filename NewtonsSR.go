package main

import (
	"fmt"
	"math"
)
// Referenced https://gist.github.com/sighmin/9173219 to for newtons method because I was unsure what it was asking.
func Sqrt(x float64) float64 {

	if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

func userInput (){

	var userGuess int //Variable to store user input
	fmt.Println("Enter number to be square rooted: " )
	fmt.Scanln(&userGuess)
	var convert float64 = float64(userGuess)//casts my int to a float64 for newtons square root
	fmt.Println("The square root of " , userGuess , " is " ,Sqrt(convert))//Passes guess to function which checks if guess is right
	
	
}

func main() {
	userInput();
}
