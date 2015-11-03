package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	apiURL     = "http://api.wallsreetcn.com/v2/admin/livenews?api_key=lQecdGAe"
	testapiURL = "http://api.wallstcn.com/v2/admin/livenews?api_key=lQecdGAe"
)

var proxy = []string{"http://123.59.83.131:23128", "http://123.59.83.140:23128", "http://123.59.83.137:23128",
	"http://123.59.83.147:23128", "http://123.59.83.139:23128", "http://123.59.83.132:23128",
	"http://123.59.83.141:23128", "http://123.59.83.148:23128", "http://123.59.83.130:23128", "http://123.59.83.145:23128"}

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

func (j Jin10) dealTime() (ts int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	now := time.Now()
	now = now.In(loc)
	return now.Unix(), nil
}

func (j *Jin10) matchResult() (content string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(j.jin10_page))
	if err != nil {
		return "", err
	}
	// firstEle, err := doc.Find("#newslist table").Eq(0).Html()
	// if err != nil {
	// 	return "", 0, err
	// }
	//contentExp, err := regexp.Compile("\\<tddddd align=\"left\" valign=\"middle\" id=\"content_[0-9]+\"\\>(.+)?\\</td\\>")
	//timeExp, err := regexp.Compile("\\<td align=\"left\" valign=\"middle\" width=\"55\"\\>(.+)?\\</td\\>")
	// if err != nil {
	// 	return "", 0, err
	// }
	ID, hasID := doc.Find("#newslist .newsline").Eq(0).Attr("id")
	if !hasID {
		log.Println(errors.New("content match failed."))
	}
	content = doc.Find("#content_" + ID).Text()
	if content == "" {
		c := doc.Find("#newslist .newsline").Eq(0).Find("table table tr").Text()
		if c == "" {
			content = ""
			err = nil
			return
		}
		re, err := regexp.Compile("\\s{2,}")
		if err != nil {
			return "", err
		}
		c = re.ReplaceAllString(c, " ")
		c = strings.TrimSpace(c)
		src := strings.Split(c, " ")
		content = src[1] + src[2] + "，" + src[4] + "，" + src[3]
	}
	keyword := []string{"jin10", "金十", "推荐阅读", "视频", "新品上线"}
	keywordSlice := make([]string, 0)
	for _, v := range keyword {
		rawBool := strings.Contains(content, v)
		keywordSlice = append(keywordSlice, strconv.FormatBool(rawBool))
	}
	hasKeyword := strings.Join(keywordSlice, ",")
	if hasK := strings.Contains(hasKeyword, "true"); hasK {
		content = ""
	}
	err = nil
	return
}

func (j *Jin10) getPage(proxy string) error {
	url := "http://jin10.com/"
	resp, err := j.getByProxy(url, proxy)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	j.jin10_page = string(body)
	return nil
}

func main() {
	var jin10 = &Jin10{}
	var end = make(chan bool)
	n := 0
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			timer := time.Now().Hour()
			if timer >= 0 && timer <= 6 {
				err := jin10.getPage(proxy[n])
				if err != nil {
					log.Println(err)
					return
				}
				content, err := jin10.matchResult()
				if err != nil {
					log.Println(err)
					return
				}
				ts, err := jin10.dealTime()
				if err != nil {
					log.Println(err)
					return
				}
				jin10.Content = content
				jin10.CodeType = "markdown"
				jin10.Channels = []int{1}
				jin10.CreateAt = ts
				p, err := json.Marshal(jin10)
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println(string(p))
				n++
				if n == 10 {
					n = 0
				}
				if content != "" {
					body := bytes.NewBuffer(p)
					res, err := http.Post(apiURL, "application/json;charset=utf-8", body)
					if err != nil {
						log.Println(err)
						return
					}
					msg, err := ioutil.ReadAll(res.Body)
					defer res.Body.Close()
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println(string(msg))
				}
			}
		}
	}()
	// time.Sleep(100 * time.Second)
	<-end
	fmt.Println("end")
}
