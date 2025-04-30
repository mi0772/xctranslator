package model

import "sync"

type XCStrings struct {
	SourceLanguage string                   `json:"sourceLanguage"`
	Strings        map[string]XCStringEntry `json:"strings"`
	Version        string                   `json:"version"`
}

type XCStringEntry struct {
	Localizations map[string]LocalizedString `json:"localizations,omitempty"`
}

type LocalizedString struct {
	StringUnit StringUnit `json:"stringUnit"`
}

type StringUnit struct {
	State string `json:"state"`
	Value string `json:"value"`
}
type TranslationContext struct {
	mu                sync.Mutex
	SourceLanguage    string
	SourceStrings     []StringEntry
	translatedStrings map[string][]StringEntry
}

func (t *TranslationContext) AddTranslatedString(lang string, entry StringEntry) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.translatedStrings == nil {
		t.translatedStrings = make(map[string][]StringEntry)
	}

	if _, exists := t.translatedStrings[lang]; !exists {
		t.translatedStrings[lang] = []StringEntry{}
	}

	t.translatedStrings[lang] = append(t.translatedStrings[lang], entry)
}

func (t *TranslationContext) GetTranslatedStrings() map[string][]StringEntry {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.translatedStrings
}

type StringEntry struct {
	Identifier int64
	Content    string
}
