package slidingMinMax

import (
	"github.com/gammazero/deque"
	"golang.org/x/exp/constraints"
)

type data interface {
	constraints.Ordered // we can work with any ordered datatype
}

type RollingMinMax[D data] struct {
	k   int            // lookback window
	min deque.Deque[D] // queue storing min values
	max deque.Deque[D] // queus storing max values
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

	for _, number := range y {
		for t.max.Len() >= t.k { // constraint 1
			t.max.PopFront()
		}
		for t.max.Len() > 0 && t.max.Front() < number { // constraint 2
			t.max.PopFront()
		}
		t.max.PushBack(number) // append new value

		for t.min.Len() >= t.k { // constraint 1
			t.min.PopFront()
		}
		for t.min.Len() > 0 && t.min.Front() > number { // constraint 2
			t.min.PopFront()
		}
		t.min.PushBack(number) // append new value
	}
	return t.min.Front(), t.max.Front()
}
