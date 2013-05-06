package newrand

import (
	"fmt"
	"math/rand"
	"time"
)

func TimeSeed() {
	rand.Seed(time.Now().UnixNano())
}

func Intr(a, b int) int {
	c := b - a
	r := rand.Intn(c + 1)
	return r + a
}

func Hit(p int) (bool, error) {
	if p > 100 {
		return false, fmt.Errorf("Percentage is %v. Must be 100 or less", p)
	}
	r := rand.Intn(100)
	if p > r {
		return true, nil
	} else {
		return false, nil
	}
}
