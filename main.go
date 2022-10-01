package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// coll'ys main entity is a Collector object.
	// Collector manages the network communication and responsible
	// for the execution of the attached callbacks while a collector job is running.
	// To work with colly, you have to initialize a Collector:
	collector := colly.NewCollector()
	collector.SetRequestTimeout(120 * time.Second)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	collector.Visit("https://www.linkedin.com/jobs/search/?currentJobId=3227124654&geoId=101282230&keywords=backend%20engineer&location=Germany&refresh=true")
}
