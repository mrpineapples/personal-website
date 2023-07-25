package models

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mrpineapples/personal-website/utils"
)

type Article struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Markdown    string    `json:"markdown"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (a *Article) ToHTML() (template.HTML, error) {
	articleBuf := bytes.Buffer{}
	if err := utils.Markdown.Convert([]byte(a.Markdown), &articleBuf); err != nil {
		return "", nil
	}

	return template.HTML(articleBuf.String()), nil
}

func GetArticleBySlug(slug string) (Article, error) {
	sqlStatement := `
		SELECT title, description, markdown, created_at, updated_at
		FROM articles
		WHERE slug = $1;
	`

	article := Article{}
	row := Pool.QueryRow(DBContext, sqlStatement, slug)
	err := row.Scan(&article.Title, &article.Description, &article.Markdown, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows were returned for the following slug:", slug)
		}
		return Article{}, err
	}

	return article, nil
}
