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
	var jsonFilePath string
	flag.StringVar(&jsonFilePath, "json", "default", "Path for the generated json file")
	var tokenPath string
	flag.StringVar(&tokenPath, "token", "default", "Path for the Google Cloud API key")

	flag.Parse()

	jsonContent, err := readFileAsString(jsonFilePath)
	if err != nil {
		panic(err)
	}

	token, err := readFileAsString(tokenPath)
	if err != nil {
		panic(err)
	}

	//fmt.Println("-- printing jsonContent --")
	//fmt.Println(jsonContent)

	var p_jsonContent *string

	p_jsonContent = &jsonContent

	*p_jsonContent = removePOSTwords(jsonContent)

	//fmt.Println(jsonContent)

	//findJsonInText(jsonContent)
	urls := findURLsinText(jsonContent)

	*p_jsonContent = removeUrls(jsonContent)

	//fmt.Println(jsonContent)

	jsons := splitByEmptyNewline(jsonContent)

	/*jsonUrlMap := make(map[string]string)

	for i, _ := range urls {
		jsonUrlMap[urls[i]] = jsons[i]
	}*/

	var trimmedUrls []string
	for _, url := range urls {

		trimmedUrls = append(trimmedUrls, strings.TrimSpace(url))
	}

	for z, _ := range urls {
		sendPostRequest(trimmedUrls[z], jsons[z], strings.TrimSpace(token))
	}

	// Now we have to process the urls and jsons list

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
	// testing
	//cleaned := jsonContent
	return cleaned
}

// url regex could be: ^(?:https?:\/\/)
func findURLsinText(jsonContent string) []string {
	//fmt.Println(jsonContent)
	str := jsonContent

	re := regexp.MustCompile(`(\ https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)

	submatchall := re.FindAllString(str, -1)

	var urls []string
	for _, match := range submatchall {
		//fmt.Println(match)
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
	//func (m RawMessage) MarshalJSON() ([]byte, error)
	//rJson := strconv.Quote(json)
	var bJson = []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bJson))
	if err != nil {
		panic(err)
	}
	//req.Header.Set("X-Custom-Header", "Real-Custom-Shit")
	//Authorization of the API call
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
