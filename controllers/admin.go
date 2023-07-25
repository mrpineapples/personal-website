package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mrpineapples/personal-website/models"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin", gin.H{
		"PageTitle": "Admin",
	})
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
		"PageTitle":    "Admin | Posts",
		"FaviconEmoji": "🛠",
		"Articles":     articles,
		"IsAdmin":      true,
	})
}

// EditArticle renders the edit article view
func EditArticle(c *gin.Context) {
	sqlStatement := `
		SELECT id, title, description, markdown, created_at, updated_at
		FROM articles
		WHERE slug = $1;
	`
	slug := c.Param("slug")
	row := models.Pool.QueryRow(models.DBContext, sqlStatement, slug)

	var article models.Article
	err := row.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows were returned!", slug)
		}
		panic(err)
	}
	article.Slug = slug

	c.HTML(http.StatusOK, "edit-article", gin.H{
		"PageTitle":    "Editing | " + article.Title,
		"FaviconEmoji": "🛠",
		"Article":      article,
	})
}
