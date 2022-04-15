package resident

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
