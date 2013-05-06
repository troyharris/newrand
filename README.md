newrand
=======

Some simple extensions of math/rand in an effort to learn Go

Probably pretty useless.

Some functions: 

newrand.TimeSeed() is just an alias for: rand.Seed(time.Now().UnixNano())

newrand.Intr(int, int) is like rand.Intn, except it takes a start and end int as a range

newrand.Hit(int) takes an integer between 0 and 100 that represents a percent chance of "hitting" and returns true or false based on those odds. 
