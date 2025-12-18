package search

import (
	"strings"

	"github.com/go-rod/rod"

	"linkedin-automation/logger"
	"linkedin-automation/stealth"
)

// SearchOptions defines what kind of profiles we are looking for
type SearchOptions struct {
	Keywords string
	Location string
	Company  string
}

// Searcher handles profile discovery
type Searcher struct {
	Logger *logger.Logger
	Seen   map[string]bool
}

// NewSearcher creates a new search manager
func NewSearcher(log *logger.Logger) *Searcher {
	return &Searcher{
		Logger: log,
		Seen:   make(map[string]bool),
	}
}

// FindProfiles searches LinkedIn and returns profile URLs
func (s *Searcher) FindProfiles(
	page *rod.Page,
	options SearchOptions,
	maxResults int,
) ([]string, error) {

	var profiles []string

	s.Logger.Info("Starting LinkedIn profile search")

	// Build a simple search URL
	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" +
		strings.ReplaceAll(options.Keywords, " ", "%20")

	page.MustNavigate(searchURL)
	stealth.RandomDelay(4, 7)

	for len(profiles) < maxResults {

		s.Logger.Info("Scanning search results page")

		profileLinks := page.MustElements(`a[href*="/in/"]`)

		for _, link := range profileLinks {
			href := link.MustAttribute("href")
			if href == nil {
				continue
			}

			profileURL := strings.Split(*href, "?")[0]

			if s.Seen[profileURL] {
				continue
			}

			s.Seen[profileURL] = true
			profiles = append(profiles, profileURL)

			s.Logger.Debug("Found profile: " + profileURL)

			if len(profiles) >= maxResults {
				break
			}
		}

		// Try moving to next page
		nextButton, err := page.Element(`button[aria-label="Next"]`)
		if err != nil {
			s.Logger.Info("No more pages available")
			break
		}

		s.Logger.Info("Moving to next page")
		nextButton.MustClick()

		stealth.RandomDelay(3, 6)
	}

	s.Logger.Info("Profile search completed")

	return profiles, nil
}
