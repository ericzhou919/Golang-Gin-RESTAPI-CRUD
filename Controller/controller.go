package controller

import (
	models "golang-gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.String(http.StatusOK, models.GetAll())
}

func GetUser(c *gin.Context) {
	id := (c.Param("id"))
	c.String(http.StatusOK, models.GetOne(id))
}
func Insert(c *gin.Context) {
	name := c.PostForm("name")
	models.Insert(name)
	c.JSON(200, gin.H{
		"message": "Success insert",
	})
}
func Update(c *gin.Context) {
	name := c.PostForm("name")
	id := (c.Param("id"))
	if models.Update(id, name) {
		c.JSON(200, gin.H{
			"message": "Success update",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Not found",
		})
	}
}
func Delete(c *gin.Context) {
	id := (c.Param("id"))
	if models.Delete(id) {
		c.JSON(200, gin.H{
			"message": "Success delete",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Not found",
		})
	}
}
