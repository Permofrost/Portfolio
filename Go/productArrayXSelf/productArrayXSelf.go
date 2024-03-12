package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i] contains the product of all the numbers to the left of i
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// Now for each i, compute product of all the numbers to the right and multiply it with the product of all the numbers to the left
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}

	return answer
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}
