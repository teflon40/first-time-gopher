// Print alphabets in lower case followed by uppercase
package main

import "fmt"

func main() {
	for i := 97; i <= 122; i++ {
		fmt.Printf("%c", i)
	}
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}

