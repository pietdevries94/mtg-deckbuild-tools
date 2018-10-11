package card

import (
	"context"
	"database/sql"
	"time"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func getCardBySetAndNumber(set, number string) (Card, error) {
	db := data.GetDB()

	q := `SELECT
		scryfall_id,
		set_code,
		set_number,
		name,
		oracle_id,
		updated_at
	FROM cards
	WHERE set_code = ?
	AND set_number = ?`
	row := db.QueryRow(q, set, number)

	ok := true
	card := Card{}
	err := row.Scan(&card.ScryfallID, &card.SetCode, &card.SetNumber, &card.Name, &card.OracleID, &card.UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			return card, err
		}
		ok = false
	}

	if !ok {
		return addCardCacheFromScryfallBySetAndNumber(set, number)
	}

	if card.UpdatedAt.Before(time.Now().Add(-24 * time.Hour)) {
		return updateCardCacheFromScryfall(card.ScryfallID)
	}

	return card, nil
}

func addCardCacheFromScryfallBySetAndNumber(set, number string) (Card, error) {
	sfClient := data.GetScryfall()

	ctx := context.Background()
	sfCard, err := sfClient.GetCardBySetCodeAndCollectorNumber(ctx, set, string(number))
	if err != nil {
		return Card{}, err
	}

	return upsertSfCard(sfCard)
}

func updateCardCacheFromScryfall(scryfallID string) (Card, error) {
	sfClient := data.GetScryfall()

	ctx := context.Background()
	sfCard, err := sfClient.GetCard(ctx, scryfallID)
	if err != nil {
		return Card{}, err
	}

	return upsertSfCard(sfCard)
}

func upsertSfCard(sfCard scryfall.Card) (Card, error) {
	card := Card{
		ScryfallID: sfCard.ID,
		SetCode:    sfCard.Set,
		SetNumber:  sfCard.CollectorNumber,
		Name:       sfCard.Name,
		OracleID:   sfCard.OracleID,
		UpdatedAt:  time.Now(),
	}

	db := data.GetDB()

	q := `insert or replace into cards (scryfall_id, set_code, set_number, name, oracle_id, updated_at) values (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, card.ScryfallID, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt)

	return card, err
}
