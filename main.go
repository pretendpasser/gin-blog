package main

import (
	"blog/routes"
	"blog/model"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}