package main

import (
	"fmt"
	"golang-practice/getsetMethods/calendar"
	"log"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	event := calendar.Event{}
	err := event.SetTitle("Old New Year")
	errCheck(err)
	err = event.SetYear(2022)
	errCheck(err)
	err = event.SetMonth(12)
	errCheck(err)
	err = event.SetDay(26)
	errCheck(err)
	fmt.Println("Title:", event.Title())
	fmt.Println("Year:", event.Year())
	fmt.Println("Month:", event.Month())
	fmt.Println("Day:", event.Day())
}
