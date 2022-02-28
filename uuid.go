package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func UUID() UUIDFunctions {
	return UUIDFunctions{}
}

type UUIDFunctions struct {
}

// GetUUID definition
func (u UUIDFunctions) Get() string {
	return fmt.Sprintf("%v", uuid.New())
}

// GetShortUUID definition
func (u UUIDFunctions) GetShort() string {
	return strings.Split(u.Get(), "-")[0]
}

// IsValidUUID definition
func (u UUIDFunctions) IsValid(uu string) bool {
	_, err := uuid.Parse(uu)
	return err == nil
}
