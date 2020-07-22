/*
Author: Jasmin A. Smith
Date: 06/04/2020
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Defining future variables
	var sizePrice, coffeePrice, flavorPrice, tip, pretip, total float64
	in := bufio.NewReader(os.Stdin)

	// Asking user for coffee size
	fmt.Print("Do you want a small, medium, or large coffee? ")

	// Taking user input
	size, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// Takes off enter from user input
	size = strings.TrimRight(size, "\r\n")
	// Converts input to a uniformed form (in this case all lower case)
	size = strings.ToLower(size)

	// Sets price for certain size
	if size == "small" {
		sizePrice = 2
	} else if size == "medium" {
		sizePrice = 3
	} else {
		sizePrice = 4
	}

	// Asking user for type of coffee
	fmt.Print("Would you like brewed, espresso, or cold press? ")
	// Takes in user input
	coffee, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// Takes off enter from user input
	coffee = strings.TrimRight(coffee, "\r\n")
	// Converts input to a uniformed form (in this case all lower case)
	coffee = strings.ToLower(coffee)

	// Sets price for certain coffee
	if coffee == "brewed" {
		coffeePrice = 0
	} else if coffee == "espresso" {
		coffeePrice = .5
	} else if coffee == "cold press" {
		coffeePrice = 1
	}

	// Asking user if they want flavor
	fmt.Print("Do you want a flavored syrup? (Yes or No) ")
	// Takes user input
	response, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Takes off enter from user input
	response = strings.TrimRight(response, "\r\n")
	// Converts input to a uniformed form (in this case all lower case)
	response = strings.ToLower(response)

	// Evaluates user answer to see if needs to add extra cost
	if response == "yes" {
		// Sees the flavoring type
		fmt.Print("Would you like hazelnut, vanilla, or caramel? ")
		flavor, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}
		flavor = strings.TrimRight(flavor, "\r\n")
		flavor = strings.ToLower(flavor)
		flavorPrice = .5
		// Repeats order back
		fmt.Println("You asked for a " + size + " cup of " + coffee + " with " + flavor + " syrup.")
	} else {
		flavorPrice = 0
		fmt.Println("You asked for a " + size + " cup of " + coffee + " with no syrup.")
	}
	// Calculates cost pre tip
	pretip = sizePrice + coffeePrice + flavorPrice
	pretip = RoundUp(pretip, 2)
	fmt.Printf("Your cup of coffee costs %0.2f \n", pretip)
	// Calculates cost with type
	tip = pretip * .15
	total = tip + pretip
	total = RoundUp(total, 2)
	fmt.Printf("The price with tip is %0.2f \n", total)
}

// Rounds the float to desired place value after the decimal
func RoundUp(input float64, places int) (newValue float64) {
	var round float64
	power := math.Pow(10, float64(places))
	digit := power * input
	round = math.Ceil(digit)
	newValue = round / power
	return
}
