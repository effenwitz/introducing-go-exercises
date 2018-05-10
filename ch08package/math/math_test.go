package math

import "testing"

//go includes special program go test
//go compiler ignores code ending with _test.go
func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverages(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error("For", pair.values,
				"expected", pair.average,
				"got", v)
		}
	}
}
