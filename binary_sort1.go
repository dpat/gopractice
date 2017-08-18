// Implement binary search of a sorted array of integers

package main

import "fmt"

func center(value int, bottom int, top int, a []int) {

		interval := top - bottom

		if interval%2 != 0 {
			interval = (interval-1)/2
		} else {
			interval = interval/2
		}

		center1 := bottom + interval

		if top-1 == bottom {
			fmt.Println("value not found")
			return
		}

		if value > a[center1] {
			bottom = center1
			center(value, bottom, top, a)
		} else if value < a[center1] {
			top = center1
			center(value, bottom, top, a)
		} else {
			value = center1
			fmt.Println(value)
		}

		
}

func main() {
	
	a := []int{1,2,5,7,8,10,14,17,18,20,25,29,31,32,45,47,50,52,55}
	top := len(a)
	bottom := 0

	center(47, bottom, top, a)
	
}
