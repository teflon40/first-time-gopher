// Print the hexadecimal digits from 0 through 9 to the letter F
package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Printf("%c", i+48)
	}
	for i := range 6 {
		fmt.Printf("%c", i+97)
	}
	fmt.Println()
}