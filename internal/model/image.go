package model

type Image struct {
	BlurHash  string `json:"blurhash"`
	URL       string `json:"url"`
	IsPreview bool   `json:"is_preview,omitempty"`
}
