package main

import (
	"fmt"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation/auth"
	"linkedin-automation/config"
	"linkedin-automation/connect"
	"linkedin-automation/logger"
	"linkedin-automation/message"
	"linkedin-automation/search"
	"linkedin-automation/storage"
)

func main() {
	// 1. Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	// 2. Logger
	log := logger.NewLogger(cfg.LogLevel)
	log.Info("LinkedIn automation bot starting")

	// 3. Storage
	store, err := storage.NewStateStore("state.json")
	if err != nil {
		log.Error("Failed to initialize storage")
		os.Exit(1)
	}

	cookieStore := storage.NewCookieStore("cookies.json")

	// 4. Browser (FIX: Disable Leakless)
	l := launcher.New().
		Headless(cfg.Headless).
		Leakless(false)

	browserURL, err := l.Launch()
	if err != nil {
		log.Error("Failed to launch browser: " + err.Error())
		os.Exit(1)
	}

	browser := rod.New().
		ControlURL(browserURL).
		MustConnect()
	defer browser.Close()

	page := browser.MustPage()

	//  Load cookies BEFORE login
	if err := cookieStore.LoadCookies(page); err != nil {
		log.Warn("Failed to load cookies: " + err.Error())
	}

	// Check if already logged in
	if auth.IsLoggedIn(page) {
		log.Info("Existing session detected, skipping login")
	} else {
		if err := auth.Login(page, cfg.Email, cfg.Password, log); err != nil {
			log.Error("Login failed: " + err.Error())
			return
		}

		// Save cookies AFTER successful login
		if err := cookieStore.SaveCookies(page); err != nil {
			log.Warn("Failed to save cookies: " + err.Error())
		}
	}

	// 5. Login
	if err := auth.Login(page, cfg.Email, cfg.Password, log); err != nil {
		log.Error("Login failed: " + err.Error())
		return
	}

	// 6. Search
	searcher := search.NewSearcher(log)
	profiles, err := searcher.FindProfiles(
		page,
		search.SearchOptions{
			Keywords: "Software Engineer",
		},
		5,
	)
	if err != nil {
		log.Error("Search failed: " + err.Error())
		return
	}

	// 7. Connect
	connector := connect.NewConnector(cfg.DailyConnectLimit, log)

	for _, profile := range profiles {
		if store.HasVisited(profile) {
			continue
		}

		if err := connector.SendConnectionRequest(
			page,
			profile,
			"Hi, would love to connect!",
		); err != nil {
			log.Warn(err.Error())
			continue
		}

		store.MarkVisited(profile)
		store.MarkConnectionSent(profile)
	}

	// 8. Message
	messenger := message.NewMessageSender(log)

	for _, profile := range profiles {
		if store.HasSentMessage(profile) {
			continue
		}

		if err := messenger.SendMessage(
			page,
			profile,
			"Thanks for connecting! Looking forward to learning from you.",
		); err == nil {
			store.MarkMessageSent(profile)
		}
	}

	log.Info("Bot run completed safely")
}
