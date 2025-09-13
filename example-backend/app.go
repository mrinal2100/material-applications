package main

import (
    "os"
    "server/common"
    "server/router"

    "github.com/gin-contrib/cors"
)

func main() {
    port := common.FallbackString(os.Getenv("PORT"), "8080")
    r := router.Router()

    // Add CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5000"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    r.Run(":" + port)
}
