package wepie2

import "testing"

func TestSchedule(t *testing.T) {
	type testData struct {
		name     string
		tasks    []string
		cooldown int
		want     int
	}

	data := []testData{
		{
			name:     "case demo",
			tasks:    []string{"A", "A", "A", "B", "B", "B"},
			cooldown: 2,
			want:     8,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			get := schedule(d.tasks, d.cooldown)
			t.Logf("case: %v, get=%v, want=%v\n", d.name, get, d.want)
			if get != d.want {
				t.Fail()
			}
		})
	}
}
