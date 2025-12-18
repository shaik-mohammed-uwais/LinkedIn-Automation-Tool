package storage

import (
	"encoding/json"
	"os"
	"sync"
)

// BotState represents everything the bot remembers
// between different runs.
type BotState struct {
	VisitedProfiles  map[string]bool `json:"visited_profiles"`
	SentConnections  map[string]bool `json:"sent_connections"`
	SentMessages     map[string]bool `json:"sent_messages"`
	ConnectionsToday int             `json:"connections_today"`
}

// StateStore is responsible for saving and loading bot state
// from a JSON file on disk.
type StateStore struct {
	filePath string
	state    *BotState
	mutex    sync.Mutex
}

// NewStateStore loads existing state or creates a fresh one.
func NewStateStore(filePath string) (*StateStore, error) {
	store := &StateStore{
		filePath: filePath,
	}

	if err := store.loadFromDisk(); err != nil {
		return nil, err
	}

	return store, nil
}

// loadFromDisk reads the state file or creates a new one if missing.
func (s *StateStore) loadFromDisk() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// First run: no state file exists
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		s.state = newEmptyState()
		return s.saveToDisk()
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	var loadedState BotState
	if err := json.Unmarshal(data, &loadedState); err != nil {
		return err
	}

	s.state = &loadedState
	return nil
}

// saveToDisk writes the current state back to the file.
func (s *StateStore) saveToDisk() error {
	data, err := json.MarshalIndent(s.state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

// newEmptyState returns a clean starting state.
func newEmptyState() *BotState {
	return &BotState{
		VisitedProfiles: make(map[string]bool),
		SentConnections: make(map[string]bool),
		SentMessages:    make(map[string]bool),
	}
}

//
// -------- Simple, human-readable helper methods --------
//

// HasVisited checks if we already opened this profile.
func (s *StateStore) HasVisited(profileURL string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.state.VisitedProfiles[profileURL]
}

// MarkVisited remembers that we visited this profile.
func (s *StateStore) MarkVisited(profileURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.state.VisitedProfiles[profileURL] = true
	return s.saveToDisk()
}

// HasSentConnection checks if a connection request was already sent.
func (s *StateStore) HasSentConnection(profileURL string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.state.SentConnections[profileURL]
}

// MarkConnectionSent records a sent connection request.
func (s *StateStore) MarkConnectionSent(profileURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.state.SentConnections[profileURL] = true
	s.state.ConnectionsToday++
	return s.saveToDisk()
}

// HasSentMessage checks if we already messaged this profile.
func (s *StateStore) HasSentMessage(profileURL string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.state.SentMessages[profileURL]
}

// MarkMessageSent records a successfully sent message.
func (s *StateStore) MarkMessageSent(profileURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.state.SentMessages[profileURL] = true
	return s.saveToDisk()
}
