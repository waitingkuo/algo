package insertion

import (
	"testing"
)

func TestSort(t *testing.T) {

	a := []int{6, 1, 4, 2, 3, 5}
	Sort(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Error()
		}
	}

}
