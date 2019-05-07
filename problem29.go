// Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

// 2^2=4,  2^3=8,   2^4=16,  2^5=32
// 3^2=9,  3^3=27,  3^4=81,  3^5=243
// 4^2=16, 4^3=64,  4^4=256, 4^5=1024
// 5^2=25, 5^3=125, 5^4=625, 5^5=3125
// If they are then placed in numerical order, with any repeats removed,
// we get the following sequence of 15 distinct terms:

// 4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

// How many distinct terms are in the sequence generated by a^b
// for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

func generateSequency(limit int, terms map[string]struct{}) {
	for i := 2; i <= limit; i++ {
		for j := 2; j <= limit; j++ {
			a := big.NewInt(int64(i))
			b := big.NewInt(int64(j))
			term := new(big.Int)
			term = term.Exp(a, b, nil)
			// p(a, " ^ ", b, " = ", term)
			terms[term.String()] = struct{}{}
		}
	}
}

func main() {
	p("Problem 29")
	terms := make(map[string]struct{}, 0)
	generateSequency(100, terms)
	p("count = ", len(terms))
}
