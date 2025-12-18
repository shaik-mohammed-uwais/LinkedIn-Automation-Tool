package stealth

import (
	"math/rand"
	"time"
)

// RandomDelay pauses execution for a random duration
// between min and max seconds.
func RandomDelay(minSeconds int, maxSeconds int) {
	if minSeconds >= maxSeconds {
		time.Sleep(time.Duration(minSeconds) * time.Second)
		return
	}

	randomSeconds := rand.Intn(maxSeconds-minSeconds+1) + minSeconds
	time.Sleep(time.Duration(randomSeconds) * time.Second)
}

// ShortPause simulates a small human hesitation.
func ShortPause() {
	RandomDelay(1, 2)
}

// MediumPause simulates reading or thinking time.
func MediumPause() {
	RandomDelay(3, 5)
}

// LongPause simulates longer breaks or distractions.
func LongPause() {
	RandomDelay(6, 10)
}
