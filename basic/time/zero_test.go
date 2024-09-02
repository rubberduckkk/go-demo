package time

import (
	"testing"
	"time"
)

func TestZeroValue(t *testing.T) {
	var tt time.Time
	t.Logf("time is: %s\n", tt)
	t.Logf("since time is: %s\n", time.Since(tt))
	t.Logf("is zero: %v\n", tt.IsZero())
}
