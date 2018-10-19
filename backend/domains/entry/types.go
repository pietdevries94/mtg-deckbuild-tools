package entry

type Entry struct {
	ScryfallID string
}

type AddEntryPayload struct {
	ScryfallID string `json:"scryfall_id"`
}
