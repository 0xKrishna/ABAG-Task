package main

func Solution(A []int) int {
	var waysToCut int
	if astheticTrees(A) {
		return 0
	}

	for i := 0; i < len(A)-1; i++ {
		a := createSubArray(A, 0, len(A)-1)
		a = append(a[:i], a[i+1:]...)
		if astheticTrees(a) {
			waysToCut++
		}
	}
	if waysToCut == 0 {
		return -1
	}
	return waysToCut
}

func createSubArray(a []int, start, end int) []int {
	var subArray []int
	for i := start; i <= end; i++ {
		subArray = append(subArray, a[i])
	}
	return subArray
}

func astheticTrees(a []int) bool {
	for i := 1; i < len(a)-1; i++ {
		if (a[i] >= a[i-1] && a[i] <= a[i+1]) || (a[i] <= a[i-1] && a[i] >= a[i+1]) {
			return false
		}
	}
	return true
}
