package main

import (
	// "compress/gzip"
	"encoding/json"
	"fmt"
	// "time"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2/bson"
)

//www.investing.com
type Investing struct {
	Html string
}

type Investing_Filed struct {
	Price  string `json:"price"`
	Prev   string `json:"prev"`
	Open   string `json:"open"`
	RangeL string `json:"rangeL"`
	RangeR string `json:"rangeR"`
	Diff   string `json:"diff"`
	DiffP  string `json:"diffP"`
}

type FiledsSlice struct {
	Investing_Fileds []Investing_Filed
}

type GetValue interface {
	responseHandler(w http.ResponseWriter, r *http.Request)
	matchResult(page *string, regExp string) (field string)
}

func (i *Investing) decodeBson(data []byte) {
	// var v Investing
	err := bson.Unmarshal(data, &i)
	if err != nil {
		log.Println(err)
		return
	}
}

func (i *Investing) getField(regexp []string, page *string) (price, prev, open, rangeL, rangeR, diff, diffP string) {
	var field = make([]string, 0)
	for _, v := range regexp {
		f := i.matchResult(page, v)
		fmt.Println(f)
		field = append(field, f)
	}
	price = field[0]
	prev = field[1]
	open = field[2]
	rangeL = field[3]
	rangeR = field[4]
	diff = field[5]
	diffP = field[6]
	if price == "" {
		price = i.getFieldPrice()
	}
	if open == "" {
		open = i.getFieldOpen()
	}
	if prev == "" {
		prev = i.getFieldPrev()
	}
	if rangeL == "" || rangeR == "" {
		rangeL, rangeR = i.getFields()
	}
	if diff == "" {
		diff = i.getFieldDiff()
	}
	if diffP == "" {
		diffP = i.getFieldDiffP()
	}
	return
}

func (i *Investing) responseHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header.Get("Accept-Encoding"))
	//reader, err := gzip.NewReader(r.Body)
	result, err := ioutil.ReadAll(r.Body)
	if len(result) == 0 {
		log.Println("no request body")
		return
	}
	defer r.Body.Close()
	if err != nil {
		log.Println("http get error.")
		return
	}
	i.decodeBson(result)
	var regexpSli = []string{".+? id=\"last_last\" dir=\"ltr\">(.+)?</span>",
		".+?Prev. Close:</span> <span dir=\"ltr\">(.+)?</span>",
		".+?Open:</span> <span dir=\"ltr\">(.+)?</span>",
		".+?Day's Range:</span> <span dir=\"ltr\">(.+)? - .+?</span>",
		".+?Day's Range:</span> <span dir=\"ltr\">.+? - (.+)?</span>",
		"<span class=\"arial_20[\\s\\S]+pid-[0-9]+-pc\" dir=\"ltr\">(.+)?</span>",
		"<span class=\"arial_20[\\S\\s]+pid-[0-9]+-pcp parentheses\" dir=\"ltr\">(.+)?%</span>"}
	//var regexpSli = []string{".+? id=\"last_last\" dir=\"ltr\">(.+)?</span>.*"}
	//price, prev, open, rangeL, rangeR, diff, diffP := <-c, <-c, <-c, <-c, <-c, <-c, <-c
	price, prev, open, rangeL, rangeR, diff, diffP := i.getField(regexpSli, &i.Html)
	var fie FiledsSlice
	fie.Investing_Fileds = append(fie.Investing_Fileds, Investing_Filed{price, prev, open, rangeL, rangeR, diff, diffP})
	p, err := json.Marshal(fie)
	if err != nil {
		log.Println("json encode error.")
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(p))
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Printf("WARN:panic in %v - %v", fn, e)
			}
		}()
		fn(w, r)
	}
}

func (i *Investing) matchResult(page *string, regExp string) (field string) {
	//pattern_Prev, err := regexp.Compile(".+?Prev. Close:</span> <span dir=\"ltr\">(.+?)</span>")
	pattern, err := regexp.Compile(regExp)
	if err != nil {
		log.Println("RegExp compile failed.")
		return
	}
	res := pattern.FindStringSubmatch(*page)
	if len(res) == 0 {
		log.Println("RegExp match failed.")
		return ""
	}
	field = res[1]
	return
}

func (i *Investing) NewGoqueryDoc() (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(i.Html))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return doc, nil
}

func (i *Investing) getFieldPrice() (price string) {
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	price = doc.Find("#last_last").Text()
	return
}

func (i *Investing) getFieldOpen() (open string) {
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	res := doc.Find(".overviewDataTable").Children().Eq(3).Children().Eq(0).Text()
	openPrice := doc.Find(".overviewDataTable").Children().Eq(4).Children().Eq(0).Text()
	if res == "Open" {
		open = doc.Find(".overviewDataTable").Children().Eq(3).Children().Eq(1).Text()
		return
	} else if openPrice == "Price Open" {
		open = doc.Find(".overviewDataTable").Children().Eq(4).Children().Eq(1).Text()
		return
	} else {
		return ""
	}
}

func (i *Investing) getFieldPrev() (prev string) {
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	res := doc.Find(".overviewDataTable").Children().Eq(0).Children().Eq(0).Text()
	if res == "Prev. Close" {
		prev = doc.Find(".overviewDataTable").Children().Eq(0).Children().Eq(1).Text()
		return
	} else {
		return ""
	}
}

func (i *Investing) getFields() (rangeL, rangeR string) {
	var f []string
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	doc.Find("#quotes_summary_secondary_data").Children().Each(func(i int, s *goquery.Selection) {
		data := s.Find("span[dir=ltr]").Text()
		f = append(f, data)
	})
	rangeLR := strings.Split(f[2], " - ")
	rangeL = rangeLR[0]
	rangeR = rangeLR[1]
	return
}

func (i *Investing) getFieldDiff() (diff string) {
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	diff = doc.Find("#last_last").Siblings().Eq(0).Text()
	return
}

func (i *Investing) getFieldDiffP() (diffP string) {
	doc, err := i.NewGoqueryDoc()
	if err != nil {
		log.Println(err)
		return
	}
	data := doc.Find("#last_last").Siblings().Eq(2).Text()
	diffP = strings.Split(data, "%")[0]
	return
}

func main() {
	var i Investing
	http.HandleFunc("/", safeHandler(i.responseHandler))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
