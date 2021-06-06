package fetch

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/enkdsn/randbooksearcher/core/models"
)

const (
	oreillyEbook = "https://www.oreilly.co.jp/ebook/#all_titles"
)

type OrEbookFetcher struct {
	url string
}

func NewOrEbookFetcher() *OrEbookFetcher {
	return &OrEbookFetcher{
		url: oreillyEbook,
	}
}

func (oef OrEbookFetcher) Fetch() ([]models.Book, error) {
	url := oef.url
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var bs []models.Book

	doc.Find("#bookTable").Children().Children().Each(func(_ int, s *goquery.Selection) {
		t := s.Find(".title")
		html := t.Find("a")
		href, _ := html.Attr("href")
		isbn := s.Find(".isbn")

		pi, err := strconv.Atoi(strings.Replace(s.Find(".price").Text(), ",", "", -1))
		if err != nil {
			pi = 0
		}

		b := models.Book{
			URL:   strings.Replace(href, "./..", "https://www.oreilly.co.jp", -1),
			Name:  t.Text(),
			Price: pi,
			ISBN:  isbn.Text(),
		}

		if b.Name == "" {
			return
		}
		bs = append(bs, b)

	})
	return bs, nil
}
