package controllers

import (
	"context"
	"dechild/ent"
	"dechild/ent/form"
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
		return nil, fmt.Errorf("failed creating new parents: %w", err)
	}
	log.Println("client was created: ", newClient)
	return newClient, nil
}

/*
func AvailabilityCheck(Data FormData, ctx context.Context, client *ent.Client) bool {
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
}*/

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

/*
func CheckManager(Data FormData) bool {
	if AvailabilityCheck(Data) != true && CheckDoesNotExist(Data) != true {
		return false
	}
	return true
}*/

// CreateForm Create a developer from its ID
func CreateFrom(Data FormData, ctx context.Context, client *ent.Client) (*ent.Form, error) {
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
		return nil, fmt.Errorf("failed creating new client: %w", err)
	}
	_, err = CreateContactParents(Data.Parents, newClient, ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed creating new parent: %w", err)
	}
	log.Println("client was created: ", newClient)
	return newClient, nil
}
