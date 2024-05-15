package ratelimit

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_fixedWindowRateLimiter_calculateCutoff(t *testing.T) {
	now := time.Date(2024, 5, 15, 14, 59, 35, 58, time.UTC)
	t.Logf("now used for test: %s", now)

	type fields struct {
		windowSize FixedWindowSize
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "case window size minute",
			fields: fields{
				windowSize: WindowSizeMinute,
			},
			args: args{
				now: now,
			},
			want: time.Date(2024, 5, 15, 14, 59, 0, 0, now.Location()),
		},
		{
			name: "case window size hour",
			fields: fields{
				windowSize: WindowSizeHour,
			},
			args: args{
				now: now,
			},
			want: time.Date(2024, 5, 15, 14, 0, 0, 0, now.Location()),
		},
		{
			name: "case window size day",
			fields: fields{
				windowSize: WindowSizeDay,
			},
			args: args{
				now: now,
			},
			want: time.Date(2024, 5, 15, 0, 0, 0, 0, now.Location()),
		},
		{
			name: "case window size month",
			fields: fields{
				windowSize: WindowSizeMonth,
			},
			args: args{
				now: now,
			},
			want: time.Date(2024, 5, 1, 0, 0, 0, 0, now.Location()),
		},
		{
			name: "case window size year",
			fields: fields{
				windowSize: WindowSizeYear,
			},
			args: args{
				now: now,
			},
			want: time.Date(2024, 1, 1, 0, 0, 0, 0, now.Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fixedWindowRateLimiter{
				windowSize: tt.fields.windowSize,
			}
			cutoff := f.calculateCutoff(tt.args.now)
			assert.Equal(t, cutoff, tt.want)
		})
	}
}

func Test_fixedWindowRateLimiter_Allow(t *testing.T) {
	cli := getLocalRedisCli()
	limiter, _ := NewFixedWindowRateLimiter(cli, WindowSizeMinute, 5)

	const collection = "rewards_fixed"
	for i := 0; i < 5; i++ {
		allow, err := limiter.Allow(collection, fmt.Sprintf("req_%v", i))
		assert.Nil(t, err)
		assert.True(t, allow)
	}

	allow, err := limiter.Allow(collection, fmt.Sprintf("req_%v", 5))
	assert.Nil(t, err)
	assert.False(t, allow)

	time.Sleep(time.Minute)
	allow, err = limiter.Allow(collection, fmt.Sprintf("req_%v", 6))
	assert.Nil(t, err)
	assert.True(t, allow)

	card := cli.ZCard(collection).Val()
	assert.Equal(t, int64(1), card)
}
