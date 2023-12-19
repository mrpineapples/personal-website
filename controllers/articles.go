package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/mrpineapples/personal-website/components"
	"github.com/mrpineapples/personal-website/models"
)

var articleInput struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Markdown    string `form:"markdown"`
}

func GetArticle(c *gin.Context) {
	article, err := models.GetArticleBySlug(c.Param("slug"))
	if err != nil {
		c.HTML(http.StatusNotFound, "", components.NotFound())
		return
	}

	// html, err := article.ToHTML()
	// if err != nil {
	// 	panic(err)
	// }

	if article.Description == "" {
		article.Description = "An article written by Michael"
	}

	c.HTML(http.StatusOK, "", components.Article(article))

	// c.HTML(http.StatusOK, "article", gin.H{
	// 	"PageTitle":       article.Title,
	// 	"PageDescription": pageDescription,
	// 	"Article":         article,
	// 	"Content":         html,
	// })
}

func GetArticles(c *gin.Context) {
	a, err := models.GetArticles()
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "", components.Articles(a, false))

	// c.HTML(http.StatusOK, "articles", gin.H{
	// 	"PageTitle":       "Michael Miranda | Blog",
	// 	"PageDescription": "All of Michael's posts",
	// 	"Articles":        articles,
	// })
}

// NewArticle renders the new article view
func NewArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.NewArticle())
	// c.HTML(http.StatusOK, "new-article", gin.H{
	// 	"PageTitle":    "Admin | Create a Post",
	// 	"FaviconEmoji": "🛠",
	// })
}

// CreateArticle creates a new article in our database
func CreateArticle(c *gin.Context) {
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
	err := models.CreateArticle(article)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusSeeOther, "/admin/blog")
}

func UpdateArticle(c *gin.Context) {
	if err := c.ShouldBind(&articleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := models.GetArticleByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	article.Title = articleInput.Title
	article.Description = articleInput.Description
	article.Markdown = articleInput.Markdown
	article.Slug = slug.Make(articleInput.Title)
	article.UpdatedAt = time.Now()

	err = models.UpdateArticle(article)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusSeeOther, "/blog/"+article.Slug)
}

func DeleteArticle(c *gin.Context) {
	err := models.DeleteArticle(c.Param("id"))
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusSeeOther, "/admin/blog")
}
