package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("%d\n", 1000)

	// Output:
	// 1,000
}
