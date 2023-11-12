package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Manufacturer struct {
	Href string
	Name string
}

func main() {
	c := colly.NewCollector()
	startTime := time.Now()
	fmt.Print("Starting Scraper", startTime.Local().Format("utc"))
	// Find and visit all links
	var manufacturerList []Manufacturer
	const scrapingUrl = "https://infinitediscs.com"
	c.OnHTML("div.hero-navs ul.d-flex li.dropper:first-child", func(e *colly.HTMLElement) {
		e.ForEach("li a", func(_ int, a *colly.HTMLElement) {
			// Extract the href attribute and append it to the list
			href := a.Attr("href")
			if strings.Contains(href, "brand") {
				manufacturer := Manufacturer{Name: a.Text, Href: href}
				manufacturerList = append(manufacturerList, manufacturer)
			} else {
				log.Println("Incorrect href found %s", href)
			}
		})
	})
	fmt.Println("starting loop")

	err := c.Visit(scrapingUrl)
	if err != nil {
		log.Fatal(err)
	}
	for _, href := range manufacturerList {
		fmt.Println("Href:", href.Name, href.Href)
	}
	endTime := time.Now().Sub(startTime)
	fmt.Printf("Scraper took %s to execute\n", endTime)
}

// saveOrUpdateData inserts or updates data in the database
func saveOrUpdateData(db *sql.DB, name, href string) error {
	// Use an UPSERT (INSERT ON CONFLICT UPDATE) statement
	query := `
		INSERT INTO your_table_name (name, href)
		VALUES ($1, $2)
		ON CONFLICT (name) DO UPDATE
		SET href = EXCLUDED.href;
	`

	_, err := db.Exec(query, name, href)
	return err
}
