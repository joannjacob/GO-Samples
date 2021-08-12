package usecase

import (
	"github.com/bxcodec/go-clean-arch/article"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*article.Article, string, error)
	GetByID(id int64) (*article.Article, error)
	Update(ar *article.Article) (*article.Article, error)
	GetByTitle(title string) (*article.Article, error)
	Store(*article.Article) (*article.Article, error)
	Delete(id int64) (bool, error)
}
