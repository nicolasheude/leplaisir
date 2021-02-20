package controllers

import (
	"context"
	"dechild/database"
	"dechild/ent"
)

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
	admin.Admin.
		Create().
		SetIdentifiant("nicolas.heude@tcm.fr").
		SetMotDePasse("1234").
		Save(database.Db.Ctx)
}

func InitAdminManager() {
	if IfAdminInit(database.Db.Ctx, database.Db.Def) == false {
		InitAdmin()
	}
}

/*
func ifIdCorrect(email, password string) int {
	lenDb := len(User.UserDB) - 1
	if lenDb < 0 {
		return 84
	}
	for i := 0; i <= lenDb; i++ {
		if User.UserDB[i].Email == email && User.UserDB[i].Password == password {
			return 0
		}
	}
	return 84
}

func singinSession(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if ifIdCorrect(email, password) == 84 {
			c.String(http.StatusBadRequest, "No user corresponds to the transmitted identifier, please try again.")
			return
		}
		emailCry := softcrypto.Encrypt(email, hex.EncodeToString([]byte(os.Getenv("KEY"))))
		c.SetCookie("email", emailCry, 3600, "/singin-session", "0.0.0.0", true, true)
		c.String(http.StatusOK, "User successfully logged in !")
	}
	return
 */
