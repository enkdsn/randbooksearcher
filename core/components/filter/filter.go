package filter

import "github.com/enkdsn/randbooksearcher/core/models"

type Filter interface {
	Filtrate([]*models.Book) ([]*models.Book, error)
}
