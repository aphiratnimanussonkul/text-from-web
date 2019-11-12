package main

import (
	//"encoding/json"
	"fmt"
	// "io/ioutil"
	
	"github.com/PuerkitoBio/goquery"

	//"github.com/fooku/sutRegApi/pkg/model"
	//"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

func main() {
	//var data model.Datajsonn
	//var url = "http://reg3.sut.ac.th/registrar/calendar.asp?schedulegroupid=101&acadyear=2562&semester=1"
	// var url = "https://www.amazon.com/Animal-Farm-GEORGE-ORWELL/dp/9386538288/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=1573527970&sr=8-1"
	var url = "https://www.amazon.com/Animal-Farm-Large-George-Orwell-dp-4871872696/dp/4871872696/ref=mt_paperback?_encoding=UTF8&me=&qid=1573546394"
	fmt.Println(exampleScrape(url))
	//fmt.Print(data)
}

func exampleScrape(url string) string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type") // Optional, better guessing
	utf8reader, err := charset.NewReader(resp.Body, contentType)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(utf8reader)
	if err != nil {
	}

	//type aa []string
	// Find the review items
	var img_url string
	var firstTime bool = true
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, ok := s.Find("img").Attr("data-a-dynamic-image")
		if ok && firstTime {
			firstTime = false
			img_url = band[2:strings.Index(band,"\":")]
			// fmt.Println(i, img_url)
		} else {
			return
		}

	})
	return img_url
	// doc.Find("span.a-declarative").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title

	// 	band := s.Find("a").Text()
	// 	if i == 9 {
	// 		fmt.Println(i, band)
	// 	}

	// })
}