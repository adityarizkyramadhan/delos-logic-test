package soaldua_test

import (
	"testing"

	soaldua "github.com/adityarizkyramadhan/delos-logic-test/soal_dua"
	"github.com/magiconair/properties/assert"
)

func Test_CandyDecision(t *testing.T) {
	tests := []struct {
		name         string
		students     int
		candies      int
		firstStudent int
		expected     int
	}{
		// TODO: Add test cases.
		{
			name:         "Test case 1 Delos",
			students:     5,
			candies:      3,
			firstStudent: 1,
			expected:     3,
		},
		{
			name:         "Test case 2 Delos",
			students:     352926151,
			candies:      380324688,
			firstStudent: 94730870,
			expected:     122129406,
		}, {
			name:         "Test case 1 pribadi",
			students:     5,
			candies:      2,
			firstStudent: 2,
			expected:     3,
		},
		{
			name:         "Test case 2 pribadi",
			students:     3,
			candies:      10,
			firstStudent: 2,
			expected:     2,
		},
		{
			name:         "Test case 3 pribadi",
			students:     3,
			candies:      10,
			firstStudent: 1,
			expected:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soaldua.SourCandyDecision(tt.students, tt.candies, tt.firstStudent)
			assert.Equal(t, tt.expected, result, "test ini gagal hasil tidak sama")
		})
	}
}
