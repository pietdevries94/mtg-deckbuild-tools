package entry

type Entry struct {
	ID         int      `json:"id"`
	ScryfallID string   `json:"scryfall_id"`
	ListID     int      `json:"list_id"`
	Tags       []string `json:"tags"`
}

type AddEntryPayload struct {
	ScryfallID string   `json:"scryfall_id"`
	ListID     int      `json:"list_id"`
	Tags       []string `json:"tags"`
}

type DeleteEntryPayload struct {
	ScryfallID string `json:"scryfall_id"`
	ListID     int    `json:"list_id"`
}
