package controllers

import (
	"demo-api/initializers"
	"demo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// get req body

	var body struct {
		Title       string
		Description string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the req",
		})
		return
	}
	// create post
	user, _ := c.Get("user")
	newPost := models.Post{Title: body.Title, Description: body.Description, UserId: user.(models.User).ID}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	// return
	c.JSON(http.StatusCreated, gin.H{})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Select("id", "title", "user_id", "description").Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// get id
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {

	// get req body

	var body struct {
		Title       string
		Description string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the req",
		})
		return
	}

	// get id
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// update post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Description: body.Description})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// get id
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(http.StatusOK)
}
