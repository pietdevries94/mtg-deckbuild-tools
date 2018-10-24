package card

func GetCardBySetAndNumber(set, number string) (Card, error) {
	return getCardBySetAndNumber(set, number)
}

func GetCardByName(name string) (Card, error) {
	return getCardByName(name)
}

func GetCardByScryfallID(scryfallID string) (Card, error) {
	return getCardByScryfallID(scryfallID)
}
