package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUserID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
