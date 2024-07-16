package csv

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	f, err := os.Open("uids.csv")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("data len: %d", len(data))
}
