package main

import (
	"dechild/controllers"
	"dechild/database"
	"dechild/server"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"net/http"
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
	ApplyRoutes(s.Def)
	s.Def.Run()
}