package util

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SendHTTPRequest(method, contentType, endpoint string, token *string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, endpoint, body)
	req.Header.Set("Content-Type", contentType)
	if token != nil {
		req.Header.Set("Authorization", *token)
	}
	log.Println("Endpoint", endpoint)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Print("response Status:", resp.Status)

	//log.Print("Authorization", *token)
	//log.Print("Content-Type", contentType)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Println("req:", contentType)
	resBody, err := ioutil.ReadAll(resp.Body)

	return resBody, err
}

func Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
