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
		"message": "insert success!",
	})

}
func Delete(c *gin.Context) {
	id := (c.Param("id"))
	models.Delete(id)
	c.JSON(200, gin.H{
		"message": "success delete",
	})
}
