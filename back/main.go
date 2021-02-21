package main

import (
	"context"
	"dechild/controllers"
	"dechild/database"
	"dechild/ent"
	"dechild/router"
	"dechild/server"
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
	s.Def.Run()
}