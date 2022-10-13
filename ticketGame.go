package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	fmt.Printf("%-16s %-4s %-10s %-5s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")
	for x := 1; x <= 10; x++ {
		company := rand.Intn(3) + 1
		if company == 1 {
			fmt.Printf("%-17s", "SpaceX")
		} else if company == 2 {
			fmt.Printf("%-17s", "Space Adventures")
		} else {
			fmt.Printf("%-17s", "Virgin Galactic")
		}

		days := rand.Intn(30) + 20
		fmt.Printf("%4d", days)

		tripType := rand.Intn(2) + 1
		if tripType == 1 {
			fmt.Printf("%-12s", " One-way")
		} else {
			fmt.Printf("%-12s", " Round-trip")
		}

		price := rand.Intn(65) + 35
		fmt.Printf("$ %3d\n", price)

	}
}
