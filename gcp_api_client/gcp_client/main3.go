package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"flag"
)

func main() {
	var jsonFilePath string

	flag.StringVar(&jsonFilePath,"json", "default", "Path for the generated json file")

	flag.Parse()

	file, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ### Put this block of code into a function
	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(" https") // this can also be a regex

	if err != nil {
		log.Fatal(err)
	}

	var urls []string

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			//fmt.Println(scanner.Text())
			text := scanner.Text()
			urls := append(urls, text)
		}
	}
	// ### and return the value if matched into a variable

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(urls[0])
	fmt.Println(urls[1])
}
