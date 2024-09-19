package ctrl

import (
	"log"
	"testing"
)

func switchStatements(val int) {
	switch val {
	case 0:
		log.Printf("val is 0\n")
		return
	case 1:
		log.Printf("val is 1\n")
		return
	}
	log.Printf("val is unknown value %d\n", val)
}

func TestSwitch(t *testing.T) {
	switchStatements(0)
	switchStatements(7)
}
