package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02-15"))
}
