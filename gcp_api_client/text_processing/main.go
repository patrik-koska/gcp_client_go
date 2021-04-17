package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"regexp"
	
)

func main() {
	jsonFilePath := "/home/kali/projects/gcp_client_go/gcp_api_client/gcp_client/basic-gcp-server.json"
	
	jsonContent , err := readFileAsString(jsonFilePath)
	if err != nil {
		panic(err)
	}

	var p_jsonContent *string

	p_jsonContent = &jsonContent

	*p_jsonContent = removePOSTwords(jsonContent)

	//fmt.Println(jsonContent)

	//findJsonInText(jsonContent)
	urls := findURLsinText(jsonContent)


	*p_jsonContent = removeUrls(jsonContent)

	//fmt.Println(jsonContent)

	jsons := splitByEmptyNewline(jsonContent)

	fmt.Println("urls:")
	fmt.Println(urls[0])
	fmt.Println(urls[1])
	
	fmt.Println("jsons:")
	fmt.Println(jsons[0])
	fmt.Println(jsons[1])
	
}

func readFileAsString(filePath string) (string,error) {
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

func sendPostRequest(url string, json string) string {
	return ""
}

func getApiKey(keyPath string) (string, error) {
	content, err := ioutil.ReadFile(keyPath)
	if err != nil {
			return "", err
	}
	key := string(content)
	return key, nil
}