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
	maxSpeedC   string
	maxSpeedKms string
	class       string
	size        string
	purchasableItem
}

type shield struct {
	class string
	size  string
	purchasableItem
	hp           string
	rechargeRate string
}

type cooler struct {
	class string
	size  string
	purchasableItem
	coolingCapacity string
}
