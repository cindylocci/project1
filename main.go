package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//GENERATE API KEY - https://yandex.com/dev/dictionary/
//LOOK AT DOCUMENTATION - https://yandex.com/dev/dictionary/doc/dg/reference/lookup.html

var (
	apiKey     string                //generated yandex dict api key
	langParam  string = `lang=en-ru` //translation direction
	text       string                //text to translate
	requestUrl string                //full request url
)

type ResponseJsonStruct struct {
	Def []struct {
		Text string `json:"text"`
		Pos  string `json:"pos"`
		Ts   string `json:"ts"`
		Tr   []struct {
			Text string `json:"text"`
			Pos  string `json:"pos"`
			Gen  string `json:"gen,omitempty"`
			Fr   int    `json:"fr"`
			Syn  []struct {
				Text string `json:"text"`
				Pos  string `json:"pos"`
				Gen  string `json:"gen,omitempty"`
				Fr   int    `json:"fr"`
			} `json:"syn,omitempty"`
			Mean []struct {
				Text string `json:"text"`
			} `json:"mean"`
		} `json:"tr"`
	} `json:"def"`
}

func main() {

	fmt.Print("> Pass API Key > ")
	fmt.Scan(&apiKey)
	//apiKey = `  [TESTING API KEY]  `

	fmt.Print("> Enter the word you want to translate > ")
	fmt.Scan(&text)

	requestUrl = fmt.Sprintf("https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=%v&%v&text=%v", apiKey, langParam, text)

	sendGetRequest(requestUrl)

}

func sendGetRequest(url string) {

	var translated []string

	resp, err := http.Get(url) //Request url and get response

	if err != nil {
		log.Fatal("Error making http request", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) //Read response

	fmt.Println(strings.Repeat("\u25a7 ", 50))

	log.Println("Client: got response!")
	log.Println("Client: status code:", resp.Status)

	var HTTPresponse ResponseJsonStruct //Creat structure to save and parse json respone data

	err = json.Unmarshal(body, &HTTPresponse) //Unmarshal - parse json data
	if err != nil {
		log.Fatal("Error with parsing json", err)
	}

	for _, el := range HTTPresponse.Def { //Get translate fields of json and append to slice
		for _, tr := range el.Tr {
			translated = append(translated, tr.Text)
		}
	}

	fmt.Println(strings.Repeat("\u25a7 ", 50))
	fmt.Println("Got Translation for:", text, translated)
	fmt.Println(strings.Repeat("\u25a7 ", 50))

}
