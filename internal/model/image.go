package model

type Image struct {
	BlurHash  string `json:"blurhash,omitempty"`
	URL       string `json:"url,omitempty"`
	IsPreview bool   `json:"is_preview,omitempty"`
}
