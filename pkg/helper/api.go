package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gojek/heimdall/v7/httpclient"
	"golang.org/x/net/context"
)

type RequestLog struct {
	Url    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   interface{}       `json:"body"`
}

// APICall returns
type APICall struct {
	URL       string            `json:"url"`
	Method    string            `json:"method"`
	FormParam string            `json:"form_param"`
	Header    map[string]string `json:"header"`
}

// URLHttpResponse return
type URLHttpResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Header     http.Header `json:"header"`
}

// ConstructRequestLog
func ConstructRequestLog(url, method string, header map[string]string, body interface{}) string {
	reqLog := &RequestLog{
		Url:    url,
		Method: method,
		Header: header,
		Body:   body,
	}

	reqLogJson, err := json.Marshal(reqLog)
	if err != nil {
		return ""
	}

	return string(reqLogJson)
}

// ParseRequestBody
func ParseRequestBody(data interface{}) (string, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(body), err
}

// Call call to third party endpoint
func (ac *APICall) CallWithJson(ctx context.Context) (result URLHttpResponse, err error) {
	client := httpclient.NewClient()

	// Create an http.Request instance
	req, err := http.NewRequest(ac.Method, ac.URL, bytes.NewBuffer([]byte(ac.FormParam)))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)
	for index, value := range ac.Header {
		req.Header.Add(index, value)
	}

	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	result.Status = res.Status
	result.StatusCode = res.StatusCode
	result.Body = string(body)
	result.Header = res.Header

	return
}
