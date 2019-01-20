package translator

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Google means use google translate
	Google = 0

	// Baidu means use baidu translate
	Baidu = 1

	// English language
	English = "en"

	// Japanese language
	Japanese = "ja"

	// Chinese simplified
	Chinese = "zh-CN"

	// Korean language
	Korean = "ko"

	// Auto detect language
	Auto = "auto"
)

// Translator represents a translator app
type Translator struct {
}

// New creates a new Translator and returns a pointer to it
func New() *Translator {
	return &Translator{}
}

// Translate translates the specified text
func Translate(engine int, from, to, text string) string {
	switch engine {
	case Google:
		return googleTranslate(from, to, text)
	default:
		return ""
	}
}

func googleTranslate(from, to, text string) string {
	reqURL := fmt.Sprintf("http://translate.google.cn/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", from, to, url.QueryEscape(text))
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return ""
	}
	jsonText := string(body)
	i := strings.IndexRune(jsonText, '"')
	j := strings.IndexRune(jsonText[i+1:], '"')
	if i < 0 || j < 0 {
		return ""
	}
	return jsonText[i+1 : i+j+1]
}
