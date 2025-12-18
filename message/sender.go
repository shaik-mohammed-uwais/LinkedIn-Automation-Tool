package message

import (
	"errors"

	"github.com/go-rod/rod"

	"linkedin-automation/logger"
	"linkedin-automation/stealth"
)

// MessageSender handles sending LinkedIn messages
type MessageSender struct {
	Logger *logger.Logger
}

// NewMessageSender creates a new message sender
func NewMessageSender(log *logger.Logger) *MessageSender {
	return &MessageSender{
		Logger: log,
	}
}

// SendMessage sends a message to an existing connection
func (m *MessageSender) SendMessage(
	page *rod.Page,
	profileURL string,
	messageText string,
) error {

	m.Logger.Info("Opening profile for messaging")

	page.MustNavigate(profileURL)
	stealth.RandomDelay(3, 6)

	messageButton, err := page.Element(`button:contains("Message")`)
	if err != nil {
		m.Logger.Warn("Message button not found — skipping")
		return errors.New("cannot message this profile")
	}

	m.Logger.Info("Opening message dialog")
	messageButton.MustClick()

	stealth.RandomDelay(2, 4)

	messageBox := page.MustElement(`div[role="textbox"]`)
	stealth.TypeLikeHuman(messageBox, messageText)

	stealth.RandomDelay(1, 2)

	sendButton := page.MustElement(`button:contains("Send")`)
	sendButton.MustClick()

	m.Logger.Info("Message sent successfully")

	stealth.RandomDelay(4, 7)

	return nil
}
