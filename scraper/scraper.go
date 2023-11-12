package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Manufacturer struct {
	Href string
	Name string
}

const scrapingUrl = "https://infinitediscs.com"

func main() {
	connectionString := getDbConnection()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	fmt.Print("Starting Scraper", startTime.Local().Format("utc"))

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0"),
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(),
	)
	// Find and visit all links
	// c.OnHTML("div.hero-navs ul.d-flex li.dropper:first-child", func(e *colly.HTMLElement) {
	// 	e.ForEach("li a", func(_ int, a *colly.HTMLElement) {
	// 		// Extract the href attribute and append it to the list
	// 		href := a.Attr("href")
	// 		if strings.Contains(href, "brand") {
	// 			repo.SaveOrUpdateManufacturer(db, a.Text, href)
	// 		} else {
	// 			log.Println("Incorrect href found %s", href)
	// 		}
	// 	})
	// })
	// scrapeErr := c.Visit(scrapingUrl)
	// if scrapeErr != nil {
	// 	log.Fatal(err)
	// }

	processNestedValues(c, db, "/brand/abc")
	endTime := time.Since(startTime)
	fmt.Printf("Scraper took %s to execute\n", endTime)
}

func processNestedValues(c *colly.Collector, db *sql.DB, href string) {
	// Construct  full URL based on the base URL and the collected href
	fullURL := scrapingUrl + href
	fmt.Println("The url", fullURL)
	// Set up a callback for nested values
	c.OnHTML("div#dv_mdl_dd", func(e *colly.HTMLElement) {
		// Extract and process nested values
		e.ForEach("div.pod h4 a", func(_ int, a *colly.HTMLElement) {
			fmt.Println(a.Text)
		})
		nestedValue := e.Text
		fmt.Printf("Nested Value for %s: %s\n", href, nestedValue)

		// You can insert or update the nested values in the database here
		// ...

		// Continue navigating if there are more nested values or details to scrape
	})

	// Set up error handling for the nested collector
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Nested request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping the nested URL
	err := c.Visit(fullURL)
	c.Wait()
	if err != nil {
		log.Printf("Error scraping nested values for %s: %v", href, err)
	}
}

func getDbConnection() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

}
