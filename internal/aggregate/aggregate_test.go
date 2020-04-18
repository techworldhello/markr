package aggregate

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	expectations := []struct {
		name  string
		marks []db.DBMarksRecord
		resp  data.Aggregate
	}{
		{
			name: "verify_library_works",
			marks: []db.DBMarksRecord{
				{"1234", 20, 20},
			},
			resp: data.Aggregate{Mean: 100, Stddev: 0, Min: 100, Max: 100, P25: 100, P50: 100, P75: 100, Count: 1},
		},
		{
			name: "dupe_record_uncounted",
			marks: []db.DBMarksRecord{
				{"1234", 19, 17},
				{"1234", 20, 15},
				{"0000", 18, 19},
				{"1111", 20, 16},
			},
			resp: data.Aggregate{Mean: 86.66666666666666, Stddev: 6.236095644623235, Min: 80, Max: 95, P25: 80, P50: 85, P75: 95, Count: 3},
		},
		{
			name: "high_available_mark_recorded",
			marks: []db.DBMarksRecord{
				{"1234", 19, 17},
				{"0000", 18, 19},
				{"1111", 20, 16},
			},
			resp: data.Aggregate{Mean: 86.66666666666666, Stddev: 6.236095644623235, Min: 80, Max: 95, P25: 80, P50: 85, P75: 95, Count: 3},
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			result := CalculateAverage(expect.marks)

			assert.Equal(t, expect.resp, result)
		})
	}
}
