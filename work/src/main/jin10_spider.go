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

	"github.com/PuerkitoBio/goquery"
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
	err = j.matchResult()
	if err != nil {
		log.Println(err)
		return
	}
}

func (j *Jin10) matchResult() error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(j.jin10_page))
	if err != nil {
		return err
	}
	firstEle, err := doc.Find("#newslist table").Eq(0).Html()
	if err != nil {
		return err
	}
	regExp, err := regexp.Compile("\\<td align=\"left\" valign=\"middle\" id=\"content_[0-9]+\"\\>(.+)?\\</td\\>")
	if err != nil {
		return err
	}
	var result string
	f := regExp.FindStringSubmatch(firstEle)
	if len(f) == 0 {
		ID, hasID := doc.Find("#newslist .newsline").Eq(0).Attr("id")
		if !hasID {
			errors.New("dom match failed.")
		}
		result = doc.Find("#content_" + ID).Text()
	} else {
		result = f[1]
	}
	fmt.Println(result)
	return nil
}

func main() {
	var j = &Jin10{}
	j.getPage()
}
