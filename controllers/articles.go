package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
	"github.com/mrpineapples/personal-website/models"
	"github.com/mrpineapples/personal-website/utils"
)

func GetArticle(c *gin.Context) {
	sqlStatement := `
		SELECT title, description, markdown, created_at, updated_at
		FROM articles
		WHERE slug = $1;
	`
	slug := c.Param("slug")
	row := models.Pool.QueryRow(models.DBContext, sqlStatement, slug)

	var article models.Article
	err := row.Scan(&article.Title, &article.Description, &article.Markdown, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows were returned for the following slug:", slug)
		}
		utils.RenderNotFound(c)
	}

	var articleBuf bytes.Buffer
	if err := utils.Markdown.Convert([]byte(article.Markdown), &articleBuf); err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "article", gin.H{
		"PageTitle": article.Title,
		"Article":   article,
		"Content":   template.HTML(articleBuf.String()),
	})
}

func GetArticles(c *gin.Context) {
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
		"PageTitle": "Michael Miranda | Articles",
		"Articles":  articles,
	})
}

// NewArticle renders the new article view
func NewArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "new-article", gin.H{
		"PageTitle": "Admin | Create An Article",
	})
}

// CreateArticle creates a new article in our database
func CreateArticle(c *gin.Context) {
	var articleInput struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Markdown    string `form:"markdown"`
	}

	if err := c.ShouldBind(&articleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:       articleInput.Title,
		Description: articleInput.Description,
		Markdown:    articleInput.Markdown,
		Slug:        slug.Make(articleInput.Title),
	}

	sqlStatement := `
		INSERT INTO articles (title, description, markdown, slug)
		VALUES ($1, $2, $3, $4)
		RETURNING slug;
	`
	var slug string
	err := models.Pool.QueryRow(models.DBContext, sqlStatement, article.Title, article.Description, article.Markdown, article.Slug).Scan(&slug)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusSeeOther, "/articles/"+article.Slug)
}

func UpdateArticle(c *gin.Context) {
	var articleInput struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Markdown    string `form:"markdown"`
	}

	if err := c.ShouldBind(&articleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:       articleInput.Title,
		Description: articleInput.Description,
		Markdown:    articleInput.Markdown,
		Slug:        slug.Make(articleInput.Title),
		UpdatedAt:   time.Now(),
	}

	sqlStatement := `
		UPDATE articles
		SET slug = $2, title = $3, description = $4, markdown = $5, updated_at = $6
		WHERE id = $1;
	`
	cmdTag, err := models.Pool.Exec(models.DBContext, sqlStatement, c.Param("id"), article.Slug, article.Title, article.Description, article.Markdown, article.UpdatedAt)
	if err != nil {
		panic(err)
	} else if cmdTag.RowsAffected() != 1 {
		panic(errors.New("No row found to update"))
	}

	c.Redirect(http.StatusSeeOther, "/articles/"+article.Slug)
}

func DeleteArticle(c *gin.Context) {
	sqlStatement := `
		DELETE FROM articles
		WHERE id = $1;
	`

	cmdTag, err := models.Pool.Exec(models.DBContext, sqlStatement, c.Param("id"))
	if err != nil {
		panic(err)
	} else if cmdTag.RowsAffected() != 1 {
		panic(errors.New("No row found to delete"))
	}
	c.Redirect(http.StatusSeeOther, "/articles")
}
