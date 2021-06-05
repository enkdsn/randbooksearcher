package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	oreilly      = "https://www.oreilly.co.jp/catalog/"
	oreillyEbook = "https://www.oreilly.co.jp/ebook/#all_titles"
)

type bookinfo struct {
	url   string
	name  string
	price int
}

type SearchBookResponse struct {
	URL       string `json:"url"`
	BookTitle string `json:"bookTitle"`
	BookPrice int    `json:"bookPrice"`
}

func main() {
	r, err := fetch()
	if err != nil {
		os.Exit(1)
	}

	result := randomize(r)

	fmt.Printf("%+v\n", result)
	fmt.Printf("total price: %d", total(result))
}

func total(bs []bookinfo) int {
	sum := 0
	for _, b := range bs {
		sum += b.price
	}
	return sum
}

func fetch() ([]bookinfo, error) {
	url := oreillyEbook
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var bookinfos []bookinfo

	doc.Find("#bookTable").Children().Children().Each(func(_ int, s *goquery.Selection) {
		t := s.Find(".title")
		html := t.Find("a")
		href, _ := html.Attr("href")

		pi, err := strconv.Atoi(strings.Replace(s.Find(".price").Text(), ",", "", -1))
		if err != nil {
			pi = 0
		}

		bi := bookinfo{
			url:   strings.Replace(href, "./..", "https://www.oreilly.co.jp", -1),
			name:  t.Text(),
			price: pi,
		}

		if bi.name == "" {
			return
		}
		bookinfos = append(bookinfos, bi)

	})
	return bookinfos, nil
}

func randomize(bs []bookinfo) []bookinfo {
	var bookinfos []bookinfo
	sum := 0
	// 11000円を超えないように探索
	for range bs {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(len(bs))
		sum += bs[idx].price

		// fmt.Printf("%+v", bs[idx])
		if sum > 11000 {
			bs = remove(bs, idx)
			continue
		}
		bookinfos = append(bookinfos, bs[idx])
		bs = remove(bs, idx)
	}
	return bookinfos
}

func remove(slice []bookinfo, s int) []bookinfo {
	return append(slice[:s], slice[s+1:]...)
}
