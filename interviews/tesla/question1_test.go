package tesla

import (
	"testing"
)

func Test_reverseDigits(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				N: 100400,
			},
		},
		{
			name: "2",
			args: args{
				N: 98456001,
			},
		},
		{
			name: "3",
			args: args{
				N: 10300400500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseDigits(tt.args.N)
		})
	}
}
