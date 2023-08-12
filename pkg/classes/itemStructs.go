package classes

type ArmorAnswer struct {
	ArmorName       string `json:"armorName"`
	ArmorType       string `json:"armorType"`
	ArmorClassCount int    `json:"armorClassCount"`
	Stealth         bool   `json:"armorStealth"`
}

type WeaponAnswer struct {
	WeaponName string `json:"weaponName"`
	WeaponType string `json:"weaponType"`
	Damage     string `json:"damage"`
	Property   string `json:"property"`
	Count      int    `json:"count"`
}

type equipPicks []struct {
	Variants []Variants `json:"variants"`
}
type Variants struct {
	ItemName string `json:"itemName"`
	Type     string `json:"type"`
	Count    int    `json:"count"`
}

type equipBasic []struct {
	ItemName string `json:"itemName"`
	Type     string `json:"type"`
	Count    int    `json:"count"`
}
