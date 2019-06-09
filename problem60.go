/*
Prime pair sets
The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
*/

package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int) (next int) {
	next = n + 2
	for !isPrime(next) {
		next += 2
	}
	return
}

func generatePrimes(limit int) (map[int]struct{}, []int) {
	primesList := make([]int, 0)
	primesMap := make(map[int]struct{})
	prime := 3
	for prime < limit {
		primesMap[prime] = struct{}{}
		primesList = append(primesList, prime)
		prime = nextPrime(prime)
	}
	return primesMap, primesList
}

func main() {
	p("Problem 60")
	primesMap, primesList := generatePrimes(1000)
	p(primesList)
	p(primesMap)
}