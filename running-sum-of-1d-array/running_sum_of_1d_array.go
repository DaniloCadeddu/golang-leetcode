package runningsumof1darray

// https://leetcode.com/problems/running-sum-of-1d-array/

func RunningSum(nums []int) []int {
	var output []int
	sum := 0
	for _, i := range nums {
		sum += i
		output = append(output, sum)
	}
	return output
}
