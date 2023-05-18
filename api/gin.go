package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/goMicroservice/database"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to building restapi using gin and gorm",
	})
	return
}

func PostArticleHaldler(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
	return
}

func GetArticleHandler(c *gin.Context) {
	id := c.Param("id")
	article, err := database.ReadArticle(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
	return
}

func GetArticlesHandler(c *gin.Context) {
	articles, err := database.ReadArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func PutArticleHandler(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.UpdateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
	return
}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article delete successfully.",
	})
	return
}

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/api/v1/articles/:id", GetArticleHandler)
	r.GET("/api/v1/articles", GetArticlesHandler)
	r.POST("/api/v1/articles", PostArticleHaldler)
	r.PUT("/api/v1/articles/:id", PutArticleHandler)
	r.DELETE("/api/v1/articles/:id", DeleteArticleHandler)
	return r
}
