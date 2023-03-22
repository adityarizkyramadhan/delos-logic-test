package soalsatu_test

import (
	"testing"

	soalsatu "github.com/adityarizkyramadhan/delos-logic-test/soal_satu"
	"github.com/stretchr/testify/assert"
)

func Test_LibraryFine(t *testing.T) {
	tests := []struct {
		name                   string
		expected               int
		d1, m1, y1, d2, m2, y2 int
	}{
		{
			name:     "tepat waktu",
			expected: 0,
			d1:       15, m1: 8, y1: 2022,
			d2: 15, m2: 8, y2: 2022,
		},
		{
			name:     "telat selama 5 hari",
			expected: 75,
			d1:       15, m1: 8, y1: 2022,
			d2: 20, m2: 8, y2: 2022,
		},
		{
			name:     "telat selama 9 bulan",
			expected: 4500,
			d1:       15, m1: 1, y1: 2022,
			d2: 15, m2: 10, y2: 2022,
		},
		{
			name:     "telat selama 1 tahun",
			expected: 12000,
			d1:       15, m1: 8, y1: 2022,
			d2: 15, m2: 8, y2: 2023,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soalsatu.FineCalculator(tt.d1, tt.m1, tt.y1, tt.d2, tt.m2, tt.y2)
			assert.Equal(t, tt.expected, result, "test ini sukses, hasil sama dengan ekspetasi")
		})
	}
}
