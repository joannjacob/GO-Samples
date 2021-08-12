package article_usecase_test

import (
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	models "github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository/mocks"
	ucase "github.com/bxcodec/go-clean-arch/article/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, &mockArticle)
	mockArticleRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockArticle.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockArticleRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}
