package card

import "time"

type Card struct {
	ScryfallID string
	SetCode    string
	SetNumber  string
	Name       string
	OracleID   string
	UpdatedAt  time.Time
}
