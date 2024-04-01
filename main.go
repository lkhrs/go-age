package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Make sure the dates are valid
// If the input is an empty string, use time.Now()
func dateCheck(input string) time.Time {
	if input == "" {
		input = time.Now().Format(time.DateOnly)
	}
	date, err := time.Parse(time.DateOnly, input)
	if err != nil {
		panic(err)
	}
	return date
}

// Handle terminal input
func prompt(prompt string) (input string) {
	// Open a reader for stdin
	r := bufio.NewReader(os.Stdin)
	// Print prompt to stderr
	fmt.Fprintf(os.Stderr, prompt)
	// Read the string up to and including new line
	input, _ = r.ReadString('\n')
	// Trim whitespace and new line
	input = strings.TrimSpace(input)
	return
}

// Calculate the time between two dates
// TODO: Calculate years, months, days, and minutes
func timeElapsed(first, second string) (firstDate, secondDate time.Time, years float64) {
	firstDate = dateCheck(first)
	secondDate = dateCheck(second)
	elapsed := secondDate.Sub(firstDate).Minutes()
	years = elapsed / 60 / 24 / 365
	return
}

// Set up the prompts, compare the inputs, and print the result.
// TODO: Print "It has been x years, x months, x days, and x minutes since date."
func main() {
	firstDate := prompt("Enter date (YYYY-MM-DD): ")
	secondDate := prompt("Enter second date (press enter for today): ")
	first, second, years := timeElapsed(firstDate, secondDate)
	fmt.Printf("%.1f years between %s and %s", years, first.Format(time.DateOnly), second.Format(time.DateOnly))
}
