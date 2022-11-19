package sum

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	if got != 3 {
		t.Errorf("sum(1,2) = %d; want 3", got)
	}
}
