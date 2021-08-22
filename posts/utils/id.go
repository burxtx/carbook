package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewPostID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
