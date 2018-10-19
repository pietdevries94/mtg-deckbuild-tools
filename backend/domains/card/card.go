package card

func GetCardBySetAndNumber(set, number string) (Card, error) {
	return getCardBySetAndNumber(set, number)
}

func GetCardByScryfallID(scryfallID string) (Card, error) {
	return getCardByScryfallID(scryfallID)
}
