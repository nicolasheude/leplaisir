// Code generated by entc, DO NOT EDIT.

package contactparents

import (
	"dechild/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Nom applies equality check predicate on the "Nom" field. It's identical to NomEQ.
func Nom(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNom), v))
	})
}

// Prenom applies equality check predicate on the "Prenom" field. It's identical to PrenomEQ.
func Prenom(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrenom), v))
	})
}

// Sexe applies equality check predicate on the "Sexe" field. It's identical to SexeEQ.
func Sexe(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSexe), v))
	})
}

// Adresse applies equality check predicate on the "Adresse" field. It's identical to AdresseEQ.
func Adresse(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdresse), v))
	})
}

// Ville applies equality check predicate on the "Ville" field. It's identical to VilleEQ.
func Ville(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVille), v))
	})
}

// CP applies equality check predicate on the "CP" field. It's identical to CPEQ.
func CP(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCP), v))
	})
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Tel1 applies equality check predicate on the "Tel1" field. It's identical to Tel1EQ.
func Tel1(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel1), v))
	})
}

// Tel2 applies equality check predicate on the "Tel2" field. It's identical to Tel2EQ.
func Tel2(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel2), v))
	})
}

// NomEQ applies the EQ predicate on the "Nom" field.
func NomEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNom), v))
	})
}

// NomNEQ applies the NEQ predicate on the "Nom" field.
func NomNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNom), v))
	})
}

// NomIn applies the In predicate on the "Nom" field.
func NomIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNom), v...))
	})
}

// NomNotIn applies the NotIn predicate on the "Nom" field.
func NomNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNom), v...))
	})
}

// NomGT applies the GT predicate on the "Nom" field.
func NomGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNom), v))
	})
}

// NomGTE applies the GTE predicate on the "Nom" field.
func NomGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNom), v))
	})
}

// NomLT applies the LT predicate on the "Nom" field.
func NomLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNom), v))
	})
}

// NomLTE applies the LTE predicate on the "Nom" field.
func NomLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNom), v))
	})
}

// NomContains applies the Contains predicate on the "Nom" field.
func NomContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNom), v))
	})
}

// NomHasPrefix applies the HasPrefix predicate on the "Nom" field.
func NomHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNom), v))
	})
}

// NomHasSuffix applies the HasSuffix predicate on the "Nom" field.
func NomHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNom), v))
	})
}

// NomEqualFold applies the EqualFold predicate on the "Nom" field.
func NomEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNom), v))
	})
}

// NomContainsFold applies the ContainsFold predicate on the "Nom" field.
func NomContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNom), v))
	})
}

// PrenomEQ applies the EQ predicate on the "Prenom" field.
func PrenomEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrenom), v))
	})
}

// PrenomNEQ applies the NEQ predicate on the "Prenom" field.
func PrenomNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrenom), v))
	})
}

// PrenomIn applies the In predicate on the "Prenom" field.
func PrenomIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrenom), v...))
	})
}

// PrenomNotIn applies the NotIn predicate on the "Prenom" field.
func PrenomNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrenom), v...))
	})
}

// PrenomGT applies the GT predicate on the "Prenom" field.
func PrenomGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrenom), v))
	})
}

// PrenomGTE applies the GTE predicate on the "Prenom" field.
func PrenomGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrenom), v))
	})
}

// PrenomLT applies the LT predicate on the "Prenom" field.
func PrenomLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrenom), v))
	})
}

// PrenomLTE applies the LTE predicate on the "Prenom" field.
func PrenomLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrenom), v))
	})
}

// PrenomContains applies the Contains predicate on the "Prenom" field.
func PrenomContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrenom), v))
	})
}

// PrenomHasPrefix applies the HasPrefix predicate on the "Prenom" field.
func PrenomHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrenom), v))
	})
}

// PrenomHasSuffix applies the HasSuffix predicate on the "Prenom" field.
func PrenomHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrenom), v))
	})
}

// PrenomEqualFold applies the EqualFold predicate on the "Prenom" field.
func PrenomEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrenom), v))
	})
}

// PrenomContainsFold applies the ContainsFold predicate on the "Prenom" field.
func PrenomContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrenom), v))
	})
}

// SexeEQ applies the EQ predicate on the "Sexe" field.
func SexeEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSexe), v))
	})
}

// SexeNEQ applies the NEQ predicate on the "Sexe" field.
func SexeNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSexe), v))
	})
}

// SexeIn applies the In predicate on the "Sexe" field.
func SexeIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSexe), v...))
	})
}

// SexeNotIn applies the NotIn predicate on the "Sexe" field.
func SexeNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSexe), v...))
	})
}

// SexeGT applies the GT predicate on the "Sexe" field.
func SexeGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSexe), v))
	})
}

// SexeGTE applies the GTE predicate on the "Sexe" field.
func SexeGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSexe), v))
	})
}

// SexeLT applies the LT predicate on the "Sexe" field.
func SexeLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSexe), v))
	})
}

// SexeLTE applies the LTE predicate on the "Sexe" field.
func SexeLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSexe), v))
	})
}

// SexeContains applies the Contains predicate on the "Sexe" field.
func SexeContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSexe), v))
	})
}

// SexeHasPrefix applies the HasPrefix predicate on the "Sexe" field.
func SexeHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSexe), v))
	})
}

// SexeHasSuffix applies the HasSuffix predicate on the "Sexe" field.
func SexeHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSexe), v))
	})
}

// SexeEqualFold applies the EqualFold predicate on the "Sexe" field.
func SexeEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSexe), v))
	})
}

// SexeContainsFold applies the ContainsFold predicate on the "Sexe" field.
func SexeContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSexe), v))
	})
}

// AdresseEQ applies the EQ predicate on the "Adresse" field.
func AdresseEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdresse), v))
	})
}

// AdresseNEQ applies the NEQ predicate on the "Adresse" field.
func AdresseNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdresse), v))
	})
}

// AdresseIn applies the In predicate on the "Adresse" field.
func AdresseIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdresse), v...))
	})
}

// AdresseNotIn applies the NotIn predicate on the "Adresse" field.
func AdresseNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdresse), v...))
	})
}

// AdresseGT applies the GT predicate on the "Adresse" field.
func AdresseGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdresse), v))
	})
}

// AdresseGTE applies the GTE predicate on the "Adresse" field.
func AdresseGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdresse), v))
	})
}

// AdresseLT applies the LT predicate on the "Adresse" field.
func AdresseLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdresse), v))
	})
}

// AdresseLTE applies the LTE predicate on the "Adresse" field.
func AdresseLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdresse), v))
	})
}

// AdresseContains applies the Contains predicate on the "Adresse" field.
func AdresseContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAdresse), v))
	})
}

// AdresseHasPrefix applies the HasPrefix predicate on the "Adresse" field.
func AdresseHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAdresse), v))
	})
}

// AdresseHasSuffix applies the HasSuffix predicate on the "Adresse" field.
func AdresseHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAdresse), v))
	})
}

// AdresseEqualFold applies the EqualFold predicate on the "Adresse" field.
func AdresseEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAdresse), v))
	})
}

// AdresseContainsFold applies the ContainsFold predicate on the "Adresse" field.
func AdresseContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAdresse), v))
	})
}

// VilleEQ applies the EQ predicate on the "Ville" field.
func VilleEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVille), v))
	})
}

// VilleNEQ applies the NEQ predicate on the "Ville" field.
func VilleNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVille), v))
	})
}

// VilleIn applies the In predicate on the "Ville" field.
func VilleIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVille), v...))
	})
}

// VilleNotIn applies the NotIn predicate on the "Ville" field.
func VilleNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVille), v...))
	})
}

// VilleGT applies the GT predicate on the "Ville" field.
func VilleGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVille), v))
	})
}

// VilleGTE applies the GTE predicate on the "Ville" field.
func VilleGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVille), v))
	})
}

// VilleLT applies the LT predicate on the "Ville" field.
func VilleLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVille), v))
	})
}

// VilleLTE applies the LTE predicate on the "Ville" field.
func VilleLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVille), v))
	})
}

// VilleContains applies the Contains predicate on the "Ville" field.
func VilleContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVille), v))
	})
}

// VilleHasPrefix applies the HasPrefix predicate on the "Ville" field.
func VilleHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVille), v))
	})
}

// VilleHasSuffix applies the HasSuffix predicate on the "Ville" field.
func VilleHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVille), v))
	})
}

// VilleEqualFold applies the EqualFold predicate on the "Ville" field.
func VilleEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVille), v))
	})
}

// VilleContainsFold applies the ContainsFold predicate on the "Ville" field.
func VilleContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVille), v))
	})
}

// CPEQ applies the EQ predicate on the "CP" field.
func CPEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCP), v))
	})
}

// CPNEQ applies the NEQ predicate on the "CP" field.
func CPNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCP), v))
	})
}

// CPIn applies the In predicate on the "CP" field.
func CPIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCP), v...))
	})
}

// CPNotIn applies the NotIn predicate on the "CP" field.
func CPNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCP), v...))
	})
}

// CPGT applies the GT predicate on the "CP" field.
func CPGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCP), v))
	})
}

// CPGTE applies the GTE predicate on the "CP" field.
func CPGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCP), v))
	})
}

// CPLT applies the LT predicate on the "CP" field.
func CPLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCP), v))
	})
}

// CPLTE applies the LTE predicate on the "CP" field.
func CPLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCP), v))
	})
}

// CPContains applies the Contains predicate on the "CP" field.
func CPContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCP), v))
	})
}

// CPHasPrefix applies the HasPrefix predicate on the "CP" field.
func CPHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCP), v))
	})
}

// CPHasSuffix applies the HasSuffix predicate on the "CP" field.
func CPHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCP), v))
	})
}

// CPEqualFold applies the EqualFold predicate on the "CP" field.
func CPEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCP), v))
	})
}

// CPContainsFold applies the ContainsFold predicate on the "CP" field.
func CPContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCP), v))
	})
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// Tel1EQ applies the EQ predicate on the "Tel1" field.
func Tel1EQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel1), v))
	})
}

// Tel1NEQ applies the NEQ predicate on the "Tel1" field.
func Tel1NEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTel1), v))
	})
}

// Tel1In applies the In predicate on the "Tel1" field.
func Tel1In(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTel1), v...))
	})
}

// Tel1NotIn applies the NotIn predicate on the "Tel1" field.
func Tel1NotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTel1), v...))
	})
}

// Tel1GT applies the GT predicate on the "Tel1" field.
func Tel1GT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTel1), v))
	})
}

// Tel1GTE applies the GTE predicate on the "Tel1" field.
func Tel1GTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTel1), v))
	})
}

// Tel1LT applies the LT predicate on the "Tel1" field.
func Tel1LT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTel1), v))
	})
}

// Tel1LTE applies the LTE predicate on the "Tel1" field.
func Tel1LTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTel1), v))
	})
}

// Tel1Contains applies the Contains predicate on the "Tel1" field.
func Tel1Contains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTel1), v))
	})
}

// Tel1HasPrefix applies the HasPrefix predicate on the "Tel1" field.
func Tel1HasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTel1), v))
	})
}

// Tel1HasSuffix applies the HasSuffix predicate on the "Tel1" field.
func Tel1HasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTel1), v))
	})
}

// Tel1EqualFold applies the EqualFold predicate on the "Tel1" field.
func Tel1EqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTel1), v))
	})
}

// Tel1ContainsFold applies the ContainsFold predicate on the "Tel1" field.
func Tel1ContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTel1), v))
	})
}

// Tel2EQ applies the EQ predicate on the "Tel2" field.
func Tel2EQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel2), v))
	})
}

// Tel2NEQ applies the NEQ predicate on the "Tel2" field.
func Tel2NEQ(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTel2), v))
	})
}

// Tel2In applies the In predicate on the "Tel2" field.
func Tel2In(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTel2), v...))
	})
}

// Tel2NotIn applies the NotIn predicate on the "Tel2" field.
func Tel2NotIn(vs ...string) predicate.ContactParents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ContactParents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTel2), v...))
	})
}

// Tel2GT applies the GT predicate on the "Tel2" field.
func Tel2GT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTel2), v))
	})
}

// Tel2GTE applies the GTE predicate on the "Tel2" field.
func Tel2GTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTel2), v))
	})
}

// Tel2LT applies the LT predicate on the "Tel2" field.
func Tel2LT(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTel2), v))
	})
}

// Tel2LTE applies the LTE predicate on the "Tel2" field.
func Tel2LTE(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTel2), v))
	})
}

// Tel2Contains applies the Contains predicate on the "Tel2" field.
func Tel2Contains(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTel2), v))
	})
}

// Tel2HasPrefix applies the HasPrefix predicate on the "Tel2" field.
func Tel2HasPrefix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTel2), v))
	})
}

// Tel2HasSuffix applies the HasSuffix predicate on the "Tel2" field.
func Tel2HasSuffix(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTel2), v))
	})
}

// Tel2IsNil applies the IsNil predicate on the "Tel2" field.
func Tel2IsNil() predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTel2)))
	})
}

// Tel2NotNil applies the NotNil predicate on the "Tel2" field.
func Tel2NotNil() predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTel2)))
	})
}

// Tel2EqualFold applies the EqualFold predicate on the "Tel2" field.
func Tel2EqualFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTel2), v))
	})
}

// Tel2ContainsFold applies the ContainsFold predicate on the "Tel2" field.
func Tel2ContainsFold(v string) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTel2), v))
	})
}

// HasChild applies the HasEdge predicate on the "child" edge.
func HasChild() predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ChildTable, ChildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildWith applies the HasEdge predicate on the "child" edge with a given conditions (other predicates).
func HasChildWith(preds ...predicate.Form) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ChildTable, ChildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ContactParents) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ContactParents) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ContactParents) predicate.ContactParents {
	return predicate.ContactParents(func(s *sql.Selector) {
		p(s.Not())
	})
}
