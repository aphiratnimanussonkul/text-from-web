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
	"time"

	"golang.org/x/net/html/charset"
)

func main() {
	//var data model.Datajsonn
	//var url = "http://reg3.sut.ac.th/registrar/calendar.asp?schedulegroupid=101&acadyear=2562&semester=1"
	var url = "https://www.amazon.com/Animal-Farm-GEORGE-ORWELL/dp/9386538288/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=1573527970&sr=8-1"
	exampleScrape(url)
	//fmt.Print(data)
}

func exampleScrape(url string) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.amazon.com/Animal-Farm-GEORGE-ORWELL/dp/9386538288/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=1573527970&sr=8-1", nil)
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

	// fmt.Println(string(body))
	defer timer()()

	// dt  ไลด์ว่างของ dayTime
	//var dt []model.DayTime
	//
	//var d model.DReg
	//var data model.Datajsonn

	// stetus 0 > เริ่มใหม่, 1 > ต่อ
	//var stetuss int
	var checkk int

	//stetuss = 0
	//checkk := 00
	res, err := http.Get(url)
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("status code error: %d %s 01 - ", res.StatusCode, res.Status)
	}

	contentType := res.Header.Get("Content-Type") // Optional, better guessing
	utf8reader, err := charset.NewReader(resp.Body, contentType)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(utf8reader)
	if err != nil {
		fmt.Println("02", err)
	}
	fmt.Println("Doc = ", doc)

	checkk = 99
	fmt.Print("Checkk = ", checkk)
	//type aa []string
	// Find the review items
	doc.Find("div#img-canvas").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title

		band, ok := s.Find("img").Attr("src")
		
		fmt.Println(i, band, ok)
	})
	//
	//	if i == 6 {
	//		c := chunkString(band, 6)
	//		data.Coursecode = c[0]
	//		data.NameEnglish = c[1]
	//	}
	//	if i == 7 {
	//		data.NameThai = band
	//
	//	}
	//	if i == 9 {
	//		data.Credit = band
	//	}
	//
	//	if checkk == i {
	//		//fmt.Println(i)
	//
	//		stetuss = 1
	//	}
	//	if strings.Contains(band, "กลุ่มวันเวลาห้องอาคารเรียนที่นั่ง(เปิด-ลง-เหลือ)หมวด") && i != 0 {
	//		checkk = i + 1
	//
	//	}
	//	if stetuss == 2 {
	//		if strings.Contains(band, "หมายเหตุ") {
	//			s := strings.Split(band, "หมายเหตุ:")
	//			t := strings.TrimSpace(s[1])
	//			d.Note = t
	//		}
	//		data.Datareg = append(data.Datareg, d)
	//		d.T = ""
	//		d.Mid = model.Date{}
	//		d.Final = model.Date{}
	//		d.Note = ""
	//		d.DayTime = dt
	//		stetuss = 1
	//	}
	//
	//	if stetuss == 1 {
	//
	//		//fmt.Println("two", i)
	//		if strings.Contains(band, "อาจารย์") {
	//			if strings.Contains(band, "หมายเหตุ") {
	//
	//			} else {
	//				s := strings.Split(band, "อาจารย์:")
	//				t := strings.TrimSpace(s[1])
	//
	//				d.T = t
	//				return
	//			}
	//
	//		} else if strings.Contains(band, "สอบกลางภาค") {
	//			s := strings.Split(band, "สอบกลางภาค:")
	//			t := strings.TrimSpace(s[1])
	//			if t == "" {
	//				d.Final = model.Date{"", ""}
	//				return
	//			}
	//			x := strings.Split(t, " อาคาร")
	//			date := strings.Split(x[0], " เวลา")
	//			time := strings.TrimSpace(date[1])
	//			d.Mid = model.Date{date[0], time}
	//		} else if strings.Contains(band, "สอบประจำภาค:") {
	//			s := strings.Split(band, "สอบประจำภาค:")
	//			t := strings.TrimSpace(s[1])
	//			if t == "" {
	//				d.Final = model.Date{"", ""}
	//				stetuss = 2
	//				return
	//			}
	//			x := strings.Split(t, " อาคาร")
	//			date := strings.Split(x[0], " เวลา")
	//			time := strings.TrimSpace(date[1])
	//			d.Final = model.Date{date[0], time}
	//			stetuss = 2
	//		}
	//
	//		//--------------------------------
	//		//--------------------------------
	//		if strings.Contains(band, "หมายเหตุ") {
	//			return
	//		}
	//
	//		if strings.Contains(band, "จันทร์") {
	//			s := strings.Split(band, "จันทร์")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "จันทร์")
	//			day := "จันทร์"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//
	//		} else if strings.Contains(band, "อังคาร") {
	//			s := strings.Split(band, "อังคาร")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "อังคาร")
	//			day := "อังคาร"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//		} else if strings.Contains(band, "พุธ") {
	//			s := strings.Split(band, "พุธ")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "พุธ")
	//			day := "พุธ"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//		} else if strings.Contains(band, "พฤหัสบดี") {
	//			s := strings.Split(band, "พฤหัสบดี")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "พฤหัสบดี")
	//			day := "พฤหัสบดี"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//		} else if strings.Contains(band, "ศุกร์") {
	//			s := strings.Split(band, "ศุกร์")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "ศุกร์")
	//			day := "ศุกร์"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//		} else if strings.Contains(band, "เสาร์") {
	//			s := strings.Split(band, "เสาร์")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "เสาร์")
	//			day := "เสาร์"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//
	//		} else if strings.Contains(band, "อาทิตย์") {
	//			s := strings.Split(band, "อาทิตย์")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			sum := cut(band, s[0])
	//			time := strings.Split(sum, "อาทิตย์")
	//			day := "อาทิตย์"
	//
	//			d.DayTime = append(d.DayTime, model.DayTime{day, time[1]})
	//		} else if strings.Contains(band, "ไม่มีข้อมูล") {
	//			s := strings.Split(band, "ไม่มีข้อมูล")
	//			t := strings.TrimSpace(s[0])
	//			if t != "" {
	//				d.Sec = t
	//			}
	//			d.DayTime = append(d.DayTime, model.DayTime{"ไม่มีข้อมูล", "ไม่มีข้อมูล"})
	//
	//		}
	//	}
	//})

	//return data
}

func cut(band string, s string) string {
	building := []string{"F", "B", "L", "M", "S", "N", "A", "E", "G", "C", "ห้อง", "สนาม", "อาคาร"}

	for _, b := range building {
		v := strings.Split(band, b)
		band = v[0]
	}
	v2 := strings.Split(band, s)
	return v2[1]
}

func ch(band string) bool {
	if strings.Contains(band, "จันทร์") {
		return true
	} else if strings.Contains(band, "อังคาร") {
		return true
	}
	return false
}

// GetDataReg is schedule data from reg
//func GetDataReg(cid string, c echo.Context, semester string, acadyear string) error {
//	//defer deleteDatax()
//	data := exampleScrape("http://reg2.sut.ac.th/registrar/class_info_2.asp?backto=home&option=0&courseid=" + cid + "&acadyear=" + acadyear + "&semester=" + semester + "&avs972184082=6")
//	//fmt.Println("#!", data)
//	fmt.Print(data)
//	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
//	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
//	c.Response().WriteHeader(http.StatusOK)
//	return json.NewEncoder(c.Response()).Encode(data)
//
//}

func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		log.Println(diff)
	}
}

//func chunkString(s string, chunkSize int) []string {
//	var chunks []string
//	runes := []rune(s)
//
//	if len(runes) == 0 {
//		return []string{s}
//	}
//
//	for i := 0; i < len(runes); i += chunkSize {
//		nn := i + chunkSize
//		if nn > len(runes) {
//			nn = len(runes)
//		}
//		chunks = append(chunks, string(runes[i:nn]), string(runes[nn:]))
//		break
//	}
//	return chunks
//}
