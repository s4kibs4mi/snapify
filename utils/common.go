package utils

import (
	"github.com/satori/go.uuid"
)

func NewUUID() string {
	v := uuid.NewV4()
	return v.String()
}
