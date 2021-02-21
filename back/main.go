package main

import (
	"context"
	"dechild/controllers"
	"dechild/database"
	"dechild/ent"
	"dechild/router"
	"dechild/server"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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


func CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	client := database.NewDatabase()
	s := server.NewServer()
	defer client.Def.Close()
	controllers.InitAdminManager()
	if IfStockInit(database.Db.Ctx, database.Db.Def) == false {
		controllers.InitStockManager(database.Db.Ctx, database.Db.Def)
	}
	s.Def.Use(CORS())
	router.ApplyRoutes(s.Def)
	s.Def.Run()
}
