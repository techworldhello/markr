package aggregate

import (
	"encoding/json"
	"github.com/montanaflynn/stats"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/data"
)

func CalculateAverage(scores []float64) string {
	var (
		mean, _ = stats.Mean(scores)
		stdDev, _ = stats.StandardDeviation(scores)
		min, _ = stats.Min(scores)
		max, _ = stats.Max(scores)
		p25, _ = stats.PercentileNearestRank(scores, 25)
		p50, _ = stats.PercentileNearestRank(scores, 50)
		p75, _ = stats.PercentileNearestRank(scores, 75)
	)

	var a data.Aggregate
	a.Mean = getPercentage(mean)
	a.Stddev = getPercentage(stdDev)
	a.Min = getPercentage(min)
	a.Max = getPercentage(max)
	a.P25 = getPercentage(p25)
	a.P50 = getPercentage(p50)
	a.P75 = getPercentage(p75)
	a.Count = len(scores)

	aggregateBytes, err := json.Marshal(&a)
	if err != nil {
		log.Errorf("error marshalling aggregate struct to json: %v", err)
	}
	return string(aggregateBytes)
}

func getPercentage(stat float64) float64 {
	return stat / 20 * 100
}
