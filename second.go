package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io/ioutil"
	"net/url"
	"os"
	//"path/filepath"
	"strings"
)

var stock = []string{}
var notfounds = []string{}
var notfoundImgPath = "/img/pc/common/error/imgNotFond.png"
var base = "https://kirei-kaigi.jp/"
var notfoundLog = "~/dev/go-crawl/notfound.log"
var urlLog = "~/dev/go-crawl/url.log"

func main() {
	var test = []string{"test", "kuwako"}
	writeFile(notfoundLog, test)
	//	result := makeUrl(base)
	//	results := getUrl(result)
	//
	//	//	for len(results) > 0 {
	//	//		results = getUrl(results)
	//	//	}
	//
	//	fmt.Println(results)
	//	fmt.Println("end_of_for")
	//	fmt.Println(notfounds)
	//	writeFile(urlLog, stock)
	//	writeFile(notfoundLog, notfounds)
}

// urlを渡したら、そこに含まれるurlのリストを取得する
func makeUrl(base string) []*url.URL {
	doc, _ := goquery.NewDocument(base)
	var result []*url.URL

	checkNotFound(doc)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		target, _ := s.Attr("href")
		base, _ := url.Parse(base)
		targets, _ := url.Parse(target)
		result = append(result, base.ResolveReference(targets))
	})

	return result
}

func getUrl(urls []*url.URL) []*url.URL {
	sourceUrl := []*url.URL{}
L:
	for _, item := range urls {
		url_string := item.String()

		// すでにstock済みのURLなら弾く
		for e := 0; e < len(stock); e++ {
			if url_string == stock[e] {
				continue L
			}
		}

		// キレイ会議外のURLなら弾く
		if !strings.Contains(url_string, base) {
			continue L
		}

		stock = append(stock, url_string)
		results := makeUrl(url_string)
		sourceUrl = append(sourceUrl, results...)
	}

	return sourceUrl
}

// NotFoundなら格納
func checkNotFound(doc *goquery.Document) {
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		img, _ := s.Attr("src")

		if img == notfoundImgPath {
			notfounds = append(notfounds, img)
		}
	})
}

func writeFile(fileName string, stringArray []string) {
	var writer *bufio.Writer

	fmt.Println(fileName)
	fmt.Println(stringArray)
	file, _ := os.OpenFile(fileName, os.O_WRONLY, 0644)
	writer = bufio.NewWriter(file)

	for idx := range stringArray {
		fmt.Println(stringArray[idx])
		writer.Write([]byte(stringArray[idx] + "\n"))
	}

	writer.Flush()
}
