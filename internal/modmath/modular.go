package modmath

// ModInverse returns x such that (a*x) % m == 1, if it exists.
// ok=false when gcd(a,m) != 1 (no inverse).
func ModInverse(a, m int) (x int, ok bool) {
	// Extended Euclidean Algorithm
	t, newT := 0, 1
	r, newR := m, ((a%m)+m)%m

	for newR != 0 {
		q := r / newR
		t, newT = newT, t-q*newT
		r, newR = newR, r-q*newR
	}

	if r != 1 { // gcd != 1 → no inverse
		return 0, false
	}

	if t < 0 {
		t += m
	}
	return t, true
}
