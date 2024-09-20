package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const maxRadiusDigits = 4

func main() {
	printIntro()
	countCircleAreaForUsersRadius()
}

func printIntro() {
	fmt.Println()
	fmt.Println("[][][][][][][][][][]")
	fmt.Println("This program cunts area of a circle.")
}

func countCircleAreaForUsersRadius() {
	area := math.Pi * math.Pow(getRadiusFromUser(), 2)
	fmt.Println("Circle's area is: ", area)
}

func getRadiusFromUser() float64 {
	fmt.Println("Please give me radius.", maxRadiusDigits, "digits max.")
	fmt.Print("Input: ")
	return readInputToFloat(maxRadiusDigits)
}

func readInputToFloat(limit int) float64 {
	reader := bufio.NewReader(os.Stdin)

	inputString, _ := reader.ReadString('\n')
	if len(inputString) > limit {
		fmt.Println("Too many symbols...", limit, "is max")
		os.Exit(1)
	}

	radius, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return radius
}
