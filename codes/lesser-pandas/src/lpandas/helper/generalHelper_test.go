package helper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "fmt"

	"lpandas/helper"
)

func TestPadString(t *testing.T) {
	str := "ab"

	assert.Equal(t, "ab   ", helper.PadString(str, " ", 5))
}