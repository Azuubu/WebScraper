package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Item struct {
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	Price string `json:"price"`
	Change string `json:"change"`
	MarketCap string `json:"marketcap"`
	VolumeInCurrency24h string `json:"volume24h"`
	CirculatingSupply string `json:"circulatinsupply"`
	SymbolImgUrl string `json:"symbolimgurl"`
}

var items []Item

func main() {
	start := time.Now()
	
	c := colly.NewCollector()

	c.OnHTML("#fin-scr-res-table tr.simpTblRow", func(h *colly.HTMLElement) { 

		item := Item {
			Symbol: h.ChildText(`td[aria-label="Symbol"] a[data-test="quoteLink"]`),
			Name: h.ChildText(`td[aria-label="Name"]`),
			Price: h.ChildText(`td[aria-label="Price (Intraday)"] `),
			Change: h.ChildText(`td[aria-label="Change"]`),
			MarketCap: h.ChildText(`td[aria-label="Market Cap"]`),
			VolumeInCurrency24h: h.ChildText(`td[aria-label="Volume in Currency (24Hr)"]`) ,
			CirculatingSupply: h.ChildText(`td[aria-label="Circulating Supply"]`),
			SymbolImgUrl: h.ChildAttr(`td[aria-label="Symbol"] img`, "src"),
		}
		items = append(items, item)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Current URL used:", r.URL.String())
	})
	
	c.Visit("https://finance.yahoo.com/crypto")

	elapsed := time.Since(start)
	fmt.Println("Scraping was finished in:", elapsed)
	
	saveDataJson()
	displayUtilityData(items)
} 


func saveDataJson() {
	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("productData.json", content, 0644)
	
}


func displayUtilityData(items []Item) {
	fmt.Println("Number of items:", len(items))
}




