package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	var jsonFilePath string
	//var api_key string

	flag.StringVar(&jsonFilePath, "json", "default", "Path for the generated json file")
	//flag.StringVar(&api_key, "key", "default", "Please specify the Google Cloud API key")
	flag.Parse()
	fmt.Println("lol")

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

func processJson(jsonFile string) ([]string, error) {
	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(" https") // this can also be a regex
	//for scanner.Scan() {
	//	if
	//}
	return
}
