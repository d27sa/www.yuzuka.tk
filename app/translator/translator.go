package translator

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

const (
	// Google means use google translate
	Google = "google"

	// Baidu means use baidu translate
	Baidu = "baidu"

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

var bdLang = map[string]string{"ja": "jp", "zh-CN": "zh", "ko": "kor"}

var c *http.Client

var gtk string

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	c = &http.Client{Jar: jar}
	_, err = c.Get("https://fanyi.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
}

func parseToken(t string) string {
	i := strings.Index(t, "token:")
	if i < 0 {
		return ""
	}
	i += len("token:")
	j := strings.IndexRune(t[i:], '\'') + i
	if j <= i {
		return ""
	}
	j++
	l := strings.IndexRune(t[j:], '\'') + j
	if l <= j {
		return ""
	}
	return t[j:l]
}

func parseGTK(t string) string {
	i := strings.Index(t, "window.gtk")
	if i < 0 {
		return ""
	}
	j := i + len("window.gtk")
	for t[j] != '\'' {
		j++
	}
	j++
	l := strings.IndexRune(t[j:], '\'') + j
	if l <= j {
		return ""
	}
	return t[j:l]
}

// Translate translates the given text
func Translate(from, to, text string, engines []string) ([]byte, error) {
	var results []*result
	for _, e := range engines {
		switch e {
		case Google:
			results = append(results, newResult(Google, googleTranslate(from, to, text)))

		case Baidu:
			results = append(results, newResult(Baidu, baiduTranslate(from, to, text)))
		}
	}
	return json.Marshal(results)
}

// googleTranslate translates the text using google translate
func googleTranslate(from, to, text string) string {
	reqURL := fmt.Sprintf("http://translate.google.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", from, to, url.QueryEscape(text))
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

// baiduTranslate translates the text using baidu translate
func baiduTranslate(from, to, text string) string {
	if bdLang[from] != "" {
		from = bdLang[from]
	}
	if bdLang[to] != "" {
		to = bdLang[to]
	}
	res, err := c.Get("https://fanyi.baidu.com")
	if err != nil {
		log.Println(err)
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
		return ""
	}
	t := string(body)
	if gtk == "" {
		gtk = parseGTK(t)
	}
	token := parseToken(t)
	bdsign, err := sign(text, gtk)
	if err != nil {
		log.Println(err)
		return ""
	}
	queryStr := fmt.Sprintf("from=%s&to=%s&query=%s&sign=%s&token=%s", from, to, text, bdsign, token)
	postData := strings.NewReader(queryStr)
	res, err = c.Post("https://fanyi.baidu.com/v2transapi", "application/x-www-form-urlencoded", postData)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return parseBaiduResult(string(body))
}

func parseBaiduResult(t string) string {
	i := strings.Index(t, `"dst"`)
	if i < 0 {
		return ""
	}
	i += len(`"dst"`)
	j := strings.IndexRune(t[i:], '"') + i
	if j <= i {
		return ""
	}
	j++
	i = strings.IndexRune(t[j:], '"') + j
	if i <= j {
		return ""
	}
	tail := t[j:i]
	var (
		builder strings.Builder
		v       rune
		err     error
	)
	for tail != "" {
		v, _, tail, err = strconv.UnquoteChar(tail, tail[0])
		if err != nil {
			return ""
		}
		_, err = builder.WriteRune(v)
		if err != nil {
			return ""
		}
	}
	return builder.String()
}
