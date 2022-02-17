package article

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"helpstack/models"
)

// NewArticleGormRepo will create a new object of
func NewArticleGormRepo(db *gorm.DB) ArticleRepository {
	return &ArticleGormRepo{conn: db}
}

// ArticleRepository specifies application article related database operations
type ArticleRepository interface {
	FetchMany(ctx context.Context, filter models.Filter) ([]Article, error)
	GetOneById(ctx context.Context, id string) (*Article, error)
	UpdateOne(ctx context.Context, article *Article) (int64, int64, error)
	DeleteOne(ctx context.Context, id string) (int64, error)
	CreateOne(ctx context.Context, article *Article) (*Article, error)
	GetArticlesFunction(ctx context.Context) ([]Article, error)
}

// ArticleGormRepo implements the item.ItemRepository interface
type ArticleGormRepo struct {
	conn *gorm.DB
}

func (u ArticleGormRepo) FetchMany(ctx context.Context, filter models.Filter) ([]Article, error) {
	var usrs []Article

	err := u.conn.Find(&usrs).Error
	if err!=nil{
		return nil, err
	}
	return usrs, err

}

func (u ArticleGormRepo) GetOneById(ctx context.Context, id string) (*Article, error) {
	article:= Article{}
	errs:=u.conn.First(&article).Error

	if errs!=nil{
		return nil, errs
	}
	return &article, nil
}

func (u ArticleGormRepo) UpdateOne(ctx context.Context, article *Article) (int64, int64, error) {
	panic("implement me")
}

func (u ArticleGormRepo) DeleteOne(ctx context.Context, id string) (int64, error) {
	panic("implement me")
}

func (u ArticleGormRepo) CreateOne(ctx context.Context, article *Article) (*Article, error) {

	err := u.conn.Create(article).Error
	if err!=nil {
		return nil, err
	}
	return article, err
}

func (u ArticleGormRepo) GetArticlesFunction(ctx context.Context) ([]Article, error) {

	var art []Article
	err := u.conn.Raw("select get_all_articles()").Scan(&art)
	fmt.Println("err", err)
	if err!=nil {
		return nil, nil
	}
	return art, nil
}

//func (u ArticleGormRepo) CountFunction(ctx context.Context) (int, error) {
//
//	var num int
//	val := u.conn.Raw("select get_numberof_articles()").Scan(&num)
//	fmt.Println("err", val)
//	if val==nil {
//		return val, errors.New("record not found")
//	}
//	return num, nil
//}



func (u ArticleGormRepo) CreateArticleFunction(ctx context.Context, article *Article) ([]Article, error) {

	var art []Article
	err := u.conn.Raw("select get_all_articles()").Scan(&art)
	if err!=nil {
		return nil, nil
	}
	return art, nil
}
