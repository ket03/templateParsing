// you can delete "time" and "user-agent" if site dont blocking your request
package main

import (
	"fmt"
	// "time"

	"github.com/gocolly/colly" // custom library
)

func main() {
	pokemons := map[string]string{}
	URL := "https://scrapeme.live/shop/" // editable
	// UserAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36" 

	c := colly.NewCollector()
	// c.UserAgent = UserAgent

	// c.OnRequest(func(r *colly.Request) { // optional
	// 	time.Sleep(1 * time.Second) 
	// })
	
	c.OnRequest(func(r *colly.Request) { 
		fmt.Println("Visiting: ", r.URL) 
	})

	// c.OnResponse(func(r *colly.Response) {
    //     fmt.Println(r.StatusCode)
    // })
	 
	c.OnError(func(_ *colly.Response, err error) { 
		fmt.Println("Something went wrong: ", err) 
	}) 

	c.OnHTML("li.product", func(e *colly.HTMLElement) { // MAIN FUNCTION
		pokemons[e.ChildText("h2")] = e.ChildText(".price")
		})

	c.Visit(URL)
	 
	// c.OnScraped(func(r *colly.Response) { // optional
	// 	fmt.Println(r.Request.URL, " scraped!") 
	// })

	for key, value := range pokemons {
		fmt.Printf("name: %s\nprice: %s\n\n", key, value)
	}
}