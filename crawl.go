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

	return nul, true
}
