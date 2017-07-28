package main

import "fmt"

func main() {
	
	popular := 0
	count := 1
	tempcount := 0
	test := [11]int{1,2,2,5,10,4,4,1,10,10,10}

	for i, num := range test {
		tempcount = 0
        for _, x := range test {
        	if num == x {
        		tempcount += 1
        	}
        if tempcount > count {
        	popular = test[i] 
        	count = tempcount
        }
        }
    }

    fmt.Println(popular)

 }