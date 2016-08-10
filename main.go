package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func main() {
	tbProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v\n", err)
	}

	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		log.Fatalf("Failed to obtain proxy dialer: %v\n", err)
	}

	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}

	resp, err := client.Get("http://ip2location.com")
	if err != nil {
		log.Fatalf("Failed to issue GET request: %v\n", err)
	}
	defer resp.Body.Close()

	log.Printf("GET returned: %v\n", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read the body: %v\n", err)
	}
	log.Printf("%s", body)
}
