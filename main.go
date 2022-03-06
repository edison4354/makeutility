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

	fmt.Println("Checking if the file" + fileName + " already exists...")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		response, err := http.Get(url_link)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		newFile, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(Green(fileName + " saved!"))
	} else {
		fmt.Println(Red(fileName + " already exists!"))
	}

	fmt.Println(Bold("Downloading html from " + url_link + " to " + fileName))
}
