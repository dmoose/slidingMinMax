package slidingMinMax

import "testing"

func TestNewRollingMinMax(t *testing.T) {
	m := NewRollingMinMax[float32](4)
	min, max := m.push(1.1, 2.1, 3.1, 3.1, 5.1, 3.1, 3.1, 2.1, 2.1, 5.1, 6.1, 3.1, 2.1, 1.1, 1.1, 1.1, 1.1, 1.1)
	if (min != 1.1) || (max != 1.1) {
		t.Errorf("Min: expected 1.1, got %v - Max: expected 1.1, got %v", min, max)
	}
	n := NewRollingMinMax[string](3)
	mins, maxs := n.push("a", "b", "c", "x", "y", "z", "a", "b", "c")
	if (mins != "a") || (maxs != "c") {
		t.Errorf("Min: expected 'a', got %v - Max: expected 'c', got %v", mins, maxs)
	}
}
