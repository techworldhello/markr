package aggregate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	scores := []float64{13, 15, 19, 12}

	result := CalculateAverage(scores)

	assert.Equal(t, `{"mean":73.75,"stddev":13.404756618454511,"min":60,"max":95,"p25":60,"p50":65,"p75":75,"count":4}`, result)
}
