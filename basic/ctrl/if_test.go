package ctrl

import (
	"log"
	"testing"
)

func checkSomeCondition(val int) bool {
	log.Printf("check condition for %v\n", val)
	return val < 10
}

func TestIf(t *testing.T) {
	oneCondition := checkSomeCondition(300)

	if anotherCond := checkSomeCondition(600); oneCondition && anotherCond {
		t.Logf("in if block")
	}
}
