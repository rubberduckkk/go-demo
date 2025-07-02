package tokeninsight

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type data struct {
		name       string
		arr1, arr2 []int
		wantRes    []int
	}

	testData := []data{
		{
			name:    "normal case 1",
			arr1:    []int{0, 2, 8},
			arr2:    []int{3, 7, 9, 11},
			wantRes: []int{0, 2, 3, 7, 8, 9, 11},
		},
		{
			name:    "nil input",
			arr1:    nil,
			arr2:    []int{},
			wantRes: []int{},
		},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			res := mergeSort(d.arr1, d.arr2)
			if !reflect.DeepEqual(res, d.wantRes) {
				t.Errorf("wrong result: %v, expected: %v\n", res, d.wantRes)
			}
		})
	}
}
