package list

import (
	"github.com/pietdevries94/mtg-deckbuild-tools/backend/domains/tag"
)

func GetLists() ([]List, error) {
	lists, err := getLists()
	if err != nil {
		return nil, err
	}

	for i, list := range lists {
		tags, err := tag.GetTagsForListID(list.ID)
		if err != nil {
			return nil, err
		}
		lists[i].IncludedTags = tags
	}

	return lists, nil
}

func AddList(payload AddListPayload) error {
	return addList(payload.Name)
}
