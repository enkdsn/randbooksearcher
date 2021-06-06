package repos_test

import (
	"testing"

	"github.com/enkdsn/randbooksearcher/core/repos"
	"github.com/stretchr/testify/assert"
)

func Test_OrEbookRepo(t *testing.T) {
	assert := assert.New(t)
	oebRepo := repos.NewOrEBookRepo()
	books, err := oebRepo.Books()
	assert.NoError(err)
	assert.NotNil(books)
}
