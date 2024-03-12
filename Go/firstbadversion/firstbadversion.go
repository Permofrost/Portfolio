package main

var badVersion int = 4 // Set this to the first bad version

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
