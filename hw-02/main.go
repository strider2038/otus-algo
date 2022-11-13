package main

import (
	"fmt"

	"github.com/strider2038/otus-algo/hw-02/luckytickets"
)

func main() {
	fmt.Println("count by brute force:", luckytickets.CountByBruteForceForN3())
	fmt.Println("count by brute force (2):", luckytickets.CountByBruteForce2ForN3())
	fmt.Println("count recursively, N=2:", luckytickets.CountRecursively(2))
	fmt.Println("count recursively, N=3:", luckytickets.CountRecursively(3))
	fmt.Println("count recursively, N=4:", luckytickets.CountRecursively(4))
	fmt.Println("count fast, N=2:", luckytickets.CountFast(2))
	fmt.Println("count fast, N=3:", luckytickets.CountFast(3))
	fmt.Println("count fast, N=4:", luckytickets.CountFast(4))
}
