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
	// var url = "https://www.amazon.com/Harraps-Slovene-Phrasebook/dp/0071546111/ref=sr_1_1?crid=10BCROLSA6WY3&keywords=harraps+book&qid=1573550166&s=books&sprefix=harra%2Cstripbooks-intl-ship%2C388&sr=1-1"
	// var url = "https://www.amazon.com/Dasd-Direct-Access-Storage-Devices/dp/0070326746/ref=sr_1_1?keywords=dasd&qid=1573554307&s=books&sr=1-1"
	// var url = "https://www.amazon.com/gp/product/1558329919/ref=s9_acsd_al_bw_c_x_3_w?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=merchandised-search-8&pf_rd_r=WND1J9YNKCC0099P8GVX&pf_rd_t=101&pf_rd_p=69e92e00-dd37-4469-9e77-9c4b5c15c5a4&pf_rd_i=283155"
	var url = "https://www.amazon.com/Wrecking-Ball-Diary-Wimpy-Book/dp/1419739034/ref=tmm_hrd_swatch_0?_encoding=UTF8&qid=&sr="
	// var url = "https://www.amazon.com/If-You-Tell-Unbreakable-Sisterhood/dp/1542005221/ref=tmm_hrd_swatch_0?_encoding=UTF8&qid=&sr="
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
		book_author, err := GetBookAuthor(doc)
		if err == nil {
			fmt.Println(book_author)
		}
		price, err := GetPrice(doc)
		if err == nil {
			fmt.Println(price)
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
}

func GetBookName(doc *goquery.Document) (reulst string, err error) {
	//type aa []string
	// Find the review items
	var book_name string
	// var firstTime bool = true
	doc.Find("h1#title").Each(func(i int, s *goquery.Selection) {
		band := s.Find("span#productTitle").Text()
		// fmt.Println(i, band)
		// if i == 0 {
		book_name = band
		return
		// }

	})
	return book_name, nil

}

func GetBookAuthor(doc *goquery.Document) (result string, err error) {
	var book_author string = ""
	doc.Find("span.author.notFaded").Each(func(j int, se *goquery.Selection) {
		// fmt.Println(se.Html())
		html, err := se.Html()
		if err != nil {
			return
		}
		if strings.Contains(html[0:100], "<a") {
			if se.Find("span").Find("span").Text() != "(Introduction)" {
				result = se.Find("a").Text()
				if book_author != "" && result != "" {
					book_author = book_author + ", "
				}
				// fmt.Println("In first Loop")
				book_author = book_author + result
				return
			}

		} else {
			// fmt.Println("In second Loop")
			se.Find("span").Each(func(i int, s *goquery.Selection) {
				if se.Find("span").Find("span").Text() != "(Introduction)" {
					result = s.Find("a").Text()
					if book_author != "" && result != "" {
						book_author = book_author + ", "
					}
					// fmt.Println(i, result)
					if i == 5 {
						book_author = book_author + result
						return
					}
				}
			})
		}

	})

	return book_author, nil
}

func GetPrice(doc *goquery.Document) (result string, err error) {
	var price string
	var isFirst bool = true
	doc.Find("span.a-color-price").Each(func(i int, s *goquery.Selection) {
		if isFirst {
			isFirst = false
			price = strings.TrimSpace(s.Text())
			return
		}
	})
	return price, nil
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
