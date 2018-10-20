package tag

func GetTags() ([]string, error) {
	return getTags()
}

func GetTagsForListID(listID int) ([]string, error) {
	return getTagsForListID(listID)
}

func GetTagsForEntryID(entryID int) ([]string, error) {
	return getTagsForEntryID(entryID)
}
