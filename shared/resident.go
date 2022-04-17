package shared

import (
	"fmt"
	"strconv"

	"github.com/boothr/godev/residents/domain"
)

var residents = domain.AllResidents()
var nextId = "4"

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

func AddResident(r *domain.Resident) bool {
	// Add the new album to the slice.
	r.Id = nextId

	n, err := strconv.ParseInt(nextId, 0, 16)
	if err != nil {
		return false
	}

	nextId = fmt.Sprint(n + 1)
	*residents = append(*residents, *r)
	return true
}

func GetNewResident() *domain.Resident {
	var newResident domain.Resident
	return &newResident
}
