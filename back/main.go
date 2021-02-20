package main

import (
	"context"
	"dechild/controllers"
	"dechild/database"
	"dechild/ent"
	"dechild/server"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"net/http"
)

func IfStockInit(ctx context.Context, client *ent.Client) bool {
	db, _ := client.StockManager.
		Query().
		All(ctx)
	if len(db) == 0 {
		return false
	}
	return true
}

func GetForm(c *gin.Context) {
	data := controllers.FormData{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	_, err = controllers.CreateFrom(data, database.Db.Ctx, database.Db.Def)
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
	controllers.InitAdminManager()
	if IfStockInit(database.Db.Ctx, database.Db.Def) == false {
		controllers.InitStockManager(database.Db.Ctx, database.Db.Def)
	}
	ApplyRoutes(s.Def)
	s.Def.Run()
}