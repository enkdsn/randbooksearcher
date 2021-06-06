package fetch_test

import (
	"testing"

	"github.com/enkdsn/randbooksearcher/core/components/fetch"
	"github.com/stretchr/testify/assert"
)

func Test_OreillyBook(t *testing.T) {
	assert := assert.New(t)
	orf := fetch.NewOreillyBookFetcher()
	bs, err := orf.Fetch()
	assert.NoError(err)
	assert.NotNil(bs)
}
