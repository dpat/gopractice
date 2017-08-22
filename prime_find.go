// Write fibbonaci iteratively and recursively (bonus: use dynamic programming)

package main

import "fmt"


func find_prime(a int) int {
	primes := make([]int, 0, a-1)
    primes = append(primes, 2, 3)
    counter := 5
    isprime := true
    prime_count := 1

    for len(primes) < a {

    	isprime = true
    	for i := prime_count; i >= 0; i-- {
	    	if counter % primes[i] == 0 {
	    		isprime = false
	    		break
	    	}
    	}

	    if isprime {
	    	primes = append(primes, counter)
	    	prime_count++
	    }

	    counter++
    }

    return primes[a-1]

}



func main() {
	
	prime_test := find_prime(1000)
	fmt.Println(prime_test)

	
}
