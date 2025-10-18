package construct

import (
	"reflect"
	"testing"
)

func Test_squareOfNonDecreasingArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "all positive elements",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "all negative elements",
			args: args{
				arr: []int{-5, -4, -3, -2, -1},
			},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "positive and negative",
			args: args{
				arr: []int{-3, -2, -1, 4, 5},
			},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "negative, positive and zero",
			args: args{
				arr: []int{-3, -2, -1, 0, 4, 5},
			},
			want: []int{0, 1, 4, 9, 16, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareOfNonDecreasingArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("squareOfNonDecreasingArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
