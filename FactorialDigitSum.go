package main

import (
	"fmt"
	"strconv"
	"strings"
	//"math/big"
	
)


//Finds the factorial of a integer passed through
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

//Sums the digits by taking in a string, splitting the string to a string
//array converting back to and int array and adds up all the digits
func sumDigits(result string) int {
	
	var allDigits [] string 
	var totalSum int
	allDigits = strings.Split(result,"")//Splits the string into an array of each character
	
	totalSum = 0
	
	for _, i := range allDigits{
		
		val,_ := strconv.Atoi(i)//converts the array value to an int
		totalSum +=val//keeps tally of the total of the digits
	
	}
	
	return totalSum

}

func main(){

	var num uint

	num =factorial(10)//Pass 10 factorial function
	fmt.Println(num)
	
	var s = strconv.FormatUint(uint64(num), 10)//Converts the int to a string
	
	fmt.Println(s)
	
	fmt.Println(sumDigits(s))//Passes the string to the count function

}