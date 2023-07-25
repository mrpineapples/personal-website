package models

import (
	"bytes"
	"errors"
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

func GetArticleByID(ID string) (Article, error) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		WHERE id = $1;
	`

	article := Article{}
	row := Pool.QueryRow(DBContext, sqlStatement, ID)
	err := row.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.Slug, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(`No rows were returned for the following ID:`, ID)
		}
		return Article{}, err
	}

	return article, nil
}

func GetArticleBySlug(slug string) (Article, error) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		WHERE slug = $1;
	`

	article := Article{}
	row := Pool.QueryRow(DBContext, sqlStatement, slug)
	err := row.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.Slug, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows were returned for the following slug:", slug)
		}
		return Article{}, err
	}

	return article, nil
}

func GetArticles() ([]Article, error) {
	sqlStatement := `
		SELECT id, title, description, markdown, slug, created_at, updated_at
		FROM articles
		ORDER BY created_at DESC;
	`
	rows, err := Pool.Query(DBContext, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]Article, 0)
	for rows.Next() {
		article := Article{}
		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Markdown, &article.Slug, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func CreateArticle(article Article) error {
	sqlStatement := `
		INSERT INTO articles (title, description, markdown, slug)
		VALUES ($1, $2, $3, $4);
	`
	_, err := Pool.Exec(DBContext, sqlStatement, article.Title, article.Description, article.Markdown, article.Slug)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArticle(article Article) error {
	sqlStatement := `
		UPDATE articles
		SET slug = $2, title = $3, description = $4, markdown = $5, updated_at = $6
		WHERE id = $1;
	`
	cmdTag, err := Pool.Exec(DBContext, sqlStatement, article.ID, article.Slug, article.Title, article.Description, article.Markdown, article.UpdatedAt)
	if err != nil {
		return err
	} else if cmdTag.RowsAffected() != 1 {
		return errors.New("No row found to update")
	}

	return nil
}

func DeleteArticle(ID string) error {
	sqlStatement := `
		DELETE FROM articles
		WHERE id = $1;
	`
	cmdTag, err := Pool.Exec(DBContext, sqlStatement, ID)
	if err != nil {
		return err
	} else if cmdTag.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}

	return nil
}
