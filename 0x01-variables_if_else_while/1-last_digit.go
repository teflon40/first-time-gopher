package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := r.Int() - math.MaxInt / 2
	l := n % 10

	if l > 5 {
		fmt.Printf("Last digit of %d is %d and is greater than 5\n", n, l)
	} else if l == 0 {
		fmt.Printf("Last digit of %d is %d and is 0\n", n, l)
	} else {
		fmt.Printf("Last digit of %d is %d and is less than 6 and not 0\n", n, l)
	}
}
