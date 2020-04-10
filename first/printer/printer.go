// Package printer is a test package to learn export functions and import packages.
package printer

import "fmt"

// Hello is an exported function.
func Hello() {
	fmt.Println("exported hello")
}
