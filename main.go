package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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

func prompt(prompt string) (input string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, prompt)
	input, _ = r.ReadString('\n')
	input = strings.TrimSpace(input)
	return
}

func timeElapsed(first, second string) (firstDate, secondDate time.Time, years float64) {
	firstDate = dateCheck(first)
	secondDate = dateCheck(second)
	elapsed := secondDate.Sub(firstDate).Minutes()
	years = elapsed / 60 / 24 / 365
	return
}

// TODO: Print "It has been x years, x months, x days, and x minutes since date."
func main() {
	firstDate := prompt("Enter date (YYYY-MM-DD): ")
	secondDate := prompt("Enter second date (press enter for today): ")
	first, second, years := timeElapsed(firstDate, secondDate)
	fmt.Printf("%.1f years between %s and %s", years, first.Format(time.DateOnly), second.Format(time.DateOnly))
}
