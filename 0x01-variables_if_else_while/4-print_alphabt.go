// Prints the alphabets in lowercase without 'q' and 'e'
package main

import "fmt"

func main() {
	for i := 97; i <= 122; i++ {
		if i == 101 || i == 113 { // Skip 'e' and 'q'
			continue
		}
		fmt.Printf("%c", i)
	}
	fmt.Println("")
}