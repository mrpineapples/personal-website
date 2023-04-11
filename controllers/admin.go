package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/models"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin", gin.H{})
}

func GetAdminArticles(c *gin.Context) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		ORDER BY created_at DESC;
	`
	rows, err := models.Pool.Query(models.DBContext, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.Slug, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "articles", gin.H{
		"PageTitle": "Michael's Site Admin | Articles",
		"Articles":  articles,
		"IsAdmin":   true,
	})
}
