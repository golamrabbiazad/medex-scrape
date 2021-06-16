package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// DrugsList Types specified
// type DrugsList struct {
// 	Title string
// 	Power string
// 	URL   string
// 	Price float64
// }

func main() {
	// drug := []DrugsList{}
	fName := "drug_list.csv"

	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Title", "Power", "Price", "URL"})

	c := colly.NewCollector(colly.AllowedDomains("medex.com.bd"))

	c.OnHTML("div.col-xs-12.data-row-top", func(e *colly.HTMLElement) {
		getTitle := e.Text + "\n"
		title := strings.TrimSpace(getTitle)

		writer.Write([]string{title})
	})

	c.OnHTML("span.grey-ligten", func(h *colly.HTMLElement) {
		quantity := strings.TrimSpace(h.Text)
		writer.Write([]string{quantity})
	})

	c.OnHTML("a.hoverable-block", func(l *colly.HTMLElement) {
		links := l.Attr("href")

		// temp := []DrugsList{}
		// temp.URL = links
		// drug = append(drug, temp)

		c.OnHTML("span.pack-size-info", func(p *colly.HTMLElement) {
			str := p.Text + "\n"
			priceStr := strings.TrimSpace(str)

			prev := strings.Trim(priceStr, "(")
			aft := strings.Trim(prev, ")")

			modStr := strings.Split(aft, " ")
			price := modStr[len(modStr)-1]

			writer.Write([]string{price})

			// if val, err := strconv.ParseFloat(price, 64); err == nil {
			// 	fmt.Println(val)
			// }
		})

		c.Visit(links)
	})

	var pages int
	var companyURL string

	fmt.Printf("How many page do you want to scrape?")
	fmt.Scanln(&pages)

	companyURL = "https://medex.com.bd/companies/73/square-pharmaceuticals-ltd"

	visitURL(1, pages, companyURL, c)

	log.Printf("Scraping finished, check file %q for results\n", fName)

	log.Println(c)

	// fmt.Println(drug)
}

func visitURL(from, to int, companyURL string, c *colly.Collector) {
	for i := from; i <= to; i++ {
		c.Visit(companyURL + "?page=" + strconv.Itoa(i))
	}
}
