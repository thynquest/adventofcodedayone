package fuel

import "testing"

func TestFuelRequired(t *testing.T) {
	data := []struct {
		mass     float64
		expected float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
		{-15, 0},
	}
	for _, v := range data {
		result := Required(v.mass)
		if result != v.expected {
			t.Errorf("expected %v got %v", v.expected, result)
		}
	}
}

func TestTotalFuelRequired(t *testing.T) {
	data := []float64{12, 14, 1969, 100756, -15}
	expected := 34241
	var total float64
	for _, v := range data {
		total += Required(v)
	}
	if total != float64(expected) {
		t.Errorf("expected %v got %v\n", expected, total)
	}
}
