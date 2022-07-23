package findpivotindex

// https://leetcode.com/problems/find-pivot-index/

func PivotIndex(nums []int) int {
	for i := range nums {
		rightSum, leftSum := calculateRightAndLeftSum(nums, i)
		if rightSum == leftSum {
			return i
		}
	}
	return -1
}

func calculateRightAndLeftSum(nums []int, pivot int) (int, int) {
	rightSum := sum(nums[0:pivot])
	leftSum := sum(nums[pivot+1:])

	return rightSum, leftSum
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
