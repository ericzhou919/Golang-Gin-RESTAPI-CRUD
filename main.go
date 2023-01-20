package main

import (
	"fmt"
	models "golang-gin/model"
	"golang-gin/router"
)

func init() {
	models.Insert("A")
	models.Insert("B")
	models.Insert("C")
}

func main() {
	fmt.Printf("start")
	router := router.InitRouter()
	router.Run(":8000")
}
