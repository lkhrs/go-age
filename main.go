package main

import (
	"bufio"
	"fmt"
	"github.com/icza/gox/timex"
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
func timeElapsed(first, second string) (firstDate, secondDate time.Time, year, month, day int) {
	firstDate = dateCheck(first)
	secondDate = dateCheck(second)
	year, month, day, _, _, _ = timex.Diff(firstDate, secondDate)
	return
}

// Set up the prompts, compare the inputs, and print the result.
func main() {
	firstDate := prompt("Enter date (YYYY-MM-DD): ")
	secondDate := prompt("Enter second date (press enter for today): ")
	first, second, year, month, day := timeElapsed(firstDate, secondDate)
	fmt.Printf("%v years, %v months, %v days between %v and %v", year, month, day, first.Format(time.DateOnly), second.Format(time.DateOnly))
}
