package logpkg

import (
	"encoding/json"
	"fmt"
)

func Print(args ...interface{}) {
	for _, arg := range args {
		raw, _ := json.Marshal(arg)
		fmt.Println(string(raw))
	}
}
