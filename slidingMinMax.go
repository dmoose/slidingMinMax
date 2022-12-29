package slidingMinMax

import (
	"github.com/gammazero/deque"
	"golang.org/x/exp/constraints"
)

type data interface {
	constraints.Ordered // we can work with any ordered datatype
}

type pair[D data] struct {
	index int64
	value D
}

type RollingMinMax[D data] struct {
	k   int                  // lookback window
	i   int64                // count of processed items
	min deque.Deque[pair[D]] // queue storing min values
	max deque.Deque[pair[D]] // queue storing max values
}

// NewRollingMinMax returns an initialized struct of appropriate type
func NewRollingMinMax[D data](k int) *RollingMinMax[D] {
	r := RollingMinMax[D]{}
	r.k = k
	return &r
}

// Push will accept a slice of values and add them to the queues returning the
// min and max after the push. This does mean that if you push more values than
// the window you could miss intermediate highs and lows since only the state of
// the data structure at the end of the operation is returned.
func (t *RollingMinMax[D]) push(y ...D) (D, D) {
	var p pair[D]
	for _, number := range y {
		t.i++
		p.index = t.i
		p.value = number
		for t.max.Len() > 0 && t.max.Front().index <= (t.i-int64(t.k)) { // constraint 1
			t.max.PopFront()
		}
		for t.max.Len() > 0 && t.max.Back().value < number { // constraint 2
			t.max.PopBack()
		}
		t.max.PushBack(p) // append new value

		for t.min.Len() > 0 && t.min.Front().index <= (t.i-int64(t.k)) { // constraint 1
			t.min.PopFront()
		}
		for t.min.Len() > 0 && t.min.Back().value > number { // constraint 2
			t.min.PopBack()
		}
		t.min.PushBack(p) // append new value
	}
	return t.min.Front().value, t.max.Front().value
}
