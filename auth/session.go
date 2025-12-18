package auth

import (
	"time"

	"github.com/go-rod/rod"
)

// IsLoggedIn checks if LinkedIn authenticated UI is present.
func IsLoggedIn(page *rod.Page) bool {
	_, err := page.Timeout(5 * time.Second).
		Element("nav.global-nav")
	return err == nil
}
