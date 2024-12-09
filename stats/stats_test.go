package stats

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestVariance(t *testing.T) {
	data := stats.Float64Data{}
	data = append(data, 2003, 1500)
	variance, _ := stats.Variance(data)
	t.Logf("variance: %f", variance)
}
