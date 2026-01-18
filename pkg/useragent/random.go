package ua

// xorshift64 - fast PRNG with 64-bit state
// Period: 2^64-1, passes BigCrush
type xorshift64 struct {
	state uint64
}

func newXorshift64(seed uint64) *xorshift64 {
	if seed == 0 { seed = 0xDEADBEEFCAFEBABE } // default seed
	return &xorshift64{state: seed}
}

func (x *xorshift64) next() uint64 {
	x.state ^= x.state << 13
	x.state ^= x.state >> 7
	x.state ^= x.state << 17
	return x.state
}

// intn returns pseudo-random int in [0, n)
func (x *xorshift64) intn(n int) int {
	if n <= 0 { return 0 }
	return int(x.next() % uint64(n))
}

// pick returns random element from slice s.
// s must be non-empty.
func pick[T any](x *xorshift64, s []T) T {
	return s[x.intn(len(s))]
}
