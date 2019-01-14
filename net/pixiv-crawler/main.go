package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/proxy"
)

// PixivSession represents a session used to access pixiv.net
type PixivSession struct {
	*http.Client
}

// NewPixivSession returns a pointer to a newly created PixivSession
func NewPixivSession() *PixivSession {
	return &PixivSession{
		Client: &http.Client{},
	}
}

// SetSocks5Proxy sets a socks5 proxy for the session
func (psess *PixivSession) SetSocks5Proxy(addr string) error {
	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		return err
	}
	psess.Transport = &http.Transport{Dial: dialer.Dial}
	return nil
}

// InitCookieJar initials the cookiejar of the session
func (psess *PixivSession) InitCookieJar() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	psess.Jar = jar
	return nil
}

// Login logs in PIXIV
func (psess *PixivSession) Login(id string, password string) error {
	psess.Get("https://www.pixiv.net/")
	postKey, err := psess.getPostKey()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	buf.WriteString("pixiv_id=")
	buf.WriteString(id)
	buf.WriteString("&password=")
	buf.WriteString(password)
	buf.WriteString("&post_key=")
	buf.WriteString(postKey)
	_, err = psess.Post("https://accounts.pixiv.net/api/login?lang=ja", "application/x-www-form-urlencoded", bytes.NewReader(buf.Bytes()))
	return err
}

// getPostKey gets the post_key for login
func (psess PixivSession) getPostKey() (string, error) {
	resp, err := psess.Get("https://accounts.pixiv.net/login?lang=ja&source=pc&view_type=page&ref=wwwtop_accounts_index")
	if err != nil {
		return "", err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}
	var postKeyNode *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		for _, a := range n.Attr {
			if a.Val == "post_key" {
				postKeyNode = n
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if postKeyNode == nil {
		return "", errors.New("cannot find post_key node")
	}
	postKey := ""
	for _, a := range postKeyNode.Attr {
		if a.Key == "value" {
			postKey = a.Val
			break
		}
	}
	if postKey == "" {
		return postKey, errors.New("cannot find post_key value")
	}
	return postKey, nil
}

// GetIllustByID downloads the illust of specified ID
func (psess PixivSession) GetIllustByID(id int, dir, filename string) error {
	var builder strings.Builder
	builder.WriteString("https://www.pixiv.net/member_illust.php?mode=medium&illust_id=")
	idStr := strconv.Itoa(id)
	builder.WriteString(idStr)
	reqURL := builder.String()
	resp, err := psess.Get(reqURL)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	text := string(body)
	i := strings.Index(text, `"original":`)
	if i < 0 {
		return errors.New("cannot find image url")
	}
	text = text[i+len(`"original":`):]
	i = strings.Index(text, `"`)
	if i < 0 {
		return errors.New("cannot find image url")
	}
	text = text[i+1:]
	j := strings.Index(text, `"`)
	if j < 0 {
		return errors.New("cannot find image url")
	}
	imgURL := text[:j]
	imgURL = strings.Replace(imgURL, "\\", "", -1)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("GET", imgURL, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Referer", reqURL)
	resp, err = psess.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ext := filepath.Ext(imgURL)
	builder.Reset()
	if filename == "" {
		builder.WriteString(idStr)
	} else {
		builder.WriteString(filename)
	}
	builder.WriteString(ext)
	path := filepath.Join(dir, builder.String())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(img)
	return err
}

func main() {
	psess := NewPixivSession()
	err := psess.InitCookieJar()
	if err != nil {
		log.Fatal(err)
	}
	err = psess.SetSocks5Proxy("127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	err = psess.Login("d27sa", "zhukm1997")
	if err != nil {
		log.Fatal(err)
	}
	err = psess.GetIllustByID(72597036, "", "")
	if err != nil {
		log.Fatal(err)
	}
}
