package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewVehicleID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
