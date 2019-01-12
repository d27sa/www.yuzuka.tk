package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := &http.Client{}
	// rc, e := http.NewRequest("CONNECT", "http://accounts.pixiv.net:443", nil)
	// rc.Header.Set("Host", "accounts.pixiv.net:443")
	// rc.Header.Set("Proxy-Connection", "keep-alive")
	// rc.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")

	// r, e := http.NewRequest("POST", "https://accounts.pixiv.net/api/login?lang=ja", strings.NewReader("pixiv_id=d27sa&captcha=&g_recaptcha_response=&password=zhukm1997&post_key=6bcffa2f3efe0f26782a6740109a7c95&source=pc&ref=wwwtop_accounts_index&return_to=https%3A%2F%2Fwww.pixiv.net%2F"))
	// fmt.Println(e)
	// r.Header.Set("Host", "accounts.pixiv.net")
	// r.Header.Set("Connection", "keep-alive")
	// r.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// r.Header.Set("Content-Length", "183")
	// r.Header.Set("Connection", "keep-alive")
	// r.Header.Set("accept", "application/json")
	// r.Header.Set("Origin", "https://accounts.pixiv.net")
	// r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	// r.Header.Set("content-type", "application/x-www-form-urlencoded")
	// r.Header.Set("Referer", "https://accounts.pixiv.net/login?lang=ja&source=pc&view_type=page&ref=wwwtop_accounts_index")
	// // r.Header.Set("Cookie", "p_ab_id=6; p_ab_id_2=6; p_ab_d_id=577343894; __utmz=235335808.1542778747.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); _ga=GA1.2.400094805.1542778747; _ga=GA1.3.400094805.1542778747; a_type=0; module_orders_mypage=%5B%7B%22name%22%3A%22sketch_live%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22tag_follow%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22recommended_illusts%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22everyone_new_illusts%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22following_new_illusts%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22mypixiv_new_illusts%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22spotlight%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22fanbox%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22featured_tags%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22contests%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22user_events%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22sensei_courses%22%2C%22visible%22%3Atrue%7D%2C%7B%22name%22%3A%22booth_follow_items%22%2C%22visible%22%3Atrue%7D%5D; privacy_policy_agreement=1; d_type=1; __gads=ID=44da4ceffe1f97aa:T=1542779043:S=ALNI_MZZQEEBob1bhm6ximNpzDPhfNtm2w; __utmc=235335808; device_token=6c07812dc9aa49b03b6e160a48339f68; c_type=unknown; b_type=0; __utmv=235335808.|2=login%20ever=yes=1^3=plan=normal=1^5=gender=male=1^6=user_id=36382838=1^9=p_ab_id=6=1^10=p_ab_id_2=6=1^11=lang=ja=1; tag_view_ranking=BU9SQkS-zU~y8GNntYHsi~gYs_SnsOIO~gooMLQqB9a~P6BOd27mH2~uusOs0ipBx~DZrrbq7Psc~Q4duCCWLbW~QmdcEUDSz4; __utma=235335808.400094805.1542778747.1547112127.1547220659.4; __utmt=1; PHPSESSID=9e117084e0ef12c1869eb49ff38f91de; __utmb=235335808.5.9.1547220690501; login_bc=1; _gid=GA1.2.308071380.1547220703; _gat=1; _gid=GA1.3.308071380.1547220703; _gat_UA-76252338-4=1")

	resp, e := c.Do(rc)
	fmt.Println(e)
	if e == nil {
		defer resp.Body.Close()
		fmt.Println(e, resp.Status)
	}
}
