package tokeninsight

func mergeSort(arr1, arr2 []int) []int {
	res := make([]int, 0, len(arr1)+len(arr2))

	var idx1, idx2 int
	for idx1 < len(arr1) && idx2 < len(arr2) {
		if arr1[idx1] < arr2[idx2] {
			res = append(res, arr1[idx1])
			idx1++
		} else {
			res = append(res, arr2[idx2])
			idx2++
		}
	}

	for idx1 < len(arr1) {
		res = append(res, arr1[idx1])
		idx1++
	}

	for idx2 < len(arr2) {
		res = append(res, arr2[idx2])
		idx2++
	}

	return res
}

// O(logk*SUM(n1, n2, .., nk))
