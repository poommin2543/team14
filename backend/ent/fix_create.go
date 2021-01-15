// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
	"github.com/tanapon395/playlist-video/ent/fixcomtype"
	"github.com/tanapon395/playlist-video/ent/personal"
)

// FixCreate is the builder for creating a Fix entity.
type FixCreate struct {
	config
	mutation *FixMutation
	hooks    []Hook
}

// SetProductnumber sets the productnumber field.
func (fc *FixCreate) SetProductnumber(s string) *FixCreate {
	fc.mutation.SetProductnumber(s)
	return fc
}

// SetProblemtype sets the problemtype field.
func (fc *FixCreate) SetProblemtype(s string) *FixCreate {
	fc.mutation.SetProblemtype(s)
	return fc
}

// SetQueue sets the queue field.
func (fc *FixCreate) SetQueue(s string) *FixCreate {
	fc.mutation.SetQueue(s)
	return fc
}

// SetDate sets the date field.
func (fc *FixCreate) SetDate(t time.Time) *FixCreate {
	fc.mutation.SetDate(t)
	return fc
}

// SetFixID sets the fix edge to Adminrepair by id.
func (fc *FixCreate) SetFixID(id int) *FixCreate {
	fc.mutation.SetFixID(id)
	return fc
}

// SetNillableFixID sets the fix edge to Adminrepair by id if the given value is not nil.
func (fc *FixCreate) SetNillableFixID(id *int) *FixCreate {
	if id != nil {
		fc = fc.SetFixID(*id)
	}
	return fc
}

// SetFix sets the fix edge to Adminrepair.
func (fc *FixCreate) SetFix(a *Adminrepair) *FixCreate {
	return fc.SetFixID(a.ID)
}

// SetFixbrandID sets the fixbrand edge to Fixbrand by id.
func (fc *FixCreate) SetFixbrandID(id int) *FixCreate {
	fc.mutation.SetFixbrandID(id)
	return fc
}

// SetNillableFixbrandID sets the fixbrand edge to Fixbrand by id if the given value is not nil.
func (fc *FixCreate) SetNillableFixbrandID(id *int) *FixCreate {
	if id != nil {
		fc = fc.SetFixbrandID(*id)
	}
	return fc
}

// SetFixbrand sets the fixbrand edge to Fixbrand.
func (fc *FixCreate) SetFixbrand(f *Fixbrand) *FixCreate {
	return fc.SetFixbrandID(f.ID)
}

// SetPersonalID sets the personal edge to Personal by id.
func (fc *FixCreate) SetPersonalID(id int) *FixCreate {
	fc.mutation.SetPersonalID(id)
	return fc
}

// SetNillablePersonalID sets the personal edge to Personal by id if the given value is not nil.
func (fc *FixCreate) SetNillablePersonalID(id *int) *FixCreate {
	if id != nil {
		fc = fc.SetPersonalID(*id)
	}
	return fc
}

// SetPersonal sets the personal edge to Personal.
func (fc *FixCreate) SetPersonal(p *Personal) *FixCreate {
	return fc.SetPersonalID(p.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (fc *FixCreate) SetCustomerID(id int) *FixCreate {
	fc.mutation.SetCustomerID(id)
	return fc
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (fc *FixCreate) SetNillableCustomerID(id *int) *FixCreate {
	if id != nil {
		fc = fc.SetCustomerID(*id)
	}
	return fc
}

// SetCustomer sets the customer edge to Customer.
func (fc *FixCreate) SetCustomer(c *Customer) *FixCreate {
	return fc.SetCustomerID(c.ID)
}

// SetFixcomtypeID sets the fixcomtype edge to Fixcomtype by id.
func (fc *FixCreate) SetFixcomtypeID(id int) *FixCreate {
	fc.mutation.SetFixcomtypeID(id)
	return fc
}

// SetNillableFixcomtypeID sets the fixcomtype edge to Fixcomtype by id if the given value is not nil.
func (fc *FixCreate) SetNillableFixcomtypeID(id *int) *FixCreate {
	if id != nil {
		fc = fc.SetFixcomtypeID(*id)
	}
	return fc
}

// SetFixcomtype sets the fixcomtype edge to Fixcomtype.
func (fc *FixCreate) SetFixcomtype(f *Fixcomtype) *FixCreate {
	return fc.SetFixcomtypeID(f.ID)
}

// Mutation returns the FixMutation object of the builder.
func (fc *FixCreate) Mutation() *FixMutation {
	return fc.mutation
}

// Save creates the Fix in the database.
func (fc *FixCreate) Save(ctx context.Context) (*Fix, error) {
	if _, ok := fc.mutation.Productnumber(); !ok {
		return nil, &ValidationError{Name: "productnumber", err: errors.New("ent: missing required field \"productnumber\"")}
	}
	if v, ok := fc.mutation.Productnumber(); ok {
		if err := fix.ProductnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "productnumber", err: fmt.Errorf("ent: validator failed for field \"productnumber\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Problemtype(); !ok {
		return nil, &ValidationError{Name: "problemtype", err: errors.New("ent: missing required field \"problemtype\"")}
	}
	if v, ok := fc.mutation.Problemtype(); ok {
		if err := fix.ProblemtypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "problemtype", err: fmt.Errorf("ent: validator failed for field \"problemtype\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Queue(); !ok {
		return nil, &ValidationError{Name: "queue", err: errors.New("ent: missing required field \"queue\"")}
	}
	if v, ok := fc.mutation.Queue(); ok {
		if err := fix.QueueValidator(v); err != nil {
			return nil, &ValidationError{Name: "queue", err: fmt.Errorf("ent: validator failed for field \"queue\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Date(); !ok {
		return nil, &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	var (
		err  error
		node *Fix
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FixCreate) SaveX(ctx context.Context) *Fix {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FixCreate) sqlSave(ctx context.Context) (*Fix, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FixCreate) createSpec() (*Fix, *sqlgraph.CreateSpec) {
	var (
		f     = &Fix{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fix.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fix.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Productnumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fix.FieldProductnumber,
		})
		f.Productnumber = value
	}
	if value, ok := fc.mutation.Problemtype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fix.FieldProblemtype,
		})
		f.Problemtype = value
	}
	if value, ok := fc.mutation.Queue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fix.FieldQueue,
		})
		f.Queue = value
	}
	if value, ok := fc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fix.FieldDate,
		})
		f.Date = value
	}
	if nodes := fc.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   fix.FixTable,
			Columns: []string{fix.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminrepair.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FixbrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fix.FixbrandTable,
			Columns: []string{fix.FixbrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixbrand.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fix.PersonalTable,
			Columns: []string{fix.PersonalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fix.CustomerTable,
			Columns: []string{fix.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FixcomtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fix.FixcomtypeTable,
			Columns: []string{fix.FixcomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixcomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}
