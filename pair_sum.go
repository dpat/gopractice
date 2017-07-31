//Find pairs in an integer array whose sum is equal to 10 
//Need to ask professor if this is in linear time or not

package main

import "fmt"

func main() {
	
	target := 10
	s := []int{1,5,7,5,9,9,3,8,7,7,2,11,15,0}

	for i, num1 := range s {
		for j := range s {
			count := 1+i+j
			if count == len(s) {
				break
			}
			sum := num1 + s[count]
			if sum == target {
				fmt.Println(s[i],", ",s[count])
			}
		}
	}
	
 }