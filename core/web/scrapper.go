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

	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currShip := ship{}

			//get names
			name := s.Find(".cdk-column-name")
			currShip.name = name.Text()

			//get manufacturers
			manufacturer := s.Find(".cdk-column-manufacturer")
			currShip.manufacturer = manufacturer.Text()

			//get prices
			basePriceDiv := s.Find(".cdk-column-basePrice")
			priceDiv := basePriceDiv.Find(".text-primary")
			if basePriceDiv.Length() > 0 {
				currShip.basePrice = priceDiv.Text()

			} else {
				currShip.basePrice = ""
			}

			//get roles
			role := s.Find(".cdk-column-role")
			currShip.role = role.Text()

			//get careers
			career := s.Find(".cdk-column-career")
			currShip.career = career.Text()

			//get crewSizes
			crewSize := s.Find(".cdk-column-crewSize")
			currShip.crewSize = crewSize.Text()

			//get cargoGrids
			cargoGrid := s.Find(".cdk-column-cargo")
			currShip.cargoGrid = cargoGrid.Text()

			//get qtTanks
			qtTank := s.Find(".cdk-column-qtFuelCapacity")
			currShip.qtFuel = qtTank.Text()

			ships = append(ships, currShip)
		})

		//TODO : Need to write the ships struct to CSV instead of printing
		for _, s := range ships {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.career + " | " + s.role + " | " + s.crewSize + " | " + s.cargoGrid + " | " + s.qtFuel + " | " + s.basePrice)
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
