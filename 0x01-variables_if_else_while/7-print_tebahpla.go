// Print the alphabets in lowercase backwards
package main

import "fmt"

func main() {
	for i := 'z'; i >= 97; i-- {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}