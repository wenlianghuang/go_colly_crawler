package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	var wg sync.WaitGroup
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML(".product-item", func(e *colly.HTMLElement) {
		title := e.ChildText(".product-item-link")
		fmt.Printf("Title: %s\n", title)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(".item.pages-item-next a.action.next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		//fmt.Println("Next page link found:", nextPage)
		time.Sleep(100 * time.Millisecond)
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Visit(nextPage)
		}()
	})

	startURL := "https://store.acer.com/zh-tw/laptops?p=1"
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Visit(startURL)
	}()

	wg.Wait()
	c.Wait()
}
