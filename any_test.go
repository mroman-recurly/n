package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {

	// Test empty queryable
	assert.False(t, S().Any())

	// Test empty collection object
	assert.False(t, Q([]int{}).Any())

	// Test value object
	assert.True(t, Q(1).Any())

	// Test string object
	assert.True(t, Q("2").Any())
}

func TestAnyWhere(t *testing.T) {

}
