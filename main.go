package main

import (
	"douyin_backend/config"
	"douyin_backend/router"
	"fmt"
)

func main() {
	r := router.Init()
	err := r.Run(fmt.Sprintf(":%d", config.Global.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
