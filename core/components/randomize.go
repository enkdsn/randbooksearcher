package components

import (
	"math/rand"
	"time"

	"github.com/enkdsn/randbooksearcher/core/models"
)

func Randomize(bs []models.Book) []models.Book {
	var bss []models.Book
	sum := 0
	// 11000円を超えないように探索
	for range bs {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(len(bs))
		sum += bs[idx].Price

		// fmt.Printf("%+v", bs[idx])
		if sum > 11000 {
			bs = remove(bs, idx)
			continue
		}
		bss = append(bss, bs[idx])
		bs = remove(bs, idx)
	}
	return bss
}

func remove(slice []models.Book, s int) []models.Book {
	return append(slice[:s], slice[s+1:]...)
}
