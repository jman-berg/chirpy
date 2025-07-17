package auth

import (
	"testing"
)

func TestHashPasword(t *testing.T) {
	_, err := HashPassword("Gurbe")
	if err != nil {
		t.Errorf("Failed to hash password %v", err)
	}
}
