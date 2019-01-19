package translator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

func test() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Dial: dialer.Dial,
	}
	c := http.Client{
		Transport: transport,
	}

	reqURL := fmt.Sprintf("http://translate.google.com/translate_a/single?client=gtx&sl=auto&tl=ja-jp&dt=t&q=%s", url.QueryEscape("who are you"))
	resp, err := c.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	jsonText := string(body)
	i := strings.IndexRune(jsonText, '"')
	j := strings.IndexRune(jsonText[i+1:], '"')

	if err == nil {
		fmt.Println(jsonText[i+1 : i+j+1])
	}
}
