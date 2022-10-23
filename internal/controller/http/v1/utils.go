package v1

import "github.com/smolneko-team/smolneko/internal/model"

type imagesResponse struct {
	Images model.Image `json:"data"`
}
