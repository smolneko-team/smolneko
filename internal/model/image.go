package model

type Image struct {
	BlurHash  string `json:"blurhash"`
	URL       string `json:"urls"`
	IsPreview bool   `json:"is_preview,omitempty"`
}
