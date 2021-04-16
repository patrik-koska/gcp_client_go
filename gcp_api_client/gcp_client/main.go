package main

import (
	"fmt"
	"os"

)
// STILL IN PROGRESS, NOT WORKING -- 2021/04/13 9:28 PM
func main() {

    computeUrl := "../urls/post_url_one.txt"
    firewallUrl := "../urls/post_url_two.txt"
    apiKey := "../credentials/api_key"
    var ComputeJsonStr = []byte(`open json file here`)

}

func sendPostRequest() string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//^- handle error here
	req.Header.Set("X-Custom-Header", "myvalue")
	//Authorization of the API call
	req.Header.Set("Authorization", "Bearer " + apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return ""
}

func getCliOpts() string {
	return ""
}

func readFileAsString(file string) (string,error) {
	content, err := ioutil.ReadFile("filepath here")
	if err != nil {
		return nil, err
	}
	text := string(content)
	return text, nil
}
