package construct

/*

算法
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。


示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序


进阶：
请你设计时间复杂度为 O(n) 的算法解决本问题

*/

func squareOfNonDecreasingArray(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	splitIdx := 0
	for splitIdx < len(arr) {
		if arr[splitIdx] < 0 {
			splitIdx++
		} else {
			break
		}
	}

	leftIdx, rightIdx := splitIdx-1, splitIdx

	ans := make([]int, 0, len(arr))
	for leftIdx >= 0 && rightIdx < len(arr) {
		left, right := arr[leftIdx], arr[rightIdx]
		leftSquare, rightSquare := left*left, right*right
		if leftSquare < rightSquare {
			ans = append(ans, leftSquare)
			leftIdx--
		} else {
			ans = append(ans, rightSquare)
			rightIdx++
		}
	}

	for leftIdx >= 0 {
		left := arr[leftIdx]
		ans = append(ans, left*left)
		leftIdx--
	}

	for rightIdx < len(arr) {
		right := arr[rightIdx]
		ans = append(ans, right*right)
		rightIdx++
	}

	return ans
}
