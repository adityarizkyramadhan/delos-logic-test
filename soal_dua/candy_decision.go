package soaldua

func SourCandyDecision(student, candies, firstStudent int) int {
	for i := 1; i < candies; i++ {
		firstStudent++
		if firstStudent > student {
			firstStudent = 1
		}
	}
	return firstStudent
}
