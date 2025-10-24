package tesla

import (
	"testing"
)

func TestConcurrentPrint(t *testing.T) {
	start, end := 1, 100

	concurrentPrint2(start, end, 2)
}
