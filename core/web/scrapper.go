package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"net/http"
	"os"
)

/* USAGE ::::::
1 - go to erkul website
2 - open the hamburger menu on the left click on one of the following items : (Ships and Vehicles, Weapons, Shields, Coolers, Quantum Drives)
3 - right click the page and save the whole HTML in the same location as the scrapper
4 - repeat for all other submenus
*/

var collector *colly.Collector
var baseURL string = "https://www.erkul.games/live"
var ships []ship
var shipGuns []shipGun
var qds []qd
var shields []shield
var coolers []cooler
var t *http.Transport
var c *http.Client
var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		fmt.Println("can't get path to current working dir")
		os.Exit(1)
	}
	collector = colly.NewCollector()
	t = &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c = &http.Client{Transport: t}

	//callbacks
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("requesting")
	})

	collector.OnResponse(func(r *colly.Response) {

		fmt.Println("got response")
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error getting to the website")
		fmt.Println("----------------------------")
		fmt.Println(err)
		fmt.Println("----------------------------")
	})
}

func GetShips() {
	file := "file://" + cwd + "/ships.html"
	var names []string
	var manufacturers []string
	var prices []string
	var roles []string
	var careers []string
	var crewSizes []string
	var cargoGrids []string
	var qtTanks []string

	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		//get names
		rows.Find(".cdk-column-name").Each(func(i int, s *goquery.Selection) {
			names = append(names, s.Text())
		})

		//get manufacturers
		rows.Find(".cdk-column-manufacturer").Each(func(i int, s *goquery.Selection) {
			manufacturers = append(manufacturers, s.Text())
		})

		//get prices
		rows.Find(".cdk-column-basePrice").Each(func(i int, s *goquery.Selection) {
			priceDiv := s.Find(".text-primary")
			if priceDiv.Length() > 0 {
				prices = append(prices, priceDiv.Text())

			} else {
				prices = append(prices, "")
			}
		})

		//get roles
		rows.Find(".cdk-column-role").Each(func(i int, s *goquery.Selection) {
			roles = append(roles, s.Text())
		})

		//get careers
		rows.Find(".cdk-column-career").Each(func(i int, s *goquery.Selection) {
			careers = append(careers, s.Text())
		})

		//get crewSizes
		rows.Find(".cdk-column-crewSize").Each(func(i int, s *goquery.Selection) {
			crewSizes = append(crewSizes, s.Text())
		})

		//get cargoGrids
		rows.Find(".cdk-column-cargo").Each(func(i int, s *goquery.Selection) {
			cargoGrids = append(cargoGrids, s.Text())
		})

		//get qtTanks
		rows.Find(".cdk-column-qtFuelCapacity").Each(func(i int, s *goquery.Selection) {
			qtTanks = append(qtTanks, s.Text())
		})

		for i := range names {
			ships = append(ships, ship{purchasableItem: purchasableItem{name: names[i],
				manufacturer: manufacturers[i],
				basePrice:    prices[i]},
				role:      roles[i],
				career:    careers[i],
				cargoGrid: cargoGrids[i],
				crewSize:  crewSizes[i],
			})
		}

		//write the ships struct to CSV instead of printing
		for _, s := range ships {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.career + " | " + s.role)
		}
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}

}

func GetShipGuns() {
	route := baseURL + "/weapons"
	collector.Visit(route)
}

func QetQDs() {
	route := baseURL + "/quantum-drives"
	collector.Visit(route)
}

func GetShields() {
	route := baseURL + "/shields"
	collector.Visit(route)
}

func GetPowerPlants() {
	route := baseURL + "/power-plants"
	collector.Visit(route)
}

func GetCoolers() {
	route := baseURL + "/coolers"
	collector.Visit(route)
}
