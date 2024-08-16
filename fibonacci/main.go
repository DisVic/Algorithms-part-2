package main

import (
	"fmt"
	"math/big"
	"time"
)

var one = big.NewInt(1)

func main() {
	t := time.Now()
	fmt.Println(F(3))
	fmt.Println(time.Since(t))
}

func F(x int) *big.Int {
	if x == 0 || x == 1 {
		return one
	}
	a, b := big.NewInt(1), big.NewInt(1)

	for i := 2; i < x; i++ {
		temp := new(big.Int).Add(a, b)
		a, b = b, temp
	}
	return b
}
