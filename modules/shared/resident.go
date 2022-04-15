package shared

import (
	"github.com/boothr/godev/modules/domain"
)

var residents = domain.AllResidents()

func GetResidents() *[]domain.Resident {
	return residents
}

func GetResidentById(id string) (result domain.Resident) {
	//return residents[]
	for _, resident := range *residents {
		if resident.Id == id {
			result = resident
			break
		}
	}

	return result
}
