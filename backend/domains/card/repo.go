package card

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

const columns = `scryfall_id, set_code, set_number, name, oracle_id, updated_at, thumbnail_url, casting_cost, online_price, copies_owned`

func getCardFromDB(q string) (Card, bool, error) {
	var db = data.GetDB()

	row := db.QueryRow(q)

	ok := true
	card := Card{}
	err := row.Scan(&card.ScryfallID, &card.SetCode, &card.SetNumber, &card.Name, &card.OracleID,
		&card.UpdatedAt, &card.ThumbnailURL, &card.CastingCost, &card.OnlinePrice, &card.CopiesOwned)
	if err != nil {
		if err != sql.ErrNoRows {
			return card, false, err
		}
		ok = false
	}

	if !ok {
		return card, false, nil
	}

	if card.UpdatedAt.Before(time.Now().Add(-24 * time.Hour)) {
		card, err := updateCardCacheFromScryfall(card.ScryfallID)
		return card, true, err
	}

	return card, true, nil
}

func getCardByScryfallID(scryfallID string) (Card, error) {
	q := fmt.Sprintf(`SELECT %s
		FROM cards
		WHERE scryfall_id = '%s'
		ORDER BY updated_at desc
		LIMIT 1;`, columns, scryfallID)

	card, ok, err := getCardFromDB(q)
	if err != nil {
		return card, err
	}

	if !ok {
		return addCardCacheFromScryfallByScryfallID(scryfallID)
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

func addCardCacheFromScryfallByName(name string) (Card, error) {
	sfClient := data.GetScryfall()

	ctx := context.Background()
	sfCard, err := sfClient.GetCardByName(ctx, name, true, scryfall.GetCardByNameOptions{})
	if err != nil {
		return Card{}, err
	}

	return insertSfCard(sfCard)
}

func getCardBySetAndNumber(set, number string) (Card, error) {
	q := fmt.Sprintf(`SELECT %s
		FROM cards
		WHERE set_code = '%s'
		AND set_number = '%s'`,
		columns, set, number)

	card, ok, err := getCardFromDB(q)
	if err != nil {
		return card, err
	}

	if !ok {
		return addCardCacheFromScryfallBySetAndNumber(set, number)
	}

	return card, nil
}

func getCardByName(name string) (Card, error) {
	q := fmt.Sprintf(`SELECT %s
		FROM cards
		WHERE name = '%s'`,
		columns, name)

	card, ok, err := getCardFromDB(q)
	if err != nil {
		return card, err
	}

	if !ok {
		return addCardCacheFromScryfallByName(name)
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

	q := `insert into cards (scryfall_id, set_code, set_number, name, oracle_id,
		updated_at, thumbnail_url, casting_cost, online_price) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, card.ScryfallID, card.SetCode, card.SetNumber, card.Name, card.OracleID,
		card.UpdatedAt, card.ThumbnailURL, card.CastingCost, card.OnlinePrice, card.CopiesOwned)

	return card, err
}

func updateSfCard(sfCard scryfall.Card) (Card, error) {
	var db = data.GetDB()

	card := sfCardToCard(sfCard)

	q := `update cards set 
			set_code = ?,
			set_number = ?,
			name = ?,
			oracle_id = ?,
			updated_at = ?,
			thumbnail_url = ?,
			casting_cost = ?,
			online_price = ?
		  where scryfall_id = ?`
	_, err := db.Exec(q, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt,
		card.ThumbnailURL, card.CastingCost, card.OnlinePrice, card.ScryfallID)

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
		CastingCost:  sfCard.ManaCost,
		OnlinePrice:  sfCard.EUR,
		UpdatedAt:    time.Now(),
	}
}
