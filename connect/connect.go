package connect

import (
	"errors"

	"github.com/go-rod/rod"

	"linkedin-automation/logger"
	"linkedin-automation/stealth"
)

// Connector holds state related to connection requests
type Connector struct {
	DailyLimit int
	SentToday  int
	Logger     *logger.Logger
}

// NewConnector creates a new connection manager
func NewConnector(limit int, log *logger.Logger) *Connector {
	return &Connector{
		DailyLimit: limit,
		SentToday:  0,
		Logger:     log,
	}
}

// SendConnectionRequest visits a profile and tries to send a connection request
func (c *Connector) SendConnectionRequest(
	page *rod.Page,
	profileURL string,
	note string,
) error {

	if c.SentToday >= c.DailyLimit {
		return errors.New("daily connection limit reached")
	}

	c.Logger.Info("Visiting profile: " + profileURL)

	err := page.Navigate(profileURL)
	if err != nil {
		return err
	}

	stealth.RandomDelay(3, 6)

	// Try to find the Connect button
	connectButton, err := page.Element(`button:contains("Connect")`)
	if err != nil {
		c.Logger.Warn("Connect button not found — skipping profile")
		return nil
	}

	c.Logger.Info("Clicking Connect button")
	connectButton.MustClick()

	stealth.RandomDelay(1, 2)

	// Optional note flow
	if note != "" {
		c.Logger.Info("Adding a personalized note")

		addNoteButton := page.MustElement(`button:contains("Add a note")`)
		addNoteButton.MustClick()

		stealth.RandomDelay(1, 2)

		noteBox := page.MustElement(`textarea`)
		stealth.TypeLikeHuman(noteBox, note)
	}

	stealth.RandomDelay(1, 2)

	sendButton := page.MustElement(`button:contains("Send")`)
	sendButton.MustClick()

	c.SentToday++
	c.Logger.Info("Connection request sent successfully")

	stealth.RandomDelay(5, 10)

	return nil
}
