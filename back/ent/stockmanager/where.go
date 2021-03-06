// Code generated by entc, DO NOT EDIT.

package stockmanager

import (
	"dechild/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
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
func IDGT(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Activite applies equality check predicate on the "Activite" field. It's identical to ActiviteEQ.
func Activite(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActivite), v))
	})
}

// SemaineA applies equality check predicate on the "SemaineA" field. It's identical to SemaineAEQ.
func SemaineA(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineA), v))
	})
}

// SemaineB applies equality check predicate on the "SemaineB" field. It's identical to SemaineBEQ.
func SemaineB(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineB), v))
	})
}

// SemaineC applies equality check predicate on the "SemaineC" field. It's identical to SemaineCEQ.
func SemaineC(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineC), v))
	})
}

// SemaineD applies equality check predicate on the "SemaineD" field. It's identical to SemaineDEQ.
func SemaineD(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineD), v))
	})
}

// SemaineE applies equality check predicate on the "SemaineE" field. It's identical to SemaineEEQ.
func SemaineE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineE), v))
	})
}

// SemaineF applies equality check predicate on the "SemaineF" field. It's identical to SemaineFEQ.
func SemaineF(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineF), v))
	})
}

// SemaineG applies equality check predicate on the "SemaineG" field. It's identical to SemaineGEQ.
func SemaineG(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineG), v))
	})
}

// SemaineH applies equality check predicate on the "SemaineH" field. It's identical to SemaineHEQ.
func SemaineH(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineH), v))
	})
}

// ActiviteEQ applies the EQ predicate on the "Activite" field.
func ActiviteEQ(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActivite), v))
	})
}

// ActiviteNEQ applies the NEQ predicate on the "Activite" field.
func ActiviteNEQ(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActivite), v))
	})
}

// ActiviteIn applies the In predicate on the "Activite" field.
func ActiviteIn(vs ...string) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActivite), v...))
	})
}

// ActiviteNotIn applies the NotIn predicate on the "Activite" field.
func ActiviteNotIn(vs ...string) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActivite), v...))
	})
}

// ActiviteGT applies the GT predicate on the "Activite" field.
func ActiviteGT(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActivite), v))
	})
}

// ActiviteGTE applies the GTE predicate on the "Activite" field.
func ActiviteGTE(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActivite), v))
	})
}

// ActiviteLT applies the LT predicate on the "Activite" field.
func ActiviteLT(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActivite), v))
	})
}

// ActiviteLTE applies the LTE predicate on the "Activite" field.
func ActiviteLTE(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActivite), v))
	})
}

// ActiviteContains applies the Contains predicate on the "Activite" field.
func ActiviteContains(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActivite), v))
	})
}

// ActiviteHasPrefix applies the HasPrefix predicate on the "Activite" field.
func ActiviteHasPrefix(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActivite), v))
	})
}

// ActiviteHasSuffix applies the HasSuffix predicate on the "Activite" field.
func ActiviteHasSuffix(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActivite), v))
	})
}

// ActiviteEqualFold applies the EqualFold predicate on the "Activite" field.
func ActiviteEqualFold(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActivite), v))
	})
}

// ActiviteContainsFold applies the ContainsFold predicate on the "Activite" field.
func ActiviteContainsFold(v string) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActivite), v))
	})
}

// SemaineAEQ applies the EQ predicate on the "SemaineA" field.
func SemaineAEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineA), v))
	})
}

// SemaineANEQ applies the NEQ predicate on the "SemaineA" field.
func SemaineANEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineA), v))
	})
}

// SemaineAIn applies the In predicate on the "SemaineA" field.
func SemaineAIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineA), v...))
	})
}

// SemaineANotIn applies the NotIn predicate on the "SemaineA" field.
func SemaineANotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineA), v...))
	})
}

// SemaineAGT applies the GT predicate on the "SemaineA" field.
func SemaineAGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineA), v))
	})
}

// SemaineAGTE applies the GTE predicate on the "SemaineA" field.
func SemaineAGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineA), v))
	})
}

// SemaineALT applies the LT predicate on the "SemaineA" field.
func SemaineALT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineA), v))
	})
}

// SemaineALTE applies the LTE predicate on the "SemaineA" field.
func SemaineALTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineA), v))
	})
}

// SemaineBEQ applies the EQ predicate on the "SemaineB" field.
func SemaineBEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineB), v))
	})
}

// SemaineBNEQ applies the NEQ predicate on the "SemaineB" field.
func SemaineBNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineB), v))
	})
}

// SemaineBIn applies the In predicate on the "SemaineB" field.
func SemaineBIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineB), v...))
	})
}

// SemaineBNotIn applies the NotIn predicate on the "SemaineB" field.
func SemaineBNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineB), v...))
	})
}

// SemaineBGT applies the GT predicate on the "SemaineB" field.
func SemaineBGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineB), v))
	})
}

// SemaineBGTE applies the GTE predicate on the "SemaineB" field.
func SemaineBGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineB), v))
	})
}

// SemaineBLT applies the LT predicate on the "SemaineB" field.
func SemaineBLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineB), v))
	})
}

// SemaineBLTE applies the LTE predicate on the "SemaineB" field.
func SemaineBLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineB), v))
	})
}

// SemaineCEQ applies the EQ predicate on the "SemaineC" field.
func SemaineCEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineC), v))
	})
}

// SemaineCNEQ applies the NEQ predicate on the "SemaineC" field.
func SemaineCNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineC), v))
	})
}

// SemaineCIn applies the In predicate on the "SemaineC" field.
func SemaineCIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineC), v...))
	})
}

// SemaineCNotIn applies the NotIn predicate on the "SemaineC" field.
func SemaineCNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineC), v...))
	})
}

// SemaineCGT applies the GT predicate on the "SemaineC" field.
func SemaineCGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineC), v))
	})
}

// SemaineCGTE applies the GTE predicate on the "SemaineC" field.
func SemaineCGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineC), v))
	})
}

// SemaineCLT applies the LT predicate on the "SemaineC" field.
func SemaineCLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineC), v))
	})
}

// SemaineCLTE applies the LTE predicate on the "SemaineC" field.
func SemaineCLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineC), v))
	})
}

// SemaineDEQ applies the EQ predicate on the "SemaineD" field.
func SemaineDEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineD), v))
	})
}

// SemaineDNEQ applies the NEQ predicate on the "SemaineD" field.
func SemaineDNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineD), v))
	})
}

// SemaineDIn applies the In predicate on the "SemaineD" field.
func SemaineDIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineD), v...))
	})
}

// SemaineDNotIn applies the NotIn predicate on the "SemaineD" field.
func SemaineDNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineD), v...))
	})
}

// SemaineDGT applies the GT predicate on the "SemaineD" field.
func SemaineDGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineD), v))
	})
}

// SemaineDGTE applies the GTE predicate on the "SemaineD" field.
func SemaineDGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineD), v))
	})
}

// SemaineDLT applies the LT predicate on the "SemaineD" field.
func SemaineDLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineD), v))
	})
}

// SemaineDLTE applies the LTE predicate on the "SemaineD" field.
func SemaineDLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineD), v))
	})
}

// SemaineEEQ applies the EQ predicate on the "SemaineE" field.
func SemaineEEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineE), v))
	})
}

// SemaineENEQ applies the NEQ predicate on the "SemaineE" field.
func SemaineENEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineE), v))
	})
}

// SemaineEIn applies the In predicate on the "SemaineE" field.
func SemaineEIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineE), v...))
	})
}

// SemaineENotIn applies the NotIn predicate on the "SemaineE" field.
func SemaineENotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineE), v...))
	})
}

// SemaineEGT applies the GT predicate on the "SemaineE" field.
func SemaineEGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineE), v))
	})
}

// SemaineEGTE applies the GTE predicate on the "SemaineE" field.
func SemaineEGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineE), v))
	})
}

// SemaineELT applies the LT predicate on the "SemaineE" field.
func SemaineELT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineE), v))
	})
}

// SemaineELTE applies the LTE predicate on the "SemaineE" field.
func SemaineELTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineE), v))
	})
}

// SemaineFEQ applies the EQ predicate on the "SemaineF" field.
func SemaineFEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineF), v))
	})
}

// SemaineFNEQ applies the NEQ predicate on the "SemaineF" field.
func SemaineFNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineF), v))
	})
}

// SemaineFIn applies the In predicate on the "SemaineF" field.
func SemaineFIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineF), v...))
	})
}

// SemaineFNotIn applies the NotIn predicate on the "SemaineF" field.
func SemaineFNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineF), v...))
	})
}

// SemaineFGT applies the GT predicate on the "SemaineF" field.
func SemaineFGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineF), v))
	})
}

// SemaineFGTE applies the GTE predicate on the "SemaineF" field.
func SemaineFGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineF), v))
	})
}

// SemaineFLT applies the LT predicate on the "SemaineF" field.
func SemaineFLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineF), v))
	})
}

// SemaineFLTE applies the LTE predicate on the "SemaineF" field.
func SemaineFLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineF), v))
	})
}

// SemaineGEQ applies the EQ predicate on the "SemaineG" field.
func SemaineGEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineG), v))
	})
}

// SemaineGNEQ applies the NEQ predicate on the "SemaineG" field.
func SemaineGNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineG), v))
	})
}

// SemaineGIn applies the In predicate on the "SemaineG" field.
func SemaineGIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineG), v...))
	})
}

// SemaineGNotIn applies the NotIn predicate on the "SemaineG" field.
func SemaineGNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineG), v...))
	})
}

// SemaineGGT applies the GT predicate on the "SemaineG" field.
func SemaineGGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineG), v))
	})
}

// SemaineGGTE applies the GTE predicate on the "SemaineG" field.
func SemaineGGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineG), v))
	})
}

// SemaineGLT applies the LT predicate on the "SemaineG" field.
func SemaineGLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineG), v))
	})
}

// SemaineGLTE applies the LTE predicate on the "SemaineG" field.
func SemaineGLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineG), v))
	})
}

// SemaineHEQ applies the EQ predicate on the "SemaineH" field.
func SemaineHEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSemaineH), v))
	})
}

// SemaineHNEQ applies the NEQ predicate on the "SemaineH" field.
func SemaineHNEQ(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSemaineH), v))
	})
}

// SemaineHIn applies the In predicate on the "SemaineH" field.
func SemaineHIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSemaineH), v...))
	})
}

// SemaineHNotIn applies the NotIn predicate on the "SemaineH" field.
func SemaineHNotIn(vs ...int) predicate.StockManager {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StockManager(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSemaineH), v...))
	})
}

// SemaineHGT applies the GT predicate on the "SemaineH" field.
func SemaineHGT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSemaineH), v))
	})
}

// SemaineHGTE applies the GTE predicate on the "SemaineH" field.
func SemaineHGTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSemaineH), v))
	})
}

// SemaineHLT applies the LT predicate on the "SemaineH" field.
func SemaineHLT(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSemaineH), v))
	})
}

// SemaineHLTE applies the LTE predicate on the "SemaineH" field.
func SemaineHLTE(v int) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSemaineH), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StockManager) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StockManager) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
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
func Not(p predicate.StockManager) predicate.StockManager {
	return predicate.StockManager(func(s *sql.Selector) {
		p(s.Not())
	})
}
