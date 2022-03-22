package validator

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRegPassword(t *testing.T) {
	assert.Equal(t, verifyPassword("1223445"), false)
	assert.Equal(t, verifyPassword("12234a"), true)
}
