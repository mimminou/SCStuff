package web

type purchasableItem struct {
	name, manufacturer, basePrice string
}

type ship struct {
	purchasableItem
	role      string
	career    string
	crewSize  string
	cargoGrid string
	qtFuel    string
}

type shipGun struct {
	purchasableItem
	gunType         string
	size            string
	dmg             string
	alpha           string
	fireRate        string
	fireRange       string
	projectileSpeed string
}

type qd struct {
	spoolUpTime string
	maxSpeedKms string
	class       string
	grade       string
	size        string
	purchasableItem
}

type shield struct {
	class string
	grade string
	size  string
	purchasableItem
	hp        string
	regenRate string
}

type powerplant struct {
	class string
	grade string
	size  string
	purchasableItem
	power           string
	drawRequestTime string
}

type cooler struct {
	class string
	size  string
	grade string
	purchasableItem
	coolingRate string
}
