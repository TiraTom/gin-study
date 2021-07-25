package main

import (
	"gin-study/config"
)

func main() {
	r := config.GetRouter()
	r.Run(":8080")
}
