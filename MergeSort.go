package main

import (
	"fmt"
	"sort"
)

func main(){

	firstN := []int{2,4,6} //first slice
	secondN := []int{1,3,5}//second slice
	thirdN := append(firstN , secondN...)//third slice = first slice +second slice
	
	fmt.Println("First slice: ",firstN)
	fmt.Println("Second slice: ",secondN)
	
	fmt.Println("Both slices merged: ",thirdN)
	
	sort.Ints(thirdN)// sorts third slice
	
	fmt.Println("Sorted merged slice: ",thirdN)
	
	
}