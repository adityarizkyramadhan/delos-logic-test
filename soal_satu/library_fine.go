package soalsatu

import "time"

func FineCalculator(d1, m1, y1, d2, m2, y2 int) int {
	dueDate := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.Now().Location())
	returnDate := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.Now().Location())
	differenceDate := returnDate.Sub(dueDate)
	if differenceDate < 0 {
		return 0
	} else if m1 == m2 && y1 == y2 {
		dayLate := int(differenceDate.Hours() / 24)
		return 15 * dayLate
	} else if y1 == y2 {
		monthLate := m2 - m1
		return 500 * monthLate
	} else if y2 > y1 {
		return 12000
	}
	return -1
}
