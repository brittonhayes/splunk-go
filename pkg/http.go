package splunk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
 * HTTP helper methods
 */

//  nosec
func httpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}
	return client
}

func (conn Connection) httpGet(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "GET", data)
}

func (conn Connection) httpPost(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "POST", data)
}

func (conn Connection) httpCall(url string, method string, data *url.Values) (string, error) {
	client := httpClient()

	var payload io.Reader
	if data != nil {
		payload = bytes.NewBufferString(data.Encode())
	}

	request, err := http.NewRequest("POST", url, payload)
	conn.addAuthHeader(request)

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()
	return string(body), nil
}

func (conn Connection) addAuthHeader(request *http.Request) {
	if conn.sessionKey.Value != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", conn.sessionKey.Value))
	} else {
		request.SetBasicAuth(conn.Username, conn.Password)
	}
}
