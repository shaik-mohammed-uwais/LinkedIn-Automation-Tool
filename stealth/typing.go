package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// TypeLikeHuman types text one character at a time
// with small random delays between keystrokes.
func TypeLikeHuman(element *rod.Element, text string) {
	for _, character := range text {
		element.MustInput(string(character))

		// Random typing delay (50ms - 150ms)
		delay := rand.Intn(100) + 50
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
