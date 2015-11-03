package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	apiURL     = "http://api.wallsreetcn.com/v2/admin/livenews?api_key=lQecdGAe"
	testapiURL = "http://api.wallstcn.com/v2/admin/livenews?api_key=lQecdGAe"
)

type Jin10 struct {
	jin10_page string `json:"-"`
	CodeType   string `json:"codeType"`
	CreateAt   int64  `json:"createAt"`
	Channels   []int  `json:"channels"`
	Content    string `json:"content"`
}

func (j Jin10) getByProxy(url_addr, proxy_addr string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url_addr, nil)
	if err != nil {
		return nil, err
	}
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	return client.Do(request)
}

func (j *Jin10) getPage() {
	proxy := "http://123.59.83.131:23128"
	url := "http://jin10.com/"
	resp, err := j.getByProxy(url, proxy)
	if err != nil {
		log.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	j.jin10_page = string(body)
	_, _, err = j.matchResult()
	if err != nil {
		log.Println(err)
		return
	}
}

func (j Jin10) dealTime(t string) (ts int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	now := time.Now()
	now = now.In(loc)
	return now.Unix(), nil
}

func (j *Jin10) matchResult() (content string, ts int64, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(j.jin10_page))
	if err != nil {
		return "", 0, err
	}
	firstEle, err := doc.Find("#newslist table").Eq(44).Html()
	if err != nil {
		return "", 0, err
	}
	contentExp, err := regexp.Compile("\\<tddddd align=\"left\" valign=\"middle\" id=\"content_[0-9]+\"\\>(.+)?\\</td\\>")
	timeExp, err := regexp.Compile("\\<td align=\"left\" valign=\"middle\" width=\"55\"\\>(.+)?\\</td\\>")
	if err != nil {
		return "", 0, err
	}
	f := contentExp.FindStringSubmatch(firstEle)
	t := timeExp.FindStringSubmatch(firstEle)
	if len(f) == 0 {
		ID, hasID := doc.Find("#newslist .newsline").Eq(44).Attr("id")
		if !hasID {
			errors.New("dom match failed.")
		}
		content = doc.Find("#content_" + ID).Text()
		if content == "" {
			c := doc.Find("#newslist .newsline").Eq(44).Find("table table tr").Text()
			re, _ := regexp.Compile("\\s{2,}")
			src := strings.Split(re.ReplaceAllString(c, " "), " ")
			content = src[2] + src[3] + "," + src[4] + "," + src[5]
		}
	} else {
		content = f[1]
	}
	ts, err = j.dealTime(t[1])
	if err != nil {
		log.Println(err)
		return "", 0, err
	}
	fmt.Println(content, ts)
	return
}

func main() {
	var j = &Jin10{}
	j.getPage()
}
