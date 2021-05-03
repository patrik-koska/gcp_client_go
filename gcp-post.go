package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Getting the values of the generated file and the access token
	// NOTE: You can retrieve the token with the 'gcloud auth print-access-token' command
	// after you initiated your project/account with the gcloud
	// put it into a file as 'gcloud auth print-access-token > $HOME/gcp_access_token'
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "json", "default", "Path for the generated json file")
	var tokenPath string
	flag.StringVar(&tokenPath, "token", "default", "Path for the Google Cloud API key")

	flag.Parse()
	// --------------------------------------------------------------------------------

	// Getting the content of the received files
	jsonContent, err := readFileAsString(jsonFilePath)
	if err != nil {
		panic(err)
	}

	token, err := readFileAsString(tokenPath)
	if err != nil {
		panic(err)
	}
	// -----------------------------------------------

	// Processing the file through a pointer. We remove the "POST" strings, collect the urls
	// remove the urls from it afterwards, and put the leftover jsons into a slice
	var p_jsonContent *string

	p_jsonContent = &jsonContent

	*p_jsonContent = removePOSTwords(jsonContent)

	urls := findURLsinText(jsonContent)

	*p_jsonContent = removeUrls(jsonContent)

	jsons := splitByEmptyNewline(jsonContent)

	var trimmedUrls []string
	for _, url := range urls {

		trimmedUrls = append(trimmedUrls, strings.TrimSpace(url))
	}
	// --------------------------------------------------------------------------------

	// Making the post request to the Google API's, with the proper url - json pairs
	for z, _ := range urls {
		sendPostRequest(trimmedUrls[z], jsons[z], strings.TrimSpace(token))
	}
	// ---------------------------------------------------------------------------------

}

func readFileAsString(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	text := string(content)
	return text, nil
}

func removePOSTwords(jsonContent string) string {
	cleaned := strings.ReplaceAll(jsonContent, "POST", "")
	return cleaned
}

// url regex could be: ^(?:https?:\/\/)
func findURLsinText(jsonContent string) []string {
	str := jsonContent

	re := regexp.MustCompile(`(\ https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)

	submatchall := re.FindAllString(str, -1)

	var urls []string
	for _, match := range submatchall {
		urls = append(urls, match)
	}
	return urls
}

func removeUrls(jsonContent string) string {
	str := jsonContent
	re, err := regexp.Compile(`(\ https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	if err != nil {
		panic(err)
	}
	str = re.ReplaceAllString(str, "")
	return str
}

func splitByEmptyNewline(jsonContent string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(jsonContent, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func sendPostRequest(url string, json string, apiKey string) {
	var bJson = []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bJson))
	if err != nil {
		panic(err)
	}
	var bearer = "Bearer " + apiKey
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	// mandatory, otherwise does not time out
	client.Timeout = time.Second * 60
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
