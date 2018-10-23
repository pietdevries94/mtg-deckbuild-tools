package entry

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/card"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/tag"
)

func AddCardEntry(payload AddEntryPayload) error {
	_, err := card.GetCardByScryfallID(payload.ScryfallID)
	if err != nil {
		return err
	}

	entry := Entry{
		ScryfallID: payload.ScryfallID,
		ListID:     payload.ListID,
	}

	id, ok := getEntryID(entry)
	if !ok {
		id, err = insertEntry(entry)
		if err != nil {
			return err
		}
	}

	return addTags(id, payload.Tags)
}

func DeleteCardEntry(payload DeleteEntryPayload) error {
	entry := Entry{
		ScryfallID: payload.ScryfallID,
		ListID:     payload.ListID,
	}

	id, ok := getEntryID(entry)
	if !ok {
		return nil
	}

	deleteEntry(id)

	return nil
}

func GetEntriesForCard(scryfallID string) ([]Entry, error) {
	entries, err := getEntriesForCard(scryfallID)
	if err != nil {
		return nil, err
	}

	for i, entry := range entries {
		tags, err := tag.GetTagsForEntryID(entry.ID)
		if err != nil {
			return nil, err
		}
		entries[i].Tags = tags
	}

	return entries, nil
}
