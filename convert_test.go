package stringtofloat

import "testing"

func TestConvertor(t *testing.T) {
	tests := []struct {
		String   string
		Excepted float64
	}{
		{"1.234,87", 1234.87},
		{"1,234.87", 1234.87},
		{"10.000,00", 10000.00},
		{"10,000.00", 10000.00},
		{"100,000,000.00", 100000000.00},
		{"100.000.000,00", 100000000.00},
		{"2024", 2024.00},
	}

	for _, test := range tests {
		val, err := Convert(test.String)
		if err != nil {
			t.Errorf("Error when trying to convert %s %v \n", test.String, err)
		}
		if val != test.Excepted {
			t.Errorf("Trying to convert %s to float. Got %.3f  Wanted: %.3f \n", test.String, val, test.Excepted)
		}
	}
}
