package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/enkdsn/randbooksearcher/core/components/fetch"
)

func main() {
	refreshOrEbookCSV()
	refreshOreillyBookCSV()
}

func refreshOrEbookCSV() {
	file, err := os.OpenFile("../../../resource/or_ebook.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	writer := csv.NewWriter(file)

	oef := fetch.NewOrEbookFetcher()
	bs, err := oef.Fetch()
	for _, b := range bs {
		writer.Write([]string{b.URL, b.Name, strconv.Itoa(b.Price), b.ISBN})
	}
	writer.Flush()

	defer file.Close()
}

func refreshOreillyBookCSV() {
	file, err := os.OpenFile("../../../resource/or_book.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	writer := csv.NewWriter(file)

	obf := fetch.NewOreillyBookFetcher()
	bs, err := obf.Fetch()
	for _, b := range bs {
		writer.Write([]string{b.URL, b.Name, strconv.Itoa(b.Price), b.ISBN})
	}
	writer.Flush()

	defer file.Close()
}
