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


// GenerateHash generates a SHA256 hash of the input string.
// It uses the GetSHA256 function from the datum package to compute the hash.
// The hash is returned as a string.
func GenerateHash(s string) string {
	return datum.GetSHA256(s)
}


// DeterministicId generates a deterministic identifier based on a prefix and an id.
// It uses the ident.DeterministicUUID function to create a UUID that is consistent for the same prefix and id.
// This is useful for generating stable identifiers that can be reproduced across different runs.
func DeterministicId(prefix string, id string) string {
	return ident.DeterministicUUID(prefix, id)
}
