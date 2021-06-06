package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/enkdsn/randbooksearcher/core/components/fetch"
	"github.com/enkdsn/randbooksearcher/core/components/randomize"
	"github.com/spf13/cobra"
)

type SimpleCmdJSONResponse struct {
	SimpleCmdJSONResponse []SimpleCmdResponse `json:"SimpleCmdJSONResponse"`
	TotalPrice            int                 `json:"totalPrice"`
}

type SimpleCmdResponse struct {
	URL       string `json:"url"`
	BookTitle string `json:"bookTitle"`
	BookPrice int    `json:"bookPrice"`
	ISBN      string `json:"isbn,omitempty"`
}

var SimpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "3 books from Oreilly Ebook",
	Long:  `-simple return books that sum of price by 11000 yen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := cmd.Flags().GetInt("integer")
		if err != nil {
			return err
		}
		b, err := cmd.Flags().GetBool("boolean")
		if err != nil {
			return err
		}
		s, err := cmd.Flags().GetString("string")
		if err != nil {
			return err
		}
		oef := fetch.NewOrEbookFetcher()
		books, err := oef.Fetch()
		if err != nil {
			return err
		}

		books = randomize.Randomize(books)

		var response SimpleCmdJSONResponse
		for _, b := range books {
			r := SimpleCmdResponse{
				URL:       b.URL,
				BookTitle: b.Name,
				BookPrice: b.Price,
				ISBN:      b.ISBN,
			}
			response.SimpleCmdJSONResponse = append(response.SimpleCmdJSONResponse, r)
			response.TotalPrice += r.BookPrice
		}
		bytes, _ := json.Marshal(response)
		fmt.Println(string(bytes))
		return nil
	},
}
