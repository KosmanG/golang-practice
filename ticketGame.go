package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	fmt.Println("You can take tickets.")
	fmt.Println("How many you need?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Amount: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	amount, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	const mars int = 62100000

	fmt.Printf("%-16s %-4s %-10s %-5s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")

	spaceX := rand.Intn(15) + 15
	spaceA := rand.Intn(15) + 15
	virginG := rand.Intn(15) + 15
	for x := 1; x <= amount; x++ {
		company := rand.Intn(3) + 1
		if company == 1 {
			fmt.Printf("%-17s", "SpaceX")
		} else if company == 2 {
			fmt.Printf("%-17s", "Space Adventures")
		} else {
			fmt.Printf("%-17s", "Virgin Galactic")
		}

		tripType := rand.Intn(2) + 1

		if company == 1 {
			days := mars / spaceX
			if tripType == 1 {
				fmt.Printf("%4d", days/86400)
			} else {
				fmt.Printf("%4d", days/43200)
			}
		} else if company == 2 {
			days := mars / spaceA
			if tripType == 1 {
				fmt.Printf("%4d", days/86400)
			} else {
				fmt.Printf("%4d", days/43200)
			}
		} else {
			days := mars / virginG
			if tripType == 1 {
				fmt.Printf("%4d", days/86400)
			} else {
				fmt.Printf("%4d", days/43200)
			}
		}

		if tripType == 1 {
			fmt.Printf("%-12s", " One-way")
		} else {
			fmt.Printf("%-12s", " Round-trip")
		}

		price := rand.Intn(15) + 35
		if tripType == 1 {
			fmt.Printf("$ %3d\n", price)
		} else {
			fmt.Printf("$ %3d\n", price*2)
		}
	}
}
