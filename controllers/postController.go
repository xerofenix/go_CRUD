package controllers

import (
	"crud_json/initializers"
	"crud_json/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// function to create the post
func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// function to show all the posts
func PostIndex(c *gin.Context) {

	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// func to show a particualr post by id
func PostShow(c *gin.Context) {

	id := c.Param("id")
	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

// func to update the post
func PostsUpdate(c *gin.Context) {
	// get the id of URL
	id := c.Param("id")

	//get the data of request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//updating with it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// function to delete the post
func PostDelete(c *gin.Context) {
	//get the id of the url
	id := c.Param("id")

	//delte the post
	err := initializers.DB.Delete(&models.Post{}, id)
	//respond back
	if err != nil {
		fmt.Println(err)
	} else {
		c.Status(200)
	}
}
