package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Drugs Types specified
type Drugs struct {
	Title       string
	Power 		string
	URL         string
	Price		float64
}

func main() {
	stories := []Drugs{}

	// fName := "drug_list.csv"

	// file, err := os.Create(fName); if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s\n", fName, err)
	// 	return
	// }

	// defer file.Close()

	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	c := colly.NewCollector(colly.AllowedDomains("medex.com.bd"))

	c.OnHTML("div.col-xs-12.data-row-top", func(e *colly.HTMLElement) {
		getTitle := e.Text + "\n"
		title := strings.TrimSpace(getTitle)

		fmt.Println(title)
	})

	c.OnHTML("span.grey-ligten", func(h *colly.HTMLElement) {
		quantity := strings.TrimSpace(h.Text)
		fmt.Println(quantity)
	})

	c.OnHTML("a.hoverable-block", func(l *colly.HTMLElement) {
		links := l.Attr("href")

		temp := Drugs{}
		temp.URL = links
		stories = append(stories, temp)

		c.OnHTML("span.pack-size-info", func(p *colly.HTMLElement) {
			str := p.Text + "\n"
			priceStr := strings.TrimSpace(str)

			prev := strings.Trim(priceStr, "(")
			aft := strings.Trim(prev, ")")
		
			modStr := strings.Split(aft, " ")
			price := modStr[len(modStr)-1]
		
			if val, err := strconv.ParseFloat(price, 64); err == nil {
				fmt.Println(val)
			}
		})

		c.Visit(links)
	})

	pages := 28

	for i := 1; i <= pages; i++ {
		c.Visit("https://medex.com.bd/companies/73/square-pharmaceuticals-ltd?page=" + strconv.Itoa((i)))
	}

	fmt.Println(stories)
}
