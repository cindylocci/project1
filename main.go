package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RespJson struct {
	Args struct {
	} `json:"args"`
	Data  string `json:"data"`
	Files struct {
	} `json:"files"`
	Form struct {
	} `json:"form"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	JSON   any    `json:"json"`
	Method string `json:"method"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {

	//sendGetRequest("https://httpbin.org/anything")
	sendPostRequest("https://httpbin.org/anything")

}

// https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go
func sendGetRequest(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Printf("Client: got response!\n")
	fmt.Printf("Client: status code: %d\n", resp.StatusCode)
	fmt.Printf("Response body raw bytes: %v\n", body)
	fmt.Printf("Response body string: %v\n", string(body))

	//Json parsing and transfer to struct
	//create struct object to save json data to
	var HTTPresponse RespJson

	//parse json data
	err = json.Unmarshal(body, &HTTPresponse)
	if err != nil {
		log.Fatal("Error with parsing json", err)
	}

	//Json is structured
	fmt.Println(HTTPresponse.Method)
}

func sendPostRequest(url string) {

	bodyBytes := []byte("testing:test")

	resp, err := http.Post(url, "text/html", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Printf("Client: got response!\n")
	fmt.Printf("Client: status code: %d\n", resp.StatusCode)
	fmt.Printf("Response body raw bytes: %v\n", body)
	fmt.Printf("Response body string: %v\n", string(body))

	// Json parsing and transfer to struct
	// create struct object to save json data to
	var HTTPresponse RespJson

	//parse json data
	err = json.Unmarshal(body, &HTTPresponse)
	if err != nil {
		log.Fatal("Error with parsing json", err)
	}

	//Json is structured - POST
	fmt.Println(HTTPresponse.Method)
}
