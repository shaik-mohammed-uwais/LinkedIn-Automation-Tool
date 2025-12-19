package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"linkedin-automation/logger"
)

// Login performs LinkedIn login with detection safeguards.
func Login(page *rod.Page, email, password string, log *logger.Logger) error {
	log.Info("Checking LinkedIn authentication state")

	// Always start from home (safe)
	if err := page.Navigate("https://www.linkedin.com/"); err != nil {
		return err
	}

	page.MustWaitLoad()
	page.MustWaitIdle()

	// Session already restored via cookies
	if IsLoggedIn(page) {
		log.Info("Already authenticated, skipping login")
		return nil
	}

	log.Info("Opening LinkedIn login page")
	if err := page.Navigate("https://www.linkedin.com/login"); err != nil {
		return err
	}
	page.WaitLoad()
	time.Sleep(2 * time.Second)

	// Human-like typing (basic)
	page.MustElement("#username").MustInput(email)
	time.Sleep(1 * time.Second)

	page.MustElement("#password").MustInput(password)
	time.Sleep(1 * time.Second)

	page.MustElement("button[type=submit]").MustClick()

	page.MustWaitLoad()
	page.MustWaitIdle()
	time.Sleep(3 * time.Second)

	currentURL := page.MustInfo().URL
	log.Debug("Post-login URL: " + currentURL)

	// Security detections
	if strings.Contains(currentURL, "checkpoint") {
		return errors.New("linkedin checkpoint triggered")
	}

	if strings.Contains(currentURL, "signup") {
		return errors.New("linkedin redirected to signup (blocked)")
	}

	has2FA, _, _ := page.Has("input[name=pin]")
	if has2FA {
		return errors.New("2FA detected")
	}

	hasCaptcha, _, _ := page.Has("iframe[src*='captcha']")
	if hasCaptcha {
		return errors.New("captcha detected")
	}

	// FINAL verification
	if !IsLoggedIn(page) {
		return fmt.Errorf("login failed: user not authenticated")
	}

	log.Info("Login successful")
	return nil
}
