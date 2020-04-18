package aggregate

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	marks := []db.DBMarksRecord{
		{"1234", 19, 17},
		{"1234", 20, 15},
		{"0000", 18, 19},
		{"1111", 20, 16},
	}

	result := CalculateAverage(marks)

	assert.Equal(t, data.Aggregate{Mean:86.66666666666666, Stddev:6.236095644623235, Min:80, Max:95, P25:80, P50:85, P75:95, Count:3}, result)
}
