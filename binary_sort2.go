// Implement binary search in a rotated array (ex. {5,6,7,8,1,2,3})

package main

import "fmt"

func center(value int, bottom int, top int, a []int) {


		interval := top - bottom
		if interval < 0 {
			interval = top + (len(a)-bottom)
		}
		

		if interval%2 != 0 {
			interval = (interval-1)/2
		} else {
			interval = interval/2
		}

		center1 := bottom + interval
		if center1 > len(a) {
			center1 = bottom + interval - len(a)
		}

		if interval < 1 {
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
	
	a := []int{31,32,45,47,50,52,55,1,2,5,7,8,10,14,17,18,20,25,29}

	last := 0
	offset := 0

	for i, num1 := range a {
		if num1 < last {
			offset = i
			break
		} else {
			last = num1
		}
	}

	bottom := offset
	top := (offset-1)

	center(50, bottom, top, a)
	
}
