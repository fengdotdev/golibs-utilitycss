package helpers

import (
	"github.com/fengdotdev/golibs-funcs/v0/datum"
	"github.com/fengdotdev/golibs-funcs/v0/ident"
)

// GenerateId generates a unique identifier.
// uses UUID4 to generate a random UUID and returns it as a string.
func GenerateId() string {
	return ident.RamdomUUID()
}

func GenerateHash(s string) string {
	return datum.GetSHA256(s)
}

func DeterministicId(prefix string, id string) string {
	return ident.DeterministicUUID(prefix, id)
}
