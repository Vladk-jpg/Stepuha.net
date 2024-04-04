package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := r.Run(":0808")
	if err != nil {
		fmt.Println("Ooops...")
		return
	}
}
