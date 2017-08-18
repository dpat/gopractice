// Find the common elements of 2 int arrays
// The requirements for this question were relatively vague


package main

import "fmt"
import "strconv"

func main() {
	
	a := []int{1,55,6,8,10,3,7,8}
	b := []int{55,2,3,7,1,19,10,6,5,2,22,33,0,1,1,1,1,2,2}
	var value int
	var tempstring string
	var common bool

	for _, num1 := range a {
		value = num1
		tempstring = strconv.Itoa(value)
		common = false
		for _, num2 := range b {
			if value == num2 {
				tempstring = tempstring
				common = true
			}
	}

// loop through, if num equals no number in part a, then loop and look for matches in same array
		if common {
			fmt.Println(tempstring)
		}
	}
	
}
