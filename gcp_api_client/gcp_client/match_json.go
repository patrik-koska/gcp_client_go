package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//fmt.Println("lol")
	json, err := readFileAsString("./json_file.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(json)
	processedJson, err := processJson(json)
	if err != nil {
		panic(err)
	}
	fmt.Println(processedJson)
}

func readFileAsString(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	text := string(content)
	return text, nil
}

func processJson(text string) ([]string, error) {
	initial := text

	out := strings.Split(initial, " ")
	return out, nil
}
