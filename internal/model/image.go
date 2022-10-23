package model

type Image struct {
	ID    string   `json:"id"`
	Count int      `json:"count"`
	URL   []string `json:"urls"`
}
