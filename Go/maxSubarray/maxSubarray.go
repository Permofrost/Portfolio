package main

func maxSubArray(nums []int) int {
	return maxSubArrayDivideAndConquer(nums, 0, len(nums)-1)
}

func maxSubArrayDivideAndConquer(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	maxLeft := maxSubArrayDivideAndConquer(nums, left, mid)
	maxRight := maxSubArrayDivideAndConquer(nums, mid+1, right)
	maxCross := maxCrossingSubArray(nums, left, mid, right)
	return max(maxLeft, max(maxRight, maxCross))
}

func maxCrossingSubArray(nums []int, left, mid, right int) int {
	leftSum := nums[mid]
	leftMaxSum := leftSum
	for i := mid - 1; i >= left; i-- {
		leftSum += nums[i]
		leftMaxSum = max(leftMaxSum, leftSum)
	}
	rightSum := 0
	rightMaxSum := rightSum
	for i := mid + 1; i <= right; i++ {
		rightSum += nums[i]
		rightMaxSum = max(rightMaxSum, rightSum)
	}
	return leftMaxSum + rightMaxSum
}
