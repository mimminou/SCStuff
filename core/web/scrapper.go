package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/mimminou/SCItems/core/csv"
	"net/http"
	"os"
)

/* USAGE ::::::
1 - go to erkul website
2 - open the hamburger menu on the left click on one of the following items : (Ships and Vehicles, Weapons, Shields, Coolers, Quantum Drives)
3 - right click the page and save the whole HTML in the same location as the scrapper
4 - repeat for all other submenus
*/

var baseURL string = "https://www.erkul.games/live"
var ships []ship
var shipGuns []shipGun
var qds []qd
var shields []shield
var powerplants []powerplant
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
	t = &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c = &http.Client{Transport: t}

}

func getUniversalInformation(selection *goquery.Selection) purchasableItem {
	item := purchasableItem{}
	//get names
	name := selection.Find(".cdk-column-name")
	item.name = name.Text()

	//get manufacturers
	manufacturer := selection.Find(".cdk-column-manufacturer")
	item.manufacturer = manufacturer.Text()

	//get prices
	basePriceDiv := selection.Find(".cdk-column-basePrice")
	priceDiv := basePriceDiv.Find(".text-primary")
	if basePriceDiv.Length() > 0 {
		item.basePrice = priceDiv.Text()

	} else {
		item.basePrice = ""
	}
	return item
}

func GetShips() {
	file := "file://" + cwd + "/resources/ships/ships.html"

	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := ship{}

			//universalInformation is name, manufacturer and base price (applies to all purchasable items)
			currentItem.purchasableItem = getUniversalInformation(s)

			//get roles
			role := s.Find(".cdk-column-role")
			currentItem.role = role.Text()

			//get careers
			career := s.Find(".cdk-column-career")
			currentItem.career = career.Text()

			//get crewSizes
			crewSize := s.Find(".cdk-column-crewSize")
			currentItem.crewSize = crewSize.Text()

			//get cargoGrids
			cargoGrid := s.Find(".cdk-column-cargo")
			currentItem.cargoGrid = cargoGrid.Text()

			//get qtTanks
			qtTank := s.Find(".cdk-column-qtFuelCapacity")
			currentItem.qtFuel = qtTank.Text()

			ships = append(ships, currentItem)
		})

		data := [][]string{}
		for _, s := range ships {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.career + " | " + s.role + " | " + s.crewSize + " | " + s.cargoGrid + " | " + s.qtFuel + " | " + s.basePrice)
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.role, s.career, s.crewSize, s.cargoGrid, s.qtFuel})
		}
		csv.WriteCSV("ships", []string{"name", "manufacturer", "basePrice", "role", "career", "crewSize", "cargoGrid", "qtFuel"}, data)
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}
}

func GetShipGuns() {
	file := "file://" + cwd + "/resources/weapons/weapons.html"
	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := shipGun{}

			currentItem.purchasableItem = getUniversalInformation(s)

			//get type
			gunType := s.Find(".cdk-column-type")
			currentItem.gunType = gunType.Text()

			//get size
			size := s.Find(".cdk-column-size")
			currentItem.size = size.Text()

			//get dmg
			dmg := s.Find(".cdk-column-dps")
			currentItem.dmg = dmg.Text()

			//get alphaDmg
			alpha := s.Find(".cdk-column-alpha")
			currentItem.alpha = alpha.Text()

			//get fireRate
			fireRate := s.Find(".cdk-column-firerate")
			currentItem.fireRate = fireRate.Text()

			//get fireRange
			fireRange := s.Find(".cdk-column-range")
			currentItem.fireRange = fireRange.Text()

			//get projectileSpeed
			projSpeed := s.Find(".cdk-column-speed")
			currentItem.projectileSpeed = projSpeed.Text()

			shipGuns = append(shipGuns, currentItem)
		})

		data := [][]string{}
		for _, s := range shipGuns {
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.size, s.gunType, s.dmg, s.alpha, s.fireRate, s.fireRange, s.projectileSpeed})
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.gunType + " | " + s.size + " | " + s.basePrice)
		}
		csvErr := csv.WriteCSV("shipGuns", []string{"name", "manufacturer", "basePrice", "size", "gunType", "dmg", "alpha", "fireRate", "fireRange", "projectileSpeed"}, data)
		if csvErr != nil {
			fmt.Println(csvErr)
		}

	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}
}

func GetQDs() {
	file := "file://" + cwd + "/resources/qds/qds.html"
	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := qd{}

			currentItem.purchasableItem = getUniversalInformation(s)

			//get class
			class := s.Find(".cdk-column-class")
			currentItem.class = class.Text()

			//get size
			size := s.Find(".cdk-column-size")
			currentItem.size = size.Text()

			//get grade
			grade := s.Find(".cdk-column-grade")
			currentItem.grade = grade.Text()

			//get maxSpeed
			maxSpeed := s.Find(".cdk-column-driveSpeed")
			currentItem.maxSpeedKms = maxSpeed.Text()

			//get spoolUpTime
			spoolTime := s.Find(".cdk-column-spoolUpTime")
			currentItem.spoolUpTime = spoolTime.Text()

			qds = append(qds, currentItem)
		})

		data := [][]string{}
		for _, s := range qds {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.grade + " | " + s.size + " | " + s.basePrice)
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.size, s.class, s.grade, s.maxSpeedKms, s.spoolUpTime})
		}
		csvErr := csv.WriteCSV("qds", []string{"name", "manufacturer", "basePrice", "size", "class", "grade", "maxSpeedKms", "spoolUpTime"}, data)
		if csvErr != nil {
			fmt.Println(csvErr)
		}
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}
}

func GetShields() {
	file := "file://" + cwd + "/resources/shields/shields.html"
	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := shield{}

			currentItem.purchasableItem = getUniversalInformation(s)

			//get class
			class := s.Find(".cdk-column-class")
			currentItem.class = class.Text()

			//get size
			size := s.Find(".cdk-column-size")
			currentItem.size = size.Text()

			//get grade
			grade := s.Find(".cdk-column-grade")
			currentItem.grade = grade.Text()

			//get regenRate
			regenRate := s.Find(".cdk-column-maxShieldRegen")
			currentItem.regenRate = regenRate.Text()

			//get hp
			hp := s.Find(".cdk-column-maxShieldHealth")
			currentItem.hp = hp.Text()

			shields = append(shields, currentItem)
		})

		data := [][]string{}
		for _, s := range shields {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.grade + " | " + s.size + " | " + s.basePrice)
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.size, s.class, s.grade, s.hp, s.regenRate})
		}
		csvErr := csv.WriteCSV("shields", []string{"name", "manufacturer", "basePrice", "size", "class", "grade", "hp", "regenRate"}, data)
		if csvErr != nil {
			fmt.Println(csvErr)
		}
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}

}

func GetPowerPlants() {
	file := "file://" + cwd + "/resources/powerplants/powerplants.html"
	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := powerplant{}

			currentItem.purchasableItem = getUniversalInformation(s)

			//get class
			class := s.Find(".cdk-column-class")
			currentItem.class = class.Text()

			//get size
			size := s.Find(".cdk-column-size")
			currentItem.size = size.Text()

			//get grade
			grade := s.Find(".cdk-column-grade")
			currentItem.grade = grade.Text()

			//get power
			power := s.Find(".cdk-column-powerDraw")
			currentItem.drawRequestTime = power.Text()

			//get drawRequestTime
			drawRequestTime := s.Find(".cdk-column-timeToReachDrawRequest")
			currentItem.drawRequestTime = drawRequestTime.Text()

			powerplants = append(powerplants, currentItem)
		})

		data := [][]string{}
		for _, s := range powerplants {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.grade + " | " + s.size + " | " + s.basePrice)
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.size, s.class, s.grade, s.power, s.drawRequestTime})
		}
		csvErr := csv.WriteCSV("powerplants", []string{"name", "manufacturer", "basePrice", "size", "class", "grade", "power", "drawRequestTime"}, data)
		if csvErr != nil {
			fmt.Println(csvErr)
		}
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}
}

func GetCoolers() {
	file := "file://" + cwd + "/resources/coolers/coolers.html"
	collector := colly.NewCollector()
	collector.OnHTML("tbody", func(h *colly.HTMLElement) {
		rows := h.DOM.Find("tr")
		rows.Each(func(i int, s *goquery.Selection) {
			currentItem := cooler{}

			currentItem.purchasableItem = getUniversalInformation(s)

			//get class
			class := s.Find(".cdk-column-class")
			currentItem.class = class.Text()

			//get size
			size := s.Find(".cdk-column-size")
			currentItem.size = size.Text()

			//get grade
			grade := s.Find(".cdk-column-grade")
			currentItem.grade = grade.Text()

			//get cooling rate
			coolingRate := s.Find(".cdk-column-coolingRate")
			currentItem.coolingRate = coolingRate.Text()

			coolers = append(coolers, currentItem)
		})

		data := [][]string{}
		for _, s := range coolers {
			fmt.Println(s.name + " | " + s.manufacturer + " | " + s.grade + " | " + s.size + " | " + s.basePrice)
			data = append(data, []string{s.name, s.manufacturer, s.basePrice, s.size, s.class, s.grade, s.coolingRate})
		}
		csvErr := csv.WriteCSV("coolers", []string{"name", "manufacturer", "basePrice", "size", "class", "grade", "coolingRate"}, data)
		if csvErr != nil {
			fmt.Println(csvErr)
		}
	})

	//get visit the URL
	collector.WithTransport(t)
	visitErr := collector.Visit(file)
	if visitErr != nil {
		fmt.Println(visitErr)
	}
}
