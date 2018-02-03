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

var rxOk = regexp.MustCompile(`http://amazon\.co\.jp\/(review\/top-reviewers.*page.*|gp\/pdp.*pic).*`)
var rxTopReviewer = regexp.MustCompile(`http://amazon\.co\.jp\/gp\/pdp.*pic.*`)

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	if rxTopReviewer.MatchString(ctx.NormalizedURL().String()) {
		isSucker, numSucks := isSucker(doc)
		if isSucker {
			fmt.Printf("%s,%d\n", ctx.URL(), numSucks)
		}
	}
	return nil, true
}

func isSucker(doc *goquery.Document) (bool, int) {
	num := len(doc.Find(".a-star-medium-1").Nodes)
	return num > 4, num
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

func sampleMain2() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 5 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.MaxVisits = 20000

	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("http://www.amazon.co.jp/review/top-reviewers/ref=cm_cr_tr_link_2?ie=UTF8&page=1")
}
