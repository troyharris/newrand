package newrand

import "testing"

func TestIntr(t *testing.T) {
	TimeSeed()
	const start, end = 100, 1000
	for i := 0; i < 10000; i++ {
		a := Intr(100, 1000)
		if a < start || a > end {
			t.Errorf("%v is out of range", a)
		}
	}
}

func TestHit(t *testing.T) {
	TimeSeed()
	const p = 50
	count := 0
	for i := 0; i < 10000; i++ {
		v, err := Hit(p)
		if err != nil {
			t.Error("Percentage out of range")
			return
		}
		if v {
			count++
		}
	}
	rate := count / 100
	if rate > 55 || rate < 45 {
		t.Errorf("Hit rate out of range. Was %v pct.", rate)
	}
}
