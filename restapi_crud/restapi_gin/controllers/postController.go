package controllers

import (
	"yashprakash13/Go-Basics/initializers"
	"yashprakash13/Go-Basics/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	// get data off request reqBody
	var reqBody struct {
		Title string
		Body  string
	}
	c.Bind(&reqBody)

	//Create a post and return it
	post := models.Post{Title: reqBody.Title, Body: reqBody.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)

	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//get id of post
	postId := c.Param("id")

	//get the specific post
	var post models.Post
	initializers.DB.First(&post, postId)

	//respond with the posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//get id of the post
	postId := c.Param("id")

	//get data from request body
	var reqBody struct {
		Title string
		Body  string
	}
	c.Bind(&reqBody)

	//get the specific post
	var post models.Post
	initializers.DB.First(&post, postId)

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: reqBody.Title, Body: reqBody.Body})

	//respond with the updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get id of the post
	postId := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, postId)

	// respond
	c.Status(200)
}
