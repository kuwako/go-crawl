package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

type Ext struct {
	*gocrawl.DefaultExtender
}

var rxOk = regexp.MustCompile(`https://kirei-kaigi\.co\.jp(/.*)`)

// urlに訪れた時の処理
func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// URL集める
	fmt.Println(doc.title)
	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

func main() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	// クローリングのオプション設定
	opts := gocrawl.NewOptions(ext)
	opts.MaxVisits = 10
	opts.LogFlags = gocrawl.LogError
	opts.CrawlDelay = 5 * time.Second

	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("https://kirei-kaigi.jp")
}
