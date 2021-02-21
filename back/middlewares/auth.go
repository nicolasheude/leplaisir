package middlewares

import (
	"dechild/controllers"
	"dechild/database"
	"dechild/softcrypto"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CkIdCorrect(identifiant string) error {
	user := database.Db.Def
	db, err := user.Admin.
		Query().
		All(database.Db.Ctx)
	identifiant = softcrypto.Decrypt(identifiant, hex.EncodeToString([]byte(os.Getenv("KEY"))))
	if err != nil || len(db) == 0{
		return fmt.Errorf("No matching data: %w", err)
	}
	DbAdmin := make([]controllers.AdminUser, len(db))
	for i := range db {
		DbAdmin[i] = controllers.AdminUser{
			Identifiant: softcrypto.Decrypt(db[i].Identifiant, hex.EncodeToString([]byte(os.Getenv("KEY")))),
		}
	}
	for i := range DbAdmin {
		if identifiant == DbAdmin[i].Identifiant {
			return nil
		}
	}
	return fmt.Errorf("No matching data: %w", err)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := c.Cookie("KEY")
		if id == "" || err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		if CkIdCorrect(id) == nil {
			c.Next()
			return
		} else {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
