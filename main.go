package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/logrusorgru/aurora"
)

func main() {
	link := ""
	if len(os.Args) == 2 {
		link = string(os.Args[1])
	} else {
		log.Fatal("\nError while running - Usage: go run main.go <link>\n Please make sure you entered the link correctly")
	}
	fmt.Println(Bold(link))
}
