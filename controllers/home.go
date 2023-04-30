package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/models"
)

func Home(c *gin.Context) {
	sqlStatement := `
		SELECT title, slug, created_at
		FROM articles
		ORDER BY created_at DESC
		LIMIT 3;
	`

	rows, err := models.Pool.Query(models.DBContext, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.Title, &article.Slug, &article.CreatedAt)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}

	c.HTML(http.StatusOK, "home", gin.H{
		"Articles": articles,
	})
}
