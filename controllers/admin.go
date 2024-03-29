package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mrpineapples/personal-website/components"
	"github.com/mrpineapples/personal-website/database"
	"github.com/mrpineapples/personal-website/models"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.AdminHome())
}

func GetAdminArticles(c *gin.Context) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		ORDER BY created_at DESC;
	`
	rows, err := database.Pool.Query(database.DBContext, sqlStatement)
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

	c.HTML(http.StatusOK, "", components.Articles(articles, true))
}

// EditArticle renders the edit article view
func EditArticle(c *gin.Context) {
	sqlStatement := `
		SELECT id, title, description, markdown, created_at, updated_at
		FROM articles
		WHERE slug = $1;
	`
	slug := c.Param("slug")
	row := database.Pool.QueryRow(database.DBContext, sqlStatement, slug)

	var article models.Article
	err := row.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows were returned!", slug)
		}
		panic(err)
	}
	article.Slug = slug

	c.HTML(http.StatusOK, "", components.EditArticle(article))
}
