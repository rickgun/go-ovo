package ovo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	APIEnvType	Env
	AppID		string
	Key			string
	TID			string
	MID			string
	StoreCode	string
	MerchantID	string
	AppSource	string
}

func NewClient() Client {
	return Client{
		APIEnvType: Staging,
	}
}

// var defHTTPTimeout = 60 * time.Second
var httpClient = &http.Client{}

func (c *Client) Call(method, path string, headers []Header, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		fmt.Println("Request creation failed : ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	for _, header := range headers {
        req.Header[header.Key] = []string{header.Value}
    }

	start := time.Now()

	fmt.Println("Request ", method, ": ", path)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Cannot send request : ", err)
		return err
	}
	fmt.Println("Completed in ", time.Since(start))
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
