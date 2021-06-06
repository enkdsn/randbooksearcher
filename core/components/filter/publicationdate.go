package filter

import (
	"errors"
	"time"

	"github.com/enkdsn/randbooksearcher/core/models"
)

type PublicationDateFilter struct {
	publicationDateFrom time.Time
	publicationDateTo   time.Time
}

func NewPublicationDateFilter(from, to time.Time) (Filter, error) {
	if from.After(to) {
		return nil, errors.New("from must before to")
	}
	return &PublicationDateFilter{
		publicationDateFrom: from,
		publicationDateTo:   to,
	}, nil
}

func (pdf *PublicationDateFilter) Filtrate() ([]models.Book, error) {
	return nil, nil
}
