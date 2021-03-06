package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/enkdsn/randbooksearcher/core/components/randomize"
	"github.com/enkdsn/randbooksearcher/core/repos"
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
	Short: "return books these sum of prices by maxprice",
	Long:  `-simple return books that sum of price by 11000 yen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mp, err := cmd.Flags().GetInt("maxprice")
		if err != nil {
			return err
		}

		oebRepo := repos.NewOrEBookRepo()
		books, err := oebRepo.Books()
		if err != nil {
			return err
		}
		if books == nil {
			return errors.New("no data")
		}

		books = randomize.RandomizeWithMaxPrice(mp, books)

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
