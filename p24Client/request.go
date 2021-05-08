package p24Client

import (
	"bytes"
	"encoding/json"
	"fmt"
	p24Error "github.com/arturwwl/p24-golang-client/model/error"
	"io/ioutil"
	"net/http"
)

const prodUrl = "https://secure.przelewy24.pl"
const sandboxUrl = "https://sandbox.przelewy24.pl"

func getUrl(isProd bool, path string) string {
	var urlString string
	if isProd {
		urlString = prodUrl
	} else {
		urlString = sandboxUrl
	}
	urlString += path
	return urlString
}

func (c *P24Client) MakeRequest(method string, path string, data interface{}) ([]byte, error) {
	var requestBody *bytes.Buffer
	var err error
	var req *http.Request

	if data != nil {
		postBody, _ := json.Marshal(data)
		requestBody = bytes.NewBuffer(postBody)
	}
	if requestBody == nil {
		req, err = http.NewRequest(method, getUrl(c.Config.IsProd, path), nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, getUrl(c.Config.IsProd, path), requestBody)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(fmt.Sprintf("%d", c.Config.PosID), c.Config.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 204 {
		var errO p24Error.ErrorStruct
		_ = json.Unmarshal(bodyBytes, &errO)
		return nil, fmt.Errorf(errO.Error)
	}

	return bodyBytes, err
}
