package soaltiga

func SumOfElements(arr []int) string {
	total := sumArray(arr)
	leftSum := 0
	for i := range arr {
		if leftSum == total-leftSum-arr[i] {
			return "YES"
		}
		leftSum += arr[i]
	}
	return "NO"
}

func sumArray(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res
}
