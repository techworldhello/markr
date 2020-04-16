package aggregate

import (
	"github.com/montanaflynn/stats"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
)

func CalculateAverage(records []db.DBMarksRecord) data.Aggregate {
	var (
		obtained []float64
		available int
	)

	duplicates := map[int]bool{}

	for _, record := range records {
		if record.Available > available {
			available = record.Available
		}
		if !duplicates[record.StudentId] {
			duplicates[record.StudentId] = true
			obtained = append(obtained, float64(record.Obtained))
		}
	}

	var (
		mean, _ = stats.Mean(obtained)
		stdDev, _ = stats.StandardDeviation(obtained)
		min, _ = stats.Min(obtained)
		max, _ = stats.Max(obtained)
		p25, _ = stats.PercentileNearestRank(obtained, 25)
		p50, _ = stats.PercentileNearestRank(obtained, 50)
		p75, _ = stats.PercentileNearestRank(obtained, 75)
	)

	return data.Aggregate{
		Mean: getPercentage(mean, available),
		Stddev: getPercentage(stdDev, available),
		Min: getPercentage(min, available),
		Max: getPercentage(max, available),
		P25: getPercentage(p25, available),
		P50: getPercentage(p50, available),
		P75: getPercentage(p75, available),
		Count: len(obtained),
	}
}

func getPercentage(stat float64, total int) float64 {
	return stat / float64(total) * 100
}
