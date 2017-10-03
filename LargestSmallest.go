package main

import (
	"fmt"
)



func main(){

	arr := []uint{10,20,40,50,60,80}//Random array of integers
	
	max:= arr[0] //
	min:= arr[0] // We need to make sure first value is max/min so that we can compare each value to this.
	
	for _, value := range arr {
		if value > max {
			max = value 
        }
		if value < min{
			min = value
		}
	}
	
	fmt.Println("The largest value is : ", max, " while the smallest is: " , min)

}