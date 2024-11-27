package main

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	var wg sync.WaitGroup
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		title := e.ChildText("h3 a")
		price := e.ChildText(".price_color")
		fmt.Printf("Title: %s, Price: %s\n", title, price)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			c.Visit(fmt.Sprintf("http://books.toscrape.com/catalogue/page-%d.html", page))
		}(i)
	}

	wg.Wait()
	c.Wait()
}
