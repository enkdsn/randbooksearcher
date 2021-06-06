package filter

import (
	"errors"
	"time"

	"github.com/enkdsn/randbooksearcher/core/models"
)

type PublicationDateFilter struct {
	From time.Time
	To   time.Time
}

func NewPublicationDateFilter(from, to time.Time) (Filter, error) {
	if from.After(to) {
		return nil, errors.New("from must before to")
	}
	return &PublicationDateFilter{
		From: from,
		To:   to,
	}, nil
}

func (pdf *PublicationDateFilter) Filtrate(bs []*models.Book) ([]*models.Book, error) {
	var filtered []*models.Book
	for _, b := range bs {
		if b.PublicationDate.After(pdf.To) || b.PublicationDate.Before(pdf.From) {
			continue
		}
		filtered = append(filtered, b)
	}
	return filtered, nil
}
