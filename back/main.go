package main

import (
	"dechild/controllers"
	"dechild/database"
	"dechild/server"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/gin-contrib/cors"
	"time"
)

func GetForm(c *gin.Context) {
	data := controllers.FormData{}
	err := c.ShouldBind(&data)
	status := controllers.CheckDoesNotExist(data, database.Db.Ctx, database.Db.Def)
	if status == true {
		controllers.CreateFrom(data, database.Db.Ctx, database.Db.Def)
	}
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.POST("/hello", GetForm)
}

func main() {
	client := database.NewDatabase()
	s := server.NewServer()
	defer client.Def.Close()

	s.Def.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))

	ApplyRoutes(s.Def)
	s.Def.Run()
}