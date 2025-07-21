package main

import (
	"demo1_refactor_function/fileutil"
	"fmt"
)

func main() {
	fmt.Println("===> Testing original ReadLines (no error handling):")
	lines := fileutil.ReadLines("sample.txt")
	fmt.Println("Output:", lines)

	// fmt.Println("\n===> Testing improved ReadLinesSafe (with error handling):")
	// linesSafe, err := fileutil.ReadLinesSafe("sample.txt")
	// if err != nil {
	// 	fmt.Println("Handled error:", err)
	// }
	// fmt.Println("Output:", linesSafe)

	// fmt.Println("\n===> Testing ReadLinesSafe with non-existent file:")
	// _, err = fileutil.ReadLinesSafe("nonexistent.txt")
	// if err != nil {
	// 	fmt.Println("Handled error:", err)
	// } else {
	// 	fmt.Println("Returned empty slice as expected.")
	// }

	fmt.Println("\n===> Done.")
}
