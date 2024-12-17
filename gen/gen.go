package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

// generateRandomChinese generates a random Chinese character.
func generateRandomChinese() rune {
	// Unicode range for common Chinese characters (CJK Unified Ideographs)
	// 0x4E00 to 0x9FFF
	return rune(0x4E00 + rand.Intn(0x9FFF-0x4E00+1))
}

// generateRandomEnglish generates a random lowercase English letter.
func generateRandomEnglish() rune {
	return rune('a' + rand.Intn(26))
}

// generateNickname generates a nickname starting with "ai-" and total length between 5 and 8 characters.
func generateNickname() string {
	// Random length between 5 and 8
	length := rand.Intn(4) + 5 // (5, 6, 7, 8)

	// Build the nickname
	nickname := "ai-"
	for utf8.RuneCountInString(nickname) < length {
		if rand.Intn(2) == 0 {
			// Add a random Chinese character
			nickname += string(generateRandomChinese())
		} else {
			// Add a random English letter
			nickname += string(generateRandomEnglish())
		}
	}

	return nickname
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create("test_nicknames.csv")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Generate and print 100 random nicknames
	for i := 0; i < 100; i++ {
		_ = csvWriter.Write([]string{generateNickname()})
	}
}
