package shared

var residents = resident.AllResidents()

func GetResidents() []resident.Resident {
	return residents
}

func GetResidentById(id string) (result resident.Resident) {
	//return residents[]
	for _, resident := range residents {
		if resident.Id == id {
			result = resident
			break
		}
	}

	return result
}
