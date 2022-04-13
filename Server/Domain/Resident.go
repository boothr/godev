package Domain

type Resident struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	ResidentName string `json:"resident-name"`
}

var residents = []Resident{
	{Id: "1", Address: "111 Fake St", ResidentName: "Test McTesterson"},
	{Id: "2", Address: "123 Fake St", ResidentName: "Tess Masters"},
	{Id: "3", Address: "654 Fake St", ResidentName: "Chesterton Roads"},
}

func GetResidents() []Resident {
	return residents
}

func GetResidentById(id string) (result Resident) {
	//return residents[]
	for _, resident := range residents {
		if resident.Id == id {
			result = resident
			break
		}
	}

	return result
}
