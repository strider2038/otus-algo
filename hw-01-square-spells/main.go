package main

import "fmt"

func main() {
	const N = 25
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			if x < y {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
