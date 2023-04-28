package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	pageToVist := "https://datatables.net/examples/styling/bootstrap5.html"
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//if the file exists, write the header

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		writter := csv.NewWriter(file)
		defer writter.Flush()
	}

	writter := csv.NewWriter(file)
	defer writter.Flush()

	c := colly.NewCollector()
	c.OnHTML("example", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writter.Write([]string{
				el.ChildText(fmt.Sprintf("td:nth-child(%d)", e.Index)),
				el.ChildText(fmt.Sprintf("td:nth-child(%d)", 2)),
				el.ChildText(fmt.Sprintf("td:nth-child(%d)", 3)),
				el.ChildText(fmt.Sprintf("td:nth-child(%d)", 4)),
				el.ChildText(fmt.Sprintf("td:nth-child(%d)", 5)),
			})
		})
		fmt.Println("Scraping Done...")
	})
	c.Visit(pageToVist)
}
