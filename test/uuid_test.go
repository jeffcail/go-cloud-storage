package test

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestGenerateUuid(t *testing.T) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
}
