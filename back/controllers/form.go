package controllers

import (
	"context"
	"dechild/database"
	"dechild/ent"
	"dechild/ent/form"
	"dechild/ent/stockmanager"
	"fmt"
	"log"
)

type FormData struct {
	Child   Form           `json:"child"`
	Parents ContactParents `json:"parents"`
}

type Form struct {
	Email        string `json:"email"`
	Nom          string `json:"nom"`
	Prenom       string `json:"prenom"`
	Anniversaire string `json:"anniversaire"`
	Nationalite  string `json:"nationalite"`
	Sexe         string `json:"sexe"`
	Niveau       string `json:"niveau"`
	Fft          string `json:"fft"`
	LocationR    string `json:"locationR"`
	Repas        string `json:"repas"`
	Formule      string `json:"formule"`
	Semaine      string `json:"semaine"`
}

type ContactParents struct {
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Sexe    string `json:"sexe"`
	Adresse string `json:"adresse"`
	Ville   string `json:"ville"`
	CP      string `json:"cp"`
	Email   string `json:"email"`
	Tel1    string `json:"tel_1"`
	Tel2    string `json:"tel_2"`
}

var Enfant Form

var Parents ContactParents

var Data FormData

func SetStockManager(activite string, nbPlace int, ctx context.Context, client *ent.Client) {
	client.StockManager.
		Create().
		SetActivite(activite).
		SetSemaineA(nbPlace).
		SetSemaineB(nbPlace).
		SetSemaineC(nbPlace).
		SetSemaineD(nbPlace).
		SetSemaineE(nbPlace).
		SetSemaineF(nbPlace).
		SetSemaineG(nbPlace).
		SetSemaineH(nbPlace).
		Save(ctx)
}

func InitStockManager(ctx context.Context, client *ent.Client) {
	SetStockManager("Mini-Tennis", 6, ctx, client)
	SetStockManager("Tennis Times", 8, ctx, client)
	SetStockManager("Tennis + Activitée", 16, ctx, client)
	SetStockManager("Tennis + Wakeboard", 12, ctx, client)
	SetStockManager("Tennis + Voile", 10, ctx, client)
}

// CreateContactParents Create a developer from its ID
func CreateContactParents(Client ContactParents, newChild *ent.Form, ctx context.Context, client *ent.Client) (*ent.ContactParents, error) {
	newClient, err := client.ContactParents.
		Create().
		SetNom(Client.Nom).
		SetPrenom(Client.Prenom).
		SetSexe(Client.Sexe).
		SetAdresse(Client.Adresse).
		SetVille(Client.Ville).
		SetCP(Client.CP).
		SetEmail(Client.Email).
		SetTel1(Client.Tel1).
		SetTel2(Client.Tel2).
		AddChild(newChild).
		Save(ctx)
	if err != nil {
		fmt.Printf("Ligne 6 : [%w]", err)
		return nil, fmt.Errorf("failed creating new parents: %w", err)
	}
	log.Println("client was created: ", newClient)
	return newClient, nil
}

func getNb(c *ent.StockManager, semaine string) (int, error) {
	if semaine == "SemaineA" {
		return c.SemaineA, nil
	}
	if semaine == "SemaineB" {
		return c.SemaineB, nil
	}
	if semaine == "SemaineC" {
		return c.SemaineC, nil
	}
	if semaine == "SemaineD" {
		return c.SemaineD, nil
	}
	if semaine == "SemaineE" {
		return c.SemaineE, nil
	}
	if semaine == "SemaineF" {
		return c.SemaineF, nil
	}
	if semaine == "SemaineG" {
		return c.SemaineG, nil
	}
	if semaine == "SemaineH" {
		return c.SemaineH, nil
	}
	return -1, fmt.Errorf("Week not found")
}

func setNb(client *ent.StockManager, nb int, semaine string) (int, error) {
	fmt.Printf("LE NOMBRE AVANT DE LE RETIRER : %d\n", nb)
	if semaine == "SemaineA" {
		client.
			Update().
			SetSemaineA(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineB" {
		client.
			Update().
			SetSemaineB(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineC" {
		client.
			Update().
			SetSemaineC(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineD" {
		client.
			Update().
			SetSemaineD(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineE" {
		client.
			Update().
			SetSemaineE(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineF" {
		client.
			Update().
			SetSemaineF(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineG" {
		client.
			Update().
			SetSemaineG(nb - 1).
			Save(database.Db.Ctx)
	}
	if semaine == "SemaineH" {
		client.
			Update().
			SetSemaineH(nb - 1).
			Save(database.Db.Ctx)
	}
	return -1, fmt.Errorf("Week not found")
}

func AvailabilityCheck(Act, Semaine string, ctx context.Context, client *ent.Client) bool {
	db, err := client.StockManager.
		Query().
		Where(stockmanager.Activite(Act)).
		Only(ctx)
	fmt.Println(err)
	if err != nil {
		return false
	}
	nb, err := getNb(db, Semaine)
	if err != nil {
		return false
	}
	if nb > 0 {
		setNb(db, nb, Semaine)
		return true
	} else {
		return false
	}
}

func CheckDoesNotExist(Data FormData, ctx context.Context, client *ent.Client) bool {
	db, _ := client.Form.
		Query().
		Where(
			form.And(
				form.Nom(Data.Child.Nom),
				form.Prenom(Data.Child.Prenom),
				form.Anniversaire(Data.Child.Anniversaire),
				form.Semaine(Data.Child.Semaine)),
		).
		All(ctx)
	if len(db) == 0 {
		return true
	}
	return false
}

func CheckManager(Data FormData) bool {
	if AvailabilityCheck(Data.Child.Formule, Data.Child.Semaine, database.Db.Ctx, database.Db.Def) == false || CheckDoesNotExist(Data, database.Db.Ctx, database.Db.Def) == false {
		return false
	}
	return true
}

// CreateForm Create a developer from its ID
func CreateFrom(Data FormData, ctx context.Context, client *ent.Client) (*ent.Form, error) {
	if CheckManager(Data) == false {
		return nil, fmt.Errorf("Place not available or child already registered")
	}
	newClient, err := client.Form.
		Create().
		SetEmail(Data.Child.Email).
		SetNom(Data.Child.Nom).
		SetPrenom(Data.Child.Prenom).
		SetAnniversaire(Data.Child.Anniversaire).
		SetNationalite(Data.Child.Nationalite).
		SetSexe(Data.Child.Sexe).
		SetNiveau(Data.Child.Niveau).
		SetFFT(Data.Child.Fft).
		SetLocationR(Data.Child.LocationR).
		SetRepas(Data.Child.Repas).
		SetFormule(Data.Child.Formule).
		SetSemaine(Data.Child.Semaine).
		Save(ctx)
	if err != nil {
		fmt.Printf("Ligne 3 : [%w]", err)
		return nil, fmt.Errorf("failed creating new client: %w", err)
	}
	_, err = CreateContactParents(Data.Parents, newClient, ctx, client)
	if err != nil {
		fmt.Printf("Ligne 4 : [%w]", err)
		return nil, fmt.Errorf("failed creating new parent: %w", err)
	}
	log.Println("client was created: ", newClient)
	return newClient, nil
}
