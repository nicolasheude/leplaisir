package router

import (
	"context"
	"dechild/controllers"
	"dechild/database"
	"dechild/ent"
	"dechild/ent/form"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataSemaine struct {
	Semaine string `json:"semaine"`
	Tab 	[]DataAdmin `json:"tab"`
}

type DataAdmin struct {
	Nom 		 	string `json:"nom"`
	Prenom 		 	string `json:"prenom"`
	Anniversaire	string `json:"anniversaire"`
	Sexe 		 	string `json:"sexe"`
	Niveau 		 	string `json:"niveau"`
	Location 	 	string `json:"location"`
	Repas 		 	string `json:"repas"`
	Formule 	 	string `json:"formule"`
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

func SelecWeek(Semaine string, ctx context.Context, client *ent.Client) (DataSemaine, error) {
	myTab := DataSemaine{}
	db, err := client.Form.
		Query().
		Where(form.Semaine(Semaine)).
		All(ctx)
	if err != nil {
		return myTab, fmt.Errorf("failed creating new parent: %w", err)
	}
	myTab.Tab = make([]DataAdmin, len(db))
	myTab.Semaine = Semaine
	for i := range db {
		myTab.Tab[i] = DataAdmin{
			Nom:          db[i].Nom,
			Prenom:       db[i].Prenom,
			Anniversaire: db[i].Anniversaire,
			Sexe:         db[i].Sexe,
			Niveau:       db[i].Niveau,
			Location:     db[i].LocationR,
			Repas:        db[i].Repas,
			Formule:     db[i].Formule,
		}
	}
	return myTab, nil
}

func GetDbWeek(c *gin.Context) {
	semaine := c.Query("semaine")
	tab, err := SelecWeek(semaine, database.Db.Ctx, database.Db.Def)
	if err != nil {
		fmt.Errorf("Error : %w", err)
		return
	}
	c.JSON(http.StatusOK, tab)
}

func ApplyRoutes(r *gin.Engine) {
	r.POST("/form", GetForm)
	r.GET("/semaine", GetDbWeek)
	r.POST("/admin", controllers.LoginSession)
}