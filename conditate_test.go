package nataelb_test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParseURL1(t *testing.T) {
	url, err := url.Parse("http://127.0.0.1:5001")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url)
}

func TestParseURL2(t *testing.T) {
	url, err := url.Parse("https://google.com:443")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url)
}
