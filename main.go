package main

import "github.com/Tiratom/gin-study/config"

func main() {
	r := config.GetRouter()
	r.Run(":8080")
}
