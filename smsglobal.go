package main

import (
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"io/ioutil"
	"errors"
)

const (
	SMSGLOBAL_API_URL = "http://www.smsglobal.com/http-api.php"
)

func SMSNotify(username, password, from, to, message string) (error) {
	values := url.Values{}
	values.Set("action", "sendsms")
	values.Set("user", username)
	values.Set("password", password)
	values.Set("from", from)
	values.Set("to", to)
	values.Set("text", message)

	reqBody := strings.NewReader(values.Encode())
	req, err := http.NewRequest("POST", SMSGLOBAL_API_URL, reqBody)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("PostRequest: Unable to send request")
	}

	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Recieved raw: %s\n", string(contents))
	resp.Body.Close()
	return nil
}