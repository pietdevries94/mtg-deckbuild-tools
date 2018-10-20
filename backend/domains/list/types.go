package list

type List struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	IncludedTags []string `json:"included_tags"`
}

type AddListPayload struct {
	Name string `json:"name"`
}
