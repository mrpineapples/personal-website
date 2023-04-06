package routes

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

func getArticle(c *gin.Context) {
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
			fmt.Println("No rows were returned!", slug)
		}
		panic(err)
	}

	var articleBuf bytes.Buffer
	if err := utils.Markdown.Convert([]byte(article.Markdown), &articleBuf); err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "article", gin.H{
		"pageTitle": "Michael's Site | " + article.Title,
		"Article":   article,
		"Content":   template.HTML(articleBuf.String()),
	})
}

func getArticles(c *gin.Context) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		ORDER BY id ASC;
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
		"pageTitle": "Michael's Site | Articles",
		"Articles":  articles,
	})
}

// newArticle renders the new article view
func newArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "new-article", gin.H{
		"pageTitle": "Michael's Site Admin | Create An Article",
	})
}

// createArticle creates a new article in our database
func createArticle(c *gin.Context) {
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

// editArticle renders the edit article view
func editArticle(c *gin.Context) {
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
		"Article": article,
	})
}

func updateArticle(c *gin.Context) {
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

func deleteArticle(c *gin.Context) {
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
