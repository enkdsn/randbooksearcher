package fetch_test

import (
	"testing"

	"github.com/enkdsn/randbooksearcher/core/components/fetch"
	"github.com/stretchr/testify/assert"
)

func Test_OrEbook(t *testing.T) {
	assert := assert.New(t)
	orf := fetch.NewOrEbookFetcher()
	bs, err := orf.Fetch()
	assert.NoError(err)
	assert.NotNil(bs)
}
