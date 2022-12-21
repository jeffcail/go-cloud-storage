package utils

import uuid "github.com/satori/go.uuid"

// GenerateUuid
func GenerateUuid() string {
	return uuid.NewV4().String()
}
