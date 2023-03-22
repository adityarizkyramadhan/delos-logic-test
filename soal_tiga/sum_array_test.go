package soaltiga_test

import (
	"testing"

	soaltiga "github.com/adityarizkyramadhan/delos-logic-test/soal_tiga"
	"github.com/magiconair/properties/assert"
)

func Test_SumArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected string
	}{
		{
			name:     "Test case 1 Delos",
			arr:      []int{1, 5, 7, 2, 4},
			expected: "YES",
		},
		{
			name:     "Test case 2 Delos",
			arr:      []int{1, 3, 4, 9},
			expected: "NO",
		},
		{
			name:     "Test case 3 Delos",
			arr:      []int{1, 1, 4, 1, 1},
			expected: "YES",
		},
		{
			name:     "Test case 4 Delos",
			arr:      []int{2, 0, 0, 0},
			expected: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soaltiga.SumOfElements(tt.arr)
			assert.Equal(t, result, tt.expected, "test ini gagal hasil tidak sama")
		})
	}
}
