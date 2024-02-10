package main

import (
	web "github.com/mimminou/SCItems/core/web"
)

func main() {
	web.GetShips()
	web.GetShipGuns()
	web.GetShields()
	web.GetQDs()
	web.GetPowerPlants()
	web.GetCoolers()
}
