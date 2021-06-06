package fetch

import "github.com/enkdsn/randbooksearcher/core/models"

type Fetcher interface {
	Fetch() ([]models.Book, error)
}
