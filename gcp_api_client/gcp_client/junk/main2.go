package main

import (
	"fmt"
	"flag"
	//"net/http"
	//"bytes"
	"strings"
	"io/ioutil"
	"regexp"
)

func main() {
	//fmt.Println("lol")
	var jsonFilePath string
	
	flag.StringVar(&jsonFilePath,"json", "default", "Path for the generated json file")

	flag.Parse()

	jsonContent, err := readFileAsString(jsonFilePath)
	if err != nil {
		panic(err)
	}
	
	jsonRmPost := strings.Replace(jsonContent, "POST", "", -1)

	fmt.Println("jsonRmPost variable:")
	fmt.Println(jsonRmPost)
	//splitted := splitByEmptyLine(jsonContent)
	//fmt.Println(splitted)
	//for _, element := range splitted {
	//	fmt.Println(element)
	//	fmt.Println("-----")
	//}
	// At this point, we made the 2 jsons and urls into 2 objects

	/* Aminek kellene történnie:
	- Eltávolítani a "POST" szót a string file-ból
	- Belerakni az url-eket egy listába mint 1-1 string
	- Belerakni a Json-okat egy listába
	- asszociálni a két listát mappé */

}

//func sendPostRequest(url string, json string, apiKey string) string {
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
//	req.Header.Set("X-Custom-Header", "myvalue")
//	// Authorization of the API call
//	req.Header.Set("Authorization", "Bearer " + apiKey)
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	fmt.Println(resp.Status)
//	fmt.Println(resp.Header)
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body))
//	return ""
//}

func readFileAsString(filepath string) (string,error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	text := string(content)
	return text, nil
}

func splitByEmptyLine(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")	
	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
}
