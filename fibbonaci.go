// Write fibbonaci iteratively and recursively (bonus: use dynamic programming)

package main

import "fmt"

func fibbyrinit(a int) int{
	mem := make([]int, a+1)
	return fibbyr(a, mem)
}

func fibbyr(a int, mem []int) int {
    if a == 0 {return 0}
    if a == 1 {return 1}

    if mem[a] == 0 {
    	mem[a] = fibbyr(a-1, mem) + fibbyr(a-2, mem)
    }

    return mem[a]
}

func fibbyi(b int) int {
	count1 := 0
	count2 := 1
	count3 := 0
	if b == 0 {return 0}
    if b == 1 {return 1}

	for i := 0; i < b-1; i++ {
		count3 = count1 + count2
		count1 = count2
		count2 = count3
	}

	return count3
}


func main() {
	
	recursive := fibbyrinit(1000)
	fmt.Println(recursive)

	iterative := fibbyi(1000)
	fmt.Println(iterative)

}
