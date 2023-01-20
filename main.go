package main

import (
	"fmt"
	models "golang-gin/model"
	"golang-gin/router"
)

var index = 4

func init() {
	models.Insert("A")
	models.Insert("B")
	models.Insert("C")
	models.Insert("D")
}

func main() {
	fmt.Printf("start")
	router := router.InitRouter()
	router.Run(":8000")
}
