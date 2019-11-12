package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/html/charset"
)

func main() {
	//var url = "http://reg3.sut.ac.th/registrar/calendar.asp?schedulegroupid=101&acadyear=2562&semester=1"
	// var url = "https://www.amazon.com/Animal-Farm-GEORGE-ORWELL/dp/9386538288/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=1573527970&sr=8-1"
	// var url = "https://www.amazon.com/Animal-Farm-Large-George-Orwell-dp-4871872696/dp/4871872696/ref=mt_paperback?_encoding=UTF8&me=&qid=1573546394"
	var url = "https://www.amazon.com/Harraps-Slovene-Phrasebook/dp/0071546111/ref=sr_1_1?crid=10BCROLSA6WY3&keywords=harraps+book&qid=1573550166&s=books&sprefix=harra%2Cstripbooks-intl-ship%2C388&sr=1-1"
	doc, err := Init(url)
	if err == nil {
		imageUrl, err := GetUrlImgage(doc)
		if err == nil {
			fmt.Println(imageUrl)
		}

		bookName, err := GetBookName(doc)
		if err == nil {
			fmt.Println(bookName)
		}

	}

	//fmt.Print(data)
}

func GetUrlImgage(doc *goquery.Document) (reulst string, err error) {
	//type aa []string
	// Find the review items
	var img_url string
	var firstTime bool = true
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, ok := s.Find("img").Attr("data-a-dynamic-image")
		if ok && firstTime {
			firstTime = false
			img_url = band[2:strings.Index(band, "\":")]
			// fmt.Println(i, img_url)
		} else {
			return
		}
	})
	return img_url, nil
	// doc.Find("span.a-declarative").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title

	// 	band := s.Find("a").Text()
	// 	if i == 9 {
	// 		fmt.Println(i, band)
	// 	}

	// })
}

func GetBookName(doc *goquery.Document) (reulst string, err error) {
	//type aa []string
	// Find the review items
	var book_name string
	// var firstTime bool = true
	doc.Find("h1#title").Each(func(i int, s *goquery.Selection) {
		band := s.Find("span").Text()
		// fmt.Println(i, band)
		if i == 0 {
			book_name = band
			return
		}

	})
	return book_name, nil

}

func Init(url string) (*goquery.Document, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := req.Header.Get("Content-Type") // Optional, better guessing
	utf8reader, err := charset.NewReader(resp.Body, contentType)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(utf8reader)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
