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
	if n > 0 {
		fmt.Println(n, "is positive")
	} else if n < 0 {
		fmt.Println(n, "is negative")
	} else {
		fmt.Println(n, "is zero")
	}
}