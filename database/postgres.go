package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

func NewPostgreSQLClient() {

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=regmi dbname=myarticles password=sainamaina sslmode=disable")
	//testing database connection
	// err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article not created.")
	}
	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var article Article
	res := db.First(&article, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("Article not found.")
	}
	return &article, nil
}

func ReadArticles() ([]*Article, error) {
	var articles []*Article
	res := db.Find(&articles)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return articles, nil
}

func UpdateArticle(article *Article) (*Article, error) {
	var updateArticle Article
	result := db.Model(&updateArticle).Where(article.ID).Update(article)
	if result.RowsAffected == 0 {
		return &Article{}, errors.New("Article not updated.")
	}
	return &updateArticle, nil
}

func DeleteArticle(id string) error {
	var deleteArticle Article
	result := db.Where(id).Delete(&deleteArticle)
	if result.RowsAffected == 0 {
		return errors.New("article data not deleted.")
	}
	return nil
}
