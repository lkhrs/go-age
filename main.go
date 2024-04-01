package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func prompt(prompt string) (input string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, prompt)
	input, _ = r.ReadString('\n')
	input = strings.TrimSpace(input)
	return
}

func timeElapsed(input string) (years float64) {
	date, err := time.Parse(time.DateOnly, input)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(date).Minutes()
	years = elapsed / 60 / 24 / 365
	return
}

// TODO: Print "It has been x years, x months, x days, and x minutes since date."
func main() {
	input := prompt("Enter date (YYYY-MM-DD): ")
	years := timeElapsed(input)
	fmt.Printf("%.2f years since %s", years, input)
}
