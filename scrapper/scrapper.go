package scrapper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Response struct {
	Quote  string   `json:"quote"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func ScrapeTheURL(url string, classes []string) {
	fmt.Printf("URL: %s, Class: %s\n", url, classes)

	// create new colly collector
	newCollector := colly.NewCollector()

	// URL to scrape
	// url := "https://quotes.toscrape.com/page/1/"

	// CSS Selectors
	// "span.text"
	// "small.author"
	// "div.tags"

	//Store classes in map[string]bool
	classesMap := make(map[string]bool)

	for _, class := range classes {
		_, ok := classesMap[class]
		if !ok {
			classesMap[class] = true
		}
	}
	// fmt.Println("Map Data:", classesMap)

	// Set up callbacks to handle scrapping events
	newCollector.OnHTML(".quote", func(h *colly.HTMLElement) {
		res := &Response{}

		// Extract data from HTML elements
		for class := range classesMap {
			if data := h.ChildText(class); data != "" {
				switch class {
				case classes[0]:
					res.Quote = strings.TrimSpace(data)
				case classes[1]:
					res.Author = strings.TrimSpace(data)
				case classes[2]:
					h.ForEach(class, func(_ int, elem *colly.HTMLElement) {
						res.Tags = append(res.Tags, strings.TrimSpace(elem.Text))
					})
				}
			}
		}

		// Print scrapped data
		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n%s\n", res.Quote, res.Author, res.Tags, strings.Repeat("-", 50))
	})

	// Visit the URL and start scrapping
	err := newCollector.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
