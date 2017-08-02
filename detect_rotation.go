// Given 2 integer arrays, determine if the 2nd array is a rotated version of the 1st array. 
// Ex. Original Array A={1,2,3,5,6,7,8} Rotated Array B={5,6,7,8,1,2,3} 

package main

import "fmt"

func main() {
	
	a := []int{2, 4, 6, 8, 2, 12}
	b := []int{8, 2, 12, 2, 4, 6}
	equal := false

	if len(a)==len(b) {
		for i, num2 := range b {
			if a[0] == num2 {
				equal = true
				count := len(b) - i
				for n := 0; n < count; n++ {
					if a[n] != b[n+i]{
						equal = false
						break
					}
				}
				for m := 0; m < i; m++ {
					if b[m] != a[m+i]{
						equal = false
						break
					}
				}
			}

			if equal {
				break
			}
		}
	}

	if equal {
		fmt.Println("rotation confirmed")
	} else {
		fmt.Println("no rotation")
	}
}