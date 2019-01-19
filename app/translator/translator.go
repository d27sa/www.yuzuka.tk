package translator

import (
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
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Panicln(err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
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
