package csv

import (
	"encoding/csv"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestCSV(t *testing.T) {
	f, err := os.Open("./nickname.csv")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	// csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("data len: %d", len(data))
	//
	// for i, row := range data {
	// 	for j, field := range row {
	// 		field = fixQuotes(field)
	// 		data[i][j] = field
	// 	}
	// }
	//
	// err = writeCSV("fix_nickname.csv", data)
	// t.Logf("write csv error: %v", err)
}

func TestRemoveEmptyLines(t *testing.T) {
	f, err := os.Open("./fix_nickname.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	newData := make([][]string, 0, len(data))
	for _, row := range data {
		if isEmptyRow(row) {
			continue
		}
		newData = append(newData, row)
	}

	err = writeCSV("./fix_empty_nickname.csv", newData)
	t.Logf("write csv len: %d, err: %v", len(newData), err)
}

// fixQuotes ensures double quotes inside quoted fields are properly escaped
func fixQuotes(field string) string {
	// Remove leading/trailing spaces
	field = strings.TrimSpace(field)

	// If field contains a bare ", replace it with escaped ""
	if strings.Count(field, "\"")%2 != 0 {
		// Remove or escape bare quotes using regex
		re := regexp.MustCompile(`([^,])"([^,])`)
		field = re.ReplaceAllString(field, `$1""$2`)
	}
	return field
}

// writeCSV writes fixed records to a new CSV file
func writeCSV(filename string, records [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)
}

// isEmptyRow checks if a row is empty (all fields are blank)
func isEmptyRow(row []string) bool {
	for _, field := range row {
		if strings.TrimSpace(field) != "" {
			return false // Not empty
		}
	}
	return true // Empty row
}
