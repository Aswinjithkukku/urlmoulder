package controllers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/aswinjithkukku/url-moulder/initializer"
	"github.com/aswinjithkukku/url-moulder/models"
	"github.com/aswinjithkukku/url-moulder/utils"
	"github.com/gin-gonic/gin"
)

// Home Url
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Get all possible urls.
func GetAllUrls(c *gin.Context) {
	var urls []models.Urls

	result := initializer.DB.Find(&urls)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Sorry Unable to fetch the data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"urls":   urls,
	})
}

// Add url.
func AddUrl(c *gin.Context) {
	body := c.PostForm("value")
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	urlRegex := regexp.MustCompile(`^(https?://|localhost:\d{4}//).*$`)
	if !urlRegex.MatchString(body) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": false,
			"error":  "Invalid Url",
		})
		return
	}

	var urls []models.Urls

	result := initializer.DB.Find(&urls)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Something went wrong. Please try again",
		})
		return
	}

	isUnique := false
	slug := ""

	for !isUnique {
		slugUrl, isExist := utils.GenerateRandomString(25, urls)

		if !isExist {
			isUnique = true
			slug = slugUrl
			break
		}
	}

	url := models.Urls{
		Url:        body,
		SlugUrl:    slug,
		ExpireDate: time.Now().Add(7 * 24 * time.Hour),
		IsExpired:  false,
	}

	result = initializer.DB.Create(&url)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Cannot generate url.Please try again!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"url":    url,
		"info":   "Your generated Url will expire within 7 Days",
	})
}

// Give data to user and Redirect.
func GiveRedirectionOutput(c *gin.Context) {
	slug := c.Param("slug")

	var url models.Urls
	fmt.Println(os.Getenv("URL_DOM") + slug)
	result := initializer.DB.First(&url, "slug_url = ? AND is_expired = ?", slug, false)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Could not find the Url",
		})
		return
	}

	if url.ExpireDate.Before(time.Now()) {
		url.IsExpired = true
		initializer.DB.Save(&url)
		c.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"error":  "Sorry!! the link you looking for is Expired. Please Generate another url.",
		})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, url.Url)
}
