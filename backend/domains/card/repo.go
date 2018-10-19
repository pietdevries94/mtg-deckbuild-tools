package card

import (
	"context"
	"database/sql"
	"time"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

func getCardByScryfallID(scryfallID string) (Card, error) {
	var db = data.GetDB()

	q := `SELECT
		scryfall_id,
		set_code,
		set_number,
		name,
		oracle_id,
		updated_at,
		thumbnail_url
	FROM cards
	WHERE scryfall_id = ?
	ORDER BY updated_at desc
	LIMIT 1;`
	row := db.QueryRow(q, scryfallID)

	ok := true
	card := Card{}
	err := row.Scan(&card.ScryfallID, &card.SetCode, &card.SetNumber, &card.Name, &card.OracleID, &card.UpdatedAt, &card.ThumbnailURL)
	if err != nil {
		if err != sql.ErrNoRows {
			return card, err
		}
		ok = false
	}

	if !ok {
		return addCardCacheFromScryfallByScryfallID(scryfallID)
	}

	if card.UpdatedAt.Before(time.Now().Add(-24 * time.Hour)) {
		return updateCardCacheFromScryfall(card.ScryfallID)
	}

	return card, nil
}

func addCardCacheFromScryfallByScryfallID(scryfallID string) (Card, error) {
	sfClient := data.GetScryfall()

	ctx := context.Background()
	sfCard, err := sfClient.GetCard(ctx, scryfallID)
	if err != nil {
		return Card{}, err
	}

	return insertSfCard(sfCard)
}

func getCardBySetAndNumber(set, number string) (Card, error) {
	var db = data.GetDB()

	q := `SELECT
		scryfall_id,
		set_code,
		set_number,
		name,
		oracle_id,
		updated_at,
		thumbnail_url
	FROM cards
	WHERE set_code = ?
	AND set_number = ?`
	row := db.QueryRow(q, set, number)

	ok := true
	card := Card{}
	err := row.Scan(&card.ScryfallID, &card.SetCode, &card.SetNumber, &card.Name, &card.OracleID, &card.UpdatedAt, &card.ThumbnailURL)
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

	return insertSfCard(sfCard)
}

func updateCardCacheFromScryfall(scryfallID string) (Card, error) {
	sfClient := data.GetScryfall()

	ctx := context.Background()
	sfCard, err := sfClient.GetCard(ctx, scryfallID)
	if err != nil {
		return Card{}, err
	}

	return updateSfCard(sfCard)
}

func insertSfCard(sfCard scryfall.Card) (Card, error) {
	var db = data.GetDB()

	card := sfCardToCard(sfCard)

	q := `insert into cards (scryfall_id, set_code, set_number, name, oracle_id, updated_at, thumbnail_url) values (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, card.ScryfallID, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt, card.ThumbnailURL)

	return card, err
}

func updateSfCard(sfCard scryfall.Card) (Card, error) {
	var db = data.GetDB()

	card := sfCardToCard(sfCard)

	q := `update cards set set_code = ?, set_number = ?, name = ?, oracle_id = ?, updated_at = ?, thumbnail_url = ?, where scryfall_id = ?`
	_, err := db.Exec(q, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt, card.ThumbnailURL, card.ScryfallID)

	return card, err
}

func sfCardToCard(sfCard scryfall.Card) Card {
	return Card{
		ScryfallID:   sfCard.ID,
		SetCode:      sfCard.Set,
		SetNumber:    sfCard.CollectorNumber,
		Name:         sfCard.Name,
		OracleID:     sfCard.OracleID,
		ThumbnailURL: sfCard.ImageURIs.Normal,
		UpdatedAt:    time.Now(),
	}
}
