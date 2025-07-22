// Print numbers from 0 to 9 as characters
package main

import "fmt"

func main() {
	for i := '0'; i <= '9'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}