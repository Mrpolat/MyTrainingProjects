package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Asset struct {
	OrderNumberForMatching	string `json:"OrderNumberForMatching"`
	Description				string `json:"Description"`
	Family                  string `json:"Family"`
	AssetType               string `json:"AssetType"`
	Racks                   string `json:"Racks"`
	NumberOfSlots           string `json:"NumberOfSlots"`
	DiscreteIO              string `json:"DiscreteIO "`
	AnalogueIO              string `json:"AnalogueIO"`
	DiscreteInputNumber		string `json:"DiscreteInputNumber"`
	DiscreteOutputNumber	string `json:"DiscreteOutputNumber"`
	DiscreteIONumber		string `json:"DiscreteIONumber"`
	Interface      		    string `json:"Interface"`
	Protocols			    string `json:"Protocols"`
	WebService				string `json:"WebService"`
	Firmware				string `json:"Firmware"`
	ImageUrl				string `json:"ImageUrl"`	
	Monitoring				string `json:"Monitoring"`
	Name					string `json:"Name"`
	EngineeringWith			string `json:"EngineeringWith"`
	WhereToGet				string `json:"WhereToGet"`

}	
func main() {
	fmt.Println("--->SCHNEİDER Scraper Initialized")

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	asset := Asset{}
	assetList := []Asset{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(`#start > div.pdp-product-info > div.pdp-product-wrapper--desktop > div.pdp-product-info__id`, func(e *colly.HTMLElement) {
		asset.OrderNumberForMatching = e.ChildText("span")
		
	})
	c.OnHTML(`#start > div.pdp-product-info`, func(e *colly.HTMLElement) {
		asset.Description = e.ChildText("h2")
	})
	c.OnHTML(`#start > div:nth-child(1) > div.product-main-info__main-product-wrapper > div`, func(e *colly.HTMLElement) {
		asset.ImageUrl = e.ChildAttr("img","src")
	})
    //   tablonun html ini aldı Selector ile 
	c.OnHTML(`#characteristics > div`, func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			switch el.ChildText("tr > th") { // sol sütünü kontrol etti 
			case "range of product": // sol sutunda for use with yazıyor ise karşısındaki değeri asset.family e yazdırdı
				asset.Family  				     = el.ChildText("tr > td")
			case "product or component type":
				asset.AssetType		 			 = el.ChildText("tr > td")
			case "number of racks":
				asset.Racks	 		 			 = el.ChildText("tr > td")
			case "number of slots":
				asset.NumberOfSlots	 			 = el.ChildText("tr > td")
			case "discrete I/O processor capacity":
				asset.DiscreteIO  	 			 = el.ChildText("tr > td")
			case "analogue I/O processor capacity":
				asset.AnalogueIO 				 = el.ChildText("tr > td")
			case "Local I/O processor capacity (discrete)":
				asset.DiscreteIO  			     = el.ChildText("tr > td")
			case "Local I/O processor capacity (analog)":
				asset.AnalogueIO 			     = el.ChildText("tr > td")
			case "discrete input number":
				asset.DiscreteInputNumber   	 = el.ChildText("tr > td")
			case "discrete output number":
				asset.DiscreteOutputNumber   	 = el.ChildText("tr > td")
			case "discrete I/O number":
				asset.DiscreteIONumber   		 = el.ChildText("tr > td")
			case "integrated connection type":
				asset.Interface     			 = el.ChildText("tr > td")
			case "communication port protocol":
				asset.Protocols	  	 			 = el.ChildText("tr > td")
			case "web services":
				asset.WebService   				 = el.ChildText("tr > td")
			case "Firmware":
				asset.Firmware   				 = el.ChildText("tr > td")
			// case "ImageUrl":
			// 	asset.ImageUrl  				 = el.ChildText("tr > td")
			case "monitoring":
				asset.Monitoring   				 = el.ChildText("tr > td")
			case "Name":
				asset.Name   				     = el.ChildText("tr > td")
			case "EngineeringWith":
				asset.EngineeringWith   	     = el.ChildText("tr > td")
			case "WhereToGet":
				asset.WhereToGet   				 = el.ChildText("tr > td")
			
			}
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	file, _ := os.Open("url.teikste")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		c.Visit(fileScanner.Text())

		assetList = append(assetList, asset)
		asset = Asset{}
	}

	fmt.Println(assetList)

	output, _ := json.MarshalIndent(assetList, "", " ")
	_ = ioutil.WriteFile("Schneider.json", output, 0644)
}
