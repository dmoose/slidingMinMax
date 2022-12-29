package slidingMinMax

import (
	"fmt"
	"testing"
)

// minmax iterate over slice and get min and max
func minmax[D data](v []D) (D, D) {
	min := v[0]
	max := v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println(v, min, max)
	return min, max
}

func TestRollingMinMaxSlice(t *testing.T) {
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
func TestRollingMinMaxFloat32(t *testing.T) {
	window := 4
	m := NewRollingMinMax[float32](window)
	m1 := []float32{1.1, 2.1, 3.1, 3.1, 5.1, 3.1, 3.1, 2.1, 2.1, 5.1, 6.1, 3.1, 2.1, 1.1, 1.1, 1.1, 1.1, 1.1}
	var mmax, mmin float32
	var tmin, tmax float32
	start := 0
	for index, number := range m1 {
		mmin, mmax = m.push(number)
		if start <= (index - window) { // leave slice at 0 until we get to window length
			start++
		}
		tmin, tmax = minmax(m1[start : index+1]) // call our test function on current window slice
		if mmin != tmin {
			t.Errorf("Min: expected %v, got %v", tmin, mmin)
		}
		if mmax != tmax {
			t.Errorf("Max: expected %v, got %v", tmax, mmax)
		}
	}
}
func TestRollingMinMaxInt32(t *testing.T) {
	window := 4
	m := NewRollingMinMax[int32](window)
	m1 := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 0, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 22, 100, 43, 19, 22, 22, 22, 22, 22, 3, 3, 3, 3, 3, 20}
	var mmax, mmin int32
	var tmin, tmax int32
	start := 0
	for index, number := range m1 {
		mmin, mmax = m.push(number)
		if start <= (index - window) { // leave slice at 0 until we get to window length
			start++
		}
		tmin, tmax = minmax(m1[start : index+1]) // call our test function on current window slice
		if mmin != tmin {
			t.Errorf("Min: expected %v, got %v", tmin, mmin)
		}
		if mmax != tmax {
			t.Errorf("Max: expected %v, got %v", tmax, mmax)
		}
	}
}
func TestRollingMinMaxString(t *testing.T) {
	window := 3
	n := NewRollingMinMax[string](window)
	n1 := []string{"a", "b", "c", "x", "y", "z", "a", "b", "c"}
	var nmax, nmin string
	var tnmin, tnmax string
	start := 0
	for index, number := range n1 {
		nmin, nmax = n.push(number)
		if start <= (index - window) { // leave slice at 0 until we get to window length
			start++
		}
		tnmin, tnmax = minmax(n1[start : index+1]) // call our test function on current window slice
		if nmin != tnmin {
			t.Errorf("Min: expected %v, got %v", tnmin, nmin)
		}
		if nmax != tnmax {
			t.Errorf("Max: expected %v, got %v", tnmax, nmax)
		}
	}
}
