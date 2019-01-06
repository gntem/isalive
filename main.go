package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HealthResponse struct {
	Domain       string  `json:"domain"`
	Port         int16   `json:"port"`
	StatusCode   int16   `json:"status_code"`
	ResponseIP   string  `json:"response_ip"`
	ResponseCode int16   `json:"response_code"`
	ResponseTime float32 `json:"response_time"`
}

func main() {
	URLToCheckPtr := flag.String("url", "", "Define url to check")

	flag.Parse()

	filterURLRegex := regexp.MustCompile("^(?:https?:)?\\/\\/")

	parsedURL := filterURLRegex.ReplaceAllString(*URLToCheckPtr, "")

	reqURL := fmt.Sprintf("https://isitup.org/%s.json", parsedURL)

	res, err := http.Get(reqURL)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var rDataHealthRes HealthResponse

	json.Unmarshal(body, &rDataHealthRes)

	fmt.Printf("Domain ::\t\t %s \n", rDataHealthRes.Domain)
	fmt.Printf("Http Status ::\t\t %d \n", rDataHealthRes.StatusCode)
	fmt.Printf("Response code ::\t %d \n", rDataHealthRes.ResponseCode)
	fmt.Printf("IP:Port ::\t\t %s:%d \n", rDataHealthRes.ResponseIP, rDataHealthRes.Port)
	fmt.Printf("Time (ms) ::\t\t %f \n", rDataHealthRes.ResponseTime)
}
