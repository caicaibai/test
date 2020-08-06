package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	fmt.Println("hello,world")

	c := colly.NewCollector()

	c.AllowedDomains = []string{"www.youzy.cn"}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		// Print link

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found on page

		// Only those links are visited which are in AllowedDomains

		c.Visit(e.Request.AbsoluteURL(link))

	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL.String())

	})

	// Start scraping on https://hackerspaces.org

	c.Visit("https://www.youzy.cn/tzy/search/majors/homepage")

}
