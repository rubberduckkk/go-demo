package time

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Ref: https://eli.thegreenplace.net/2020/unmarshaling-time-values-from-json/

type Config struct {
	Activity ActivityConfig `json:"activity"`
}

type ActivityConfig struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

const rawConfig = `
{
    "activity": {
        "act_id": "2",
        "start_time": "2024-06-03T10:00:00+08:00",
        "end_time": "2024-07-03T23:59:59+08:00",
        "reward_skin_id": 320001,
        "max_ad_count": 5
    }
}
`

func TestUnmarshalTime(t *testing.T) {
	var cfg Config
	err := json.Unmarshal([]byte(rawConfig), &cfg)
	assert.Nil(t, err)
	fmt.Printf("%+v\n", cfg)
}
