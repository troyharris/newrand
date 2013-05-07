package newrand

import (
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

func Hit(p int) bool {
	if p > 100 {
		p = 100
	}
	r := rand.Intn(100)
	if p >= r {
		return true
	} else {
		return false
	}
}
