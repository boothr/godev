package domain

type Resident struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	ResidentName string `json:"resident-name"`
}

var Residents = []Resident{
	{Id: "1", Address: "111 Fake St", ResidentName: "Test McTesterson"},
	{Id: "2", Address: "123 Fake St", ResidentName: "Tess Masters"},
	{Id: "3", Address: "654 Fake St", ResidentName: "Chesterton Roads"},
}

func AllResidents() *[]Resident {
	return &Residents
}
