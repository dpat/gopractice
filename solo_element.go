// Find the only element in an array that only occurs once

package main

import "fmt"

func main() {
	
	a := []int{1,55,6,8,10,3,7,1,55,7,10,6,8}
	var solo int

	for i, num1 := range a {
		count := 0
		for _, num2 := range a {
			if num1 == num2 {
				count++
			}
			if count > 1 {
				break
			}
		}
		if count <= 1 {
			solo = a[i]
			break
		}
	}

	fmt.Println(solo)
}
