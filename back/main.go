package main

import (
	"context"
	"dechild/controllers"
	"dechild/database"
	"dechild/ent"
	"dechild/router"
	"dechild/server"
	"time"

	"github.com/gin-contrib/cors"
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

func main() {
	client := database.NewDatabase()
	s := server.NewServer()
	defer client.Def.Close()
	controllers.InitAdminManager()
	if IfStockInit(database.Db.Ctx, database.Db.Def) == false {
		controllers.InitStockManager(database.Db.Ctx, database.Db.Def)
	}
	router.ApplyRoutes(s.Def)
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
	s.Def.Run()
}
