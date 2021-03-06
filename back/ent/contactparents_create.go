// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"dechild/ent/contactparents"
	"dechild/ent/form"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContactParentsCreate is the builder for creating a ContactParents entity.
type ContactParentsCreate struct {
	config
	mutation *ContactParentsMutation
	hooks    []Hook
}

// SetNom sets the "Nom" field.
func (cpc *ContactParentsCreate) SetNom(s string) *ContactParentsCreate {
	cpc.mutation.SetNom(s)
	return cpc
}

// SetPrenom sets the "Prenom" field.
func (cpc *ContactParentsCreate) SetPrenom(s string) *ContactParentsCreate {
	cpc.mutation.SetPrenom(s)
	return cpc
}

// SetSexe sets the "Sexe" field.
func (cpc *ContactParentsCreate) SetSexe(s string) *ContactParentsCreate {
	cpc.mutation.SetSexe(s)
	return cpc
}

// SetAdresse sets the "Adresse" field.
func (cpc *ContactParentsCreate) SetAdresse(s string) *ContactParentsCreate {
	cpc.mutation.SetAdresse(s)
	return cpc
}

// SetVille sets the "Ville" field.
func (cpc *ContactParentsCreate) SetVille(s string) *ContactParentsCreate {
	cpc.mutation.SetVille(s)
	return cpc
}

// SetCP sets the "CP" field.
func (cpc *ContactParentsCreate) SetCP(s string) *ContactParentsCreate {
	cpc.mutation.SetCP(s)
	return cpc
}

// SetEmail sets the "Email" field.
func (cpc *ContactParentsCreate) SetEmail(s string) *ContactParentsCreate {
	cpc.mutation.SetEmail(s)
	return cpc
}

// SetTel1 sets the "Tel1" field.
func (cpc *ContactParentsCreate) SetTel1(s string) *ContactParentsCreate {
	cpc.mutation.SetTel1(s)
	return cpc
}

// SetTel2 sets the "Tel2" field.
func (cpc *ContactParentsCreate) SetTel2(s string) *ContactParentsCreate {
	cpc.mutation.SetTel2(s)
	return cpc
}

// SetNillableTel2 sets the "Tel2" field if the given value is not nil.
func (cpc *ContactParentsCreate) SetNillableTel2(s *string) *ContactParentsCreate {
	if s != nil {
		cpc.SetTel2(*s)
	}
	return cpc
}

// SetID sets the "id" field.
func (cpc *ContactParentsCreate) SetID(i int) *ContactParentsCreate {
	cpc.mutation.SetID(i)
	return cpc
}

// AddChildIDs adds the "child" edge to the Form entity by IDs.
func (cpc *ContactParentsCreate) AddChildIDs(ids ...int) *ContactParentsCreate {
	cpc.mutation.AddChildIDs(ids...)
	return cpc
}

// AddChild adds the "child" edges to the Form entity.
func (cpc *ContactParentsCreate) AddChild(f ...*Form) *ContactParentsCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cpc.AddChildIDs(ids...)
}

// Mutation returns the ContactParentsMutation object of the builder.
func (cpc *ContactParentsCreate) Mutation() *ContactParentsMutation {
	return cpc.mutation
}

// Save creates the ContactParents in the database.
func (cpc *ContactParentsCreate) Save(ctx context.Context) (*ContactParents, error) {
	var (
		err  error
		node *ContactParents
	)
	if len(cpc.hooks) == 0 {
		if err = cpc.check(); err != nil {
			return nil, err
		}
		node, err = cpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactParentsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpc.check(); err != nil {
				return nil, err
			}
			cpc.mutation = mutation
			node, err = cpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpc.hooks) - 1; i >= 0; i-- {
			mut = cpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *ContactParentsCreate) SaveX(ctx context.Context) *ContactParents {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cpc *ContactParentsCreate) check() error {
	if _, ok := cpc.mutation.Nom(); !ok {
		return &ValidationError{Name: "Nom", err: errors.New("ent: missing required field \"Nom\"")}
	}
	if _, ok := cpc.mutation.Prenom(); !ok {
		return &ValidationError{Name: "Prenom", err: errors.New("ent: missing required field \"Prenom\"")}
	}
	if _, ok := cpc.mutation.Sexe(); !ok {
		return &ValidationError{Name: "Sexe", err: errors.New("ent: missing required field \"Sexe\"")}
	}
	if _, ok := cpc.mutation.Adresse(); !ok {
		return &ValidationError{Name: "Adresse", err: errors.New("ent: missing required field \"Adresse\"")}
	}
	if _, ok := cpc.mutation.Ville(); !ok {
		return &ValidationError{Name: "Ville", err: errors.New("ent: missing required field \"Ville\"")}
	}
	if _, ok := cpc.mutation.CP(); !ok {
		return &ValidationError{Name: "CP", err: errors.New("ent: missing required field \"CP\"")}
	}
	if _, ok := cpc.mutation.Email(); !ok {
		return &ValidationError{Name: "Email", err: errors.New("ent: missing required field \"Email\"")}
	}
	if _, ok := cpc.mutation.Tel1(); !ok {
		return &ValidationError{Name: "Tel1", err: errors.New("ent: missing required field \"Tel1\"")}
	}
	if v, ok := cpc.mutation.ID(); ok {
		if err := contactparents.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	if len(cpc.mutation.ChildIDs()) == 0 {
		return &ValidationError{Name: "child", err: errors.New("ent: missing required edge \"child\"")}
	}
	return nil
}

func (cpc *ContactParentsCreate) sqlSave(ctx context.Context) (*ContactParents, error) {
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (cpc *ContactParentsCreate) createSpec() (*ContactParents, *sqlgraph.CreateSpec) {
	var (
		_node = &ContactParents{config: cpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contactparents.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contactparents.FieldID,
			},
		}
	)
	if id, ok := cpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cpc.mutation.Nom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldNom,
		})
		_node.Nom = value
	}
	if value, ok := cpc.mutation.Prenom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldPrenom,
		})
		_node.Prenom = value
	}
	if value, ok := cpc.mutation.Sexe(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldSexe,
		})
		_node.Sexe = value
	}
	if value, ok := cpc.mutation.Adresse(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldAdresse,
		})
		_node.Adresse = value
	}
	if value, ok := cpc.mutation.Ville(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldVille,
		})
		_node.Ville = value
	}
	if value, ok := cpc.mutation.CP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldCP,
		})
		_node.CP = value
	}
	if value, ok := cpc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := cpc.mutation.Tel1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldTel1,
		})
		_node.Tel1 = value
	}
	if value, ok := cpc.mutation.Tel2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactparents.FieldTel2,
		})
		_node.Tel2 = value
	}
	if nodes := cpc.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   contactparents.ChildTable,
			Columns: []string{contactparents.ChildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: form.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContactParentsCreateBulk is the builder for creating many ContactParents entities in bulk.
type ContactParentsCreateBulk struct {
	config
	builders []*ContactParentsCreate
}

// Save creates the ContactParents entities in the database.
func (cpcb *ContactParentsCreateBulk) Save(ctx context.Context) ([]*ContactParents, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*ContactParents, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactParentsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *ContactParentsCreateBulk) SaveX(ctx context.Context) []*ContactParents {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
