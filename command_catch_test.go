package main

import "testing"

func TestCatchProbabilityDecreasesWithExperience(t *testing.T) {
	prev := catchProbability(0)
	if prev != 1 {
		t.Fatalf("expected certain catch at 0 xp, got %v", prev)
	}
	for _, xp := range []int{50, 112, 300, 608} {
		p := catchProbability(xp)
		if p <= 0 || p >= prev {
			t.Errorf("xp %d: probability %v should be in (0, %v)", xp, p, prev)
		}
		prev = p
	}
}

func TestAttemptCatch(t *testing.T) {
	cases := []struct {
		xp   int
		roll float64
		want bool
	}{
		{100, 0.49, true},
		{100, 0.5, false},
		{300, 0.24, true},
		{300, 0.3, false},
		{0, 0.999, true},
	}
	for _, c := range cases {
		if got := attemptCatch(c.xp, c.roll); got != c.want {
			t.Errorf("attemptCatch(%d, %v) = %v, want %v", c.xp, c.roll, got, c.want)
		}
	}
}
