package httpClientHelper

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// NewClient Get new HTTP-Client with seperate cookie jar
func NewClient() *http.Client {
	// Eigenes CookieJar erstellen (für die Session)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar, // Dem Client das CookieJar zuweisen
	}
	return client
}

// HttpGetRequest Generic HTTP-Client Get function
func HttpGetRequest(client *http.Client, url string) (*http.Response, error) {
	httpResult, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error while HTTP GET request on %s: %v", url, err)
	}
	return httpResult, nil
}

// HttpPostRequest Generic HTTP-Client Post function
func HttpPostRequest(client *http.Client, url string, formBody url.Values) (*http.Response, error) {
	bodyReader := strings.NewReader(formBody.Encode())
	httpResult, err := client.Post(url, "application/x-www-form-urlencoded", bodyReader)
	if err != nil {
		log.Fatalf("Error while HTTP POST request on %s: %v", url, err)
	}
	return httpResult, err
}
