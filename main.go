package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	. "github.com/logrusorgru/aurora"
)

func main() {
	url_link := ""

	if len(os.Args) == 2 {
		url_link = string(os.Args[1])
	} else {
		log.Fatal("\nError while running - Usage: go run main.go <link>\nPlease make sure you entered the link")
	}

	fileName := path.Base(url_link)

	response, _ := http.Get(url_link)
	defer response.Body.Close()

	newFile, _ := os.Create(fileName)
	defer newFile.Close()

	_, _ = io.Copy(newFile, response.Body)

	fmt.Println(Bold("Downloading html from " + url_link + " to " + fileName))
}
