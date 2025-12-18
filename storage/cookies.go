package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// CookieStore handles browser session persistence.
type CookieStore struct {
	filePath string
}

// NewCookieStore creates a cookie store.
func NewCookieStore(filePath string) *CookieStore {
	return &CookieStore{filePath: filePath}
}

// SaveCookies saves all browser cookies to disk.
func (c *CookieStore) SaveCookies(page *rod.Page) error {
	cookies, err := page.Cookies([]string{})
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.filePath, data, 0644)
}

// LoadCookies loads cookies into the browser session.
func (c *CookieStore) LoadCookies(page *rod.Page) error {
	if _, err := os.Stat(c.filePath); os.IsNotExist(err) {
		return nil // first run is OK
	}

	data, err := os.ReadFile(c.filePath)
	if err != nil {
		return err
	}

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return err
	}

	var params []*proto.NetworkCookieParam
	for _, c := range storedCookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	return page.SetCookies(params)
}
