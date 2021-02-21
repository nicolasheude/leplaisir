package controllers

import (
	"context"
	"dechild/database"
	"dechild/ent"
	"dechild/softcrypto"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type AdminUser struct {
	Identifiant string `json:"identifiant"`
	MotDePasse string `json:"mot_de_passe"`
}

func IfAdminInit(ctx context.Context, client *ent.Client) bool {
	db, _ := client.Admin.
		Query().
		All(ctx)
	if len(db) == 0 {
		return false
	}
	return true
}

func InitAdmin() {
	admin := database.Db.Def
	identifiant := softcrypto.Encrypt(os.Getenv("ID"), hex.EncodeToString([]byte(os.Getenv("KEY"))))
	password := softcrypto.Encrypt(os.Getenv("PW"), hex.EncodeToString([]byte(os.Getenv("KEY"))))
	admin.Admin.
		Create().
		SetIdentifiant(identifiant).
		SetMotDePasse(password).
		Save(database.Db.Ctx)
}

func InitAdminManager() {
	if IfAdminInit(database.Db.Ctx, database.Db.Def) == false {
		InitAdmin()
	}
}

func ifIdCorrect(identifiant, password string) error {
	user := database.Db.Def
	db, err := user.Admin.
		Query().
		All(database.Db.Ctx)
	if err != nil || len(db) == 0{
		return fmt.Errorf("No matching data: %w", err)
	}
	DbAdmin := make([]AdminUser, len(db))
	for i := range db {
		DbAdmin[i] = AdminUser{
			Identifiant: softcrypto.Decrypt(db[i].Identifiant, hex.EncodeToString([]byte(os.Getenv("KEY")))),
			MotDePasse:  softcrypto.Decrypt(db[i].MotDePasse, hex.EncodeToString([]byte(os.Getenv("KEY")))),
		}
	}
	for i := range DbAdmin {
		if identifiant == DbAdmin[i].Identifiant && password == DbAdmin[i].MotDePasse {
			return nil
		}
	}
	return fmt.Errorf("No matching data: %w", err)
}

func LoginSession(c *gin.Context) {
	identifiant := c.PostForm("identifiant")
	password := c.PostForm("password")
	if identifiant == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if ifIdCorrect(identifiant, password) != nil {
			c.String(http.StatusBadRequest, "No user corresponds to the transmitted identifier, please try again.")
			return
		}
		c.SetCookie("KEY", softcrypto.Encrypt(identifiant, hex.EncodeToString([]byte(os.Getenv("KEY")))), 3600, "/", "127.0.0.1", false, true)
		c.String(http.StatusOK, "User successfully logged in !")
		fmt.Println("HEY JE SUIS LOGIN")
	}
	return
}
