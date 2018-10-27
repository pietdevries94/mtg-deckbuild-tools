package card

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/data"
)

const columns = `scryfall_id, set_code, set_number, name, oracle_id, updated_at,
	thumbnail_url, casting_cost, online_price, copies_owned, buy_link, type_line`

func getCardFromDB(q string) (Card, bool, error) {
	var db = data.GetDB()

	row := db.QueryRow(q)

	ok := true
	card := Card{}
	err := row.Scan(&card.ScryfallID, &card.SetCode, &card.SetNumber, &card.Name, &card.OracleID, &card.UpdatedAt,
		&card.ThumbnailURL, &card.CastingCost, &card.OnlinePrice, &card.CopiesOwned, &card.BuyLink, &card.TypeLine)
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
	escapedName := strings.Replace(name, `'`, `''`, -1)

	q := fmt.Sprintf(`SELECT %s
		FROM cards
		WHERE name = '%s'`,
		columns, escapedName)

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
		updated_at, thumbnail_url, casting_cost, online_price, copies_owned, buy_link, type_line) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, card.ScryfallID, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt,
		card.ThumbnailURL, card.CastingCost, card.OnlinePrice, card.CopiesOwned, card.BuyLink, card.TypeLine)

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
			online_price = ?,
			buy_link = ?,
			type_line = ?
		  where scryfall_id = ?`
	_, err := db.Exec(q, card.SetCode, card.SetNumber, card.Name, card.OracleID, card.UpdatedAt,
		card.ThumbnailURL, card.CastingCost, card.OnlinePrice, card.BuyLink, card.TypeLine, card.ScryfallID)

	return card, err
}

func sfCardToCard(sfCard scryfall.Card) Card {
	img := ``
	if sfCard.ImageURIs == nil {
		if len(sfCard.CardFaces) > 0 {
			img = sfCard.CardFaces[0].ImageURIs.Normal
		}
	} else {
		img = sfCard.ImageURIs.Normal
	}

	return Card{
		ScryfallID:   sfCard.ID,
		SetCode:      sfCard.Set,
		SetNumber:    sfCard.CollectorNumber,
		Name:         sfCard.Name,
		OracleID:     sfCard.OracleID,
		ThumbnailURL: img,
		CastingCost:  sfCard.ManaCost,
		OnlinePrice:  sfCard.EUR,
		BuyLink:      sfCard.PurchaseURIs.MagicCardMarket,
		TypeLine:     sfCard.TypeLine,
		UpdatedAt:    time.Now(),
	}
}
