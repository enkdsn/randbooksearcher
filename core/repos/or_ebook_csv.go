package repos

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/enkdsn/randbooksearcher/core/models"
)

type OrEBookRepo struct {
}

func NewOrEBookRepo() *OrEBookRepo {
	return &OrEBookRepo{}
}

type OrEBookStr struct {
	URL             string
	Name            string
	Price           string
	ISBN            string
	PublicationDate string
}

const (
	URLIdx = iota
	NameIdx
	PriceIdx
	ISBNIdx
	PublicationDateIdx
)

func (oer *OrEBookRepo) Books() ([]*models.Book, error) {
	var bs []*models.Book
	file, err := os.Open("../resource/or_ebook.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		price, err := strconv.Atoi(line[PriceIdx])
		if err != nil {
			return nil, err
		}

		b := &models.Book{
			URL:   line[URLIdx],
			Name:  line[NameIdx],
			Price: price,
		}
		bs = append(bs, b)

	}
	return bs, nil
}
