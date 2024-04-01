package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func prompt(prompt string) (s string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, prompt)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	return
}

func timeElapsed(s string) (years float64) {
	date, err := time.Parse(time.DateOnly, s)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(date).Minutes()
	years = elapsed / 60 / 24 / 365
	return
}

// TODO: Print "It has been x years, x months, x days, and x minutes since date."
func main() {
	s := prompt("Enter date (YYYY-MM-DD): ")
	years := timeElapsed(s)
	fmt.Printf("%.2f years since %s", years, s)
}
