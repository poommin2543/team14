// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/receipt"
	"github.com/tanapon395/playlist-video/ent/title"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks      []Hook
	mutation   *CustomerMutation
	predicates []predicate.Customer
}

// Where adds a new predicate for the builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetCustomername sets the Customername field.
func (cu *CustomerUpdate) SetCustomername(s string) *CustomerUpdate {
	cu.mutation.SetCustomername(s)
	return cu
}

// SetAddress sets the Address field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetPhonenumber sets the Phonenumber field.
func (cu *CustomerUpdate) SetPhonenumber(s string) *CustomerUpdate {
	cu.mutation.SetPhonenumber(s)
	return cu
}

// SetIdentificationnumber sets the Identificationnumber field.
func (cu *CustomerUpdate) SetIdentificationnumber(s string) *CustomerUpdate {
	cu.mutation.SetIdentificationnumber(s)
	return cu
}

// SetGenderID sets the gender edge to Gender by id.
func (cu *CustomerUpdate) SetGenderID(id int) *CustomerUpdate {
	cu.mutation.SetGenderID(id)
	return cu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (cu *CustomerUpdate) SetNillableGenderID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetGenderID(*id)
	}
	return cu
}

// SetGender sets the gender edge to Gender.
func (cu *CustomerUpdate) SetGender(g *Gender) *CustomerUpdate {
	return cu.SetGenderID(g.ID)
}

// SetPersonalID sets the personal edge to Personal by id.
func (cu *CustomerUpdate) SetPersonalID(id int) *CustomerUpdate {
	cu.mutation.SetPersonalID(id)
	return cu
}

// SetNillablePersonalID sets the personal edge to Personal by id if the given value is not nil.
func (cu *CustomerUpdate) SetNillablePersonalID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetPersonalID(*id)
	}
	return cu
}

// SetPersonal sets the personal edge to Personal.
func (cu *CustomerUpdate) SetPersonal(p *Personal) *CustomerUpdate {
	return cu.SetPersonalID(p.ID)
}

// SetTitleID sets the title edge to Title by id.
func (cu *CustomerUpdate) SetTitleID(id int) *CustomerUpdate {
	cu.mutation.SetTitleID(id)
	return cu
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (cu *CustomerUpdate) SetNillableTitleID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetTitleID(*id)
	}
	return cu
}

// SetTitle sets the title edge to Title.
func (cu *CustomerUpdate) SetTitle(t *Title) *CustomerUpdate {
	return cu.SetTitleID(t.ID)
}

// AddFixIDs adds the fix edge to Fix by ids.
func (cu *CustomerUpdate) AddFixIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddFixIDs(ids...)
	return cu
}

// AddFix adds the fix edges to Fix.
func (cu *CustomerUpdate) AddFix(f ...*Fix) *CustomerUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.AddFixIDs(ids...)
}

// AddReceiptIDs adds the receipt edge to Receipt by ids.
func (cu *CustomerUpdate) AddReceiptIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddReceiptIDs(ids...)
	return cu
}

// AddReceipt adds the receipt edges to Receipt.
func (cu *CustomerUpdate) AddReceipt(r ...*Receipt) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReceiptIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearGender clears the gender edge to Gender.
func (cu *CustomerUpdate) ClearGender() *CustomerUpdate {
	cu.mutation.ClearGender()
	return cu
}

// ClearPersonal clears the personal edge to Personal.
func (cu *CustomerUpdate) ClearPersonal() *CustomerUpdate {
	cu.mutation.ClearPersonal()
	return cu
}

// ClearTitle clears the title edge to Title.
func (cu *CustomerUpdate) ClearTitle() *CustomerUpdate {
	cu.mutation.ClearTitle()
	return cu
}

// RemoveFixIDs removes the fix edge to Fix by ids.
func (cu *CustomerUpdate) RemoveFixIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveFixIDs(ids...)
	return cu
}

// RemoveFix removes fix edges to Fix.
func (cu *CustomerUpdate) RemoveFix(f ...*Fix) *CustomerUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.RemoveFixIDs(ids...)
}

// RemoveReceiptIDs removes the receipt edge to Receipt by ids.
func (cu *CustomerUpdate) RemoveReceiptIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveReceiptIDs(ids...)
	return cu
}

// RemoveReceipt removes receipt edges to Receipt.
func (cu *CustomerUpdate) RemoveReceipt(r ...*Receipt) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReceiptIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cu.mutation.Customername(); ok {
		if err := customer.CustomernameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Customername", err: fmt.Errorf("ent: validator failed for field \"Customername\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return 0, &ValidationError{Name: "Address", err: fmt.Errorf("ent: validator failed for field \"Address\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Phonenumber(); ok {
		if err := customer.PhonenumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "Phonenumber", err: fmt.Errorf("ent: validator failed for field \"Phonenumber\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Identificationnumber(); ok {
		if err := customer.IdentificationnumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "Identificationnumber", err: fmt.Errorf("ent: validator failed for field \"Identificationnumber\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Customername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCustomername,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cu.mutation.Phonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhonenumber,
		})
	}
	if value, ok := cu.mutation.Identificationnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldIdentificationnumber,
		})
	}
	if cu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.GenderTable,
			Columns: []string{customer.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.GenderTable,
			Columns: []string{customer.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PersonalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.PersonalTable,
			Columns: []string{customer.PersonalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.PersonalTable,
			Columns: []string{customer.PersonalColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.TitleTable,
			Columns: []string{customer.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.TitleTable,
			Columns: []string{customer.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cu.mutation.RemovedFixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FixTable,
			Columns: []string{customer.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FixTable,
			Columns: []string{customer.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cu.mutation.RemovedReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceiptTable,
			Columns: []string{customer.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: receipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceiptTable,
			Columns: []string{customer.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: receipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetCustomername sets the Customername field.
func (cuo *CustomerUpdateOne) SetCustomername(s string) *CustomerUpdateOne {
	cuo.mutation.SetCustomername(s)
	return cuo
}

// SetAddress sets the Address field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetPhonenumber sets the Phonenumber field.
func (cuo *CustomerUpdateOne) SetPhonenumber(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhonenumber(s)
	return cuo
}

// SetIdentificationnumber sets the Identificationnumber field.
func (cuo *CustomerUpdateOne) SetIdentificationnumber(s string) *CustomerUpdateOne {
	cuo.mutation.SetIdentificationnumber(s)
	return cuo
}

// SetGenderID sets the gender edge to Gender by id.
func (cuo *CustomerUpdateOne) SetGenderID(id int) *CustomerUpdateOne {
	cuo.mutation.SetGenderID(id)
	return cuo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableGenderID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetGenderID(*id)
	}
	return cuo
}

// SetGender sets the gender edge to Gender.
func (cuo *CustomerUpdateOne) SetGender(g *Gender) *CustomerUpdateOne {
	return cuo.SetGenderID(g.ID)
}

// SetPersonalID sets the personal edge to Personal by id.
func (cuo *CustomerUpdateOne) SetPersonalID(id int) *CustomerUpdateOne {
	cuo.mutation.SetPersonalID(id)
	return cuo
}

// SetNillablePersonalID sets the personal edge to Personal by id if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillablePersonalID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetPersonalID(*id)
	}
	return cuo
}

// SetPersonal sets the personal edge to Personal.
func (cuo *CustomerUpdateOne) SetPersonal(p *Personal) *CustomerUpdateOne {
	return cuo.SetPersonalID(p.ID)
}

// SetTitleID sets the title edge to Title by id.
func (cuo *CustomerUpdateOne) SetTitleID(id int) *CustomerUpdateOne {
	cuo.mutation.SetTitleID(id)
	return cuo
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableTitleID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetTitleID(*id)
	}
	return cuo
}

// SetTitle sets the title edge to Title.
func (cuo *CustomerUpdateOne) SetTitle(t *Title) *CustomerUpdateOne {
	return cuo.SetTitleID(t.ID)
}

// AddFixIDs adds the fix edge to Fix by ids.
func (cuo *CustomerUpdateOne) AddFixIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddFixIDs(ids...)
	return cuo
}

// AddFix adds the fix edges to Fix.
func (cuo *CustomerUpdateOne) AddFix(f ...*Fix) *CustomerUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.AddFixIDs(ids...)
}

// AddReceiptIDs adds the receipt edge to Receipt by ids.
func (cuo *CustomerUpdateOne) AddReceiptIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddReceiptIDs(ids...)
	return cuo
}

// AddReceipt adds the receipt edges to Receipt.
func (cuo *CustomerUpdateOne) AddReceipt(r ...*Receipt) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReceiptIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearGender clears the gender edge to Gender.
func (cuo *CustomerUpdateOne) ClearGender() *CustomerUpdateOne {
	cuo.mutation.ClearGender()
	return cuo
}

// ClearPersonal clears the personal edge to Personal.
func (cuo *CustomerUpdateOne) ClearPersonal() *CustomerUpdateOne {
	cuo.mutation.ClearPersonal()
	return cuo
}

// ClearTitle clears the title edge to Title.
func (cuo *CustomerUpdateOne) ClearTitle() *CustomerUpdateOne {
	cuo.mutation.ClearTitle()
	return cuo
}

// RemoveFixIDs removes the fix edge to Fix by ids.
func (cuo *CustomerUpdateOne) RemoveFixIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveFixIDs(ids...)
	return cuo
}

// RemoveFix removes fix edges to Fix.
func (cuo *CustomerUpdateOne) RemoveFix(f ...*Fix) *CustomerUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.RemoveFixIDs(ids...)
}

// RemoveReceiptIDs removes the receipt edge to Receipt by ids.
func (cuo *CustomerUpdateOne) RemoveReceiptIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveReceiptIDs(ids...)
	return cuo
}

// RemoveReceipt removes receipt edges to Receipt.
func (cuo *CustomerUpdateOne) RemoveReceipt(r ...*Receipt) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReceiptIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	if v, ok := cuo.mutation.Customername(); ok {
		if err := customer.CustomernameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Customername", err: fmt.Errorf("ent: validator failed for field \"Customername\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "Address", err: fmt.Errorf("ent: validator failed for field \"Address\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Phonenumber(); ok {
		if err := customer.PhonenumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Phonenumber", err: fmt.Errorf("ent: validator failed for field \"Phonenumber\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Identificationnumber(); ok {
		if err := customer.IdentificationnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Identificationnumber", err: fmt.Errorf("ent: validator failed for field \"Identificationnumber\": %w", err)}
		}
	}

	var (
		err  error
		node *Customer
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (c *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Customer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Customername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldCustomername,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.Phonenumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhonenumber,
		})
	}
	if value, ok := cuo.mutation.Identificationnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldIdentificationnumber,
		})
	}
	if cuo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.GenderTable,
			Columns: []string{customer.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.GenderTable,
			Columns: []string{customer.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PersonalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.PersonalTable,
			Columns: []string{customer.PersonalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.PersonalTable,
			Columns: []string{customer.PersonalColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.TitleTable,
			Columns: []string{customer.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.TitleTable,
			Columns: []string{customer.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cuo.mutation.RemovedFixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FixTable,
			Columns: []string{customer.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.FixTable,
			Columns: []string{customer.FixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cuo.mutation.RemovedReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceiptTable,
			Columns: []string{customer.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: receipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.ReceiptTable,
			Columns: []string{customer.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: receipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Customer{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
