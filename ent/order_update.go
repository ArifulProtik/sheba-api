// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ArifulProtik/sheba-api/ent/order"
	"github.com/ArifulProtik/sheba-api/ent/predicate"
	"github.com/ArifulProtik/sheba-api/ent/user"
	"github.com/google/uuid"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetServiceid sets the "serviceid" field.
func (ou *OrderUpdate) SetServiceid(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetServiceid(u)
	return ou
}

// SetProviderid sets the "providerid" field.
func (ou *OrderUpdate) SetProviderid(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetProviderid(u)
	return ou
}

// SetTotalcost sets the "totalcost" field.
func (ou *OrderUpdate) SetTotalcost(f float64) *OrderUpdate {
	ou.mutation.ResetTotalcost()
	ou.mutation.SetTotalcost(f)
	return ou
}

// AddTotalcost adds f to the "totalcost" field.
func (ou *OrderUpdate) AddTotalcost(f float64) *OrderUpdate {
	ou.mutation.AddTotalcost(f)
	return ou
}

// SetAddress sets the "address" field.
func (ou *OrderUpdate) SetAddress(s []string) *OrderUpdate {
	ou.mutation.SetAddress(s)
	return ou
}

// SetIsDeclined sets the "is_declined" field.
func (ou *OrderUpdate) SetIsDeclined(b bool) *OrderUpdate {
	ou.mutation.SetIsDeclined(b)
	return ou
}

// SetNillableIsDeclined sets the "is_declined" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableIsDeclined(b *bool) *OrderUpdate {
	if b != nil {
		ou.SetIsDeclined(*b)
	}
	return ou
}

// SetPaymentOk sets the "payment_ok" field.
func (ou *OrderUpdate) SetPaymentOk(b bool) *OrderUpdate {
	ou.mutation.SetPaymentOk(b)
	return ou
}

// SetNillablePaymentOk sets the "payment_ok" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePaymentOk(b *bool) *OrderUpdate {
	if b != nil {
		ou.SetPaymentOk(*b)
	}
	return ou
}

// SetIsAccepted sets the "is_accepted" field.
func (ou *OrderUpdate) SetIsAccepted(b bool) *OrderUpdate {
	ou.mutation.SetIsAccepted(b)
	return ou
}

// SetNillableIsAccepted sets the "is_accepted" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableIsAccepted(b *bool) *OrderUpdate {
	if b != nil {
		ou.SetIsAccepted(*b)
	}
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OrderUpdate) SetUserID(id uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if _, ok := ou.mutation.UserID(); ou.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.user"`)
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Serviceid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldServiceid,
		})
	}
	if value, ok := ou.mutation.Providerid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldProviderid,
		})
	}
	if value, ok := ou.mutation.Totalcost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalcost,
		})
	}
	if value, ok := ou.mutation.AddedTotalcost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalcost,
		})
	}
	if value, ok := ou.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldAddress,
		})
	}
	if value, ok := ou.mutation.IsDeclined(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldIsDeclined,
		})
	}
	if value, ok := ou.mutation.PaymentOk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldPaymentOk,
		})
	}
	if value, ok := ou.mutation.IsAccepted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldIsAccepted,
		})
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetServiceid sets the "serviceid" field.
func (ouo *OrderUpdateOne) SetServiceid(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetServiceid(u)
	return ouo
}

// SetProviderid sets the "providerid" field.
func (ouo *OrderUpdateOne) SetProviderid(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetProviderid(u)
	return ouo
}

// SetTotalcost sets the "totalcost" field.
func (ouo *OrderUpdateOne) SetTotalcost(f float64) *OrderUpdateOne {
	ouo.mutation.ResetTotalcost()
	ouo.mutation.SetTotalcost(f)
	return ouo
}

// AddTotalcost adds f to the "totalcost" field.
func (ouo *OrderUpdateOne) AddTotalcost(f float64) *OrderUpdateOne {
	ouo.mutation.AddTotalcost(f)
	return ouo
}

// SetAddress sets the "address" field.
func (ouo *OrderUpdateOne) SetAddress(s []string) *OrderUpdateOne {
	ouo.mutation.SetAddress(s)
	return ouo
}

// SetIsDeclined sets the "is_declined" field.
func (ouo *OrderUpdateOne) SetIsDeclined(b bool) *OrderUpdateOne {
	ouo.mutation.SetIsDeclined(b)
	return ouo
}

// SetNillableIsDeclined sets the "is_declined" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableIsDeclined(b *bool) *OrderUpdateOne {
	if b != nil {
		ouo.SetIsDeclined(*b)
	}
	return ouo
}

// SetPaymentOk sets the "payment_ok" field.
func (ouo *OrderUpdateOne) SetPaymentOk(b bool) *OrderUpdateOne {
	ouo.mutation.SetPaymentOk(b)
	return ouo
}

// SetNillablePaymentOk sets the "payment_ok" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePaymentOk(b *bool) *OrderUpdateOne {
	if b != nil {
		ouo.SetPaymentOk(*b)
	}
	return ouo
}

// SetIsAccepted sets the "is_accepted" field.
func (ouo *OrderUpdateOne) SetIsAccepted(b bool) *OrderUpdateOne {
	ouo.mutation.SetIsAccepted(b)
	return ouo
}

// SetNillableIsAccepted sets the "is_accepted" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableIsAccepted(b *bool) *OrderUpdateOne {
	if b != nil {
		ouo.SetIsAccepted(*b)
	}
	return ouo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OrderUpdateOne) SetUserID(id uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if _, ok := ouo.mutation.UserID(); ouo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Order.user"`)
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Serviceid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldServiceid,
		})
	}
	if value, ok := ouo.mutation.Providerid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldProviderid,
		})
	}
	if value, ok := ouo.mutation.Totalcost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalcost,
		})
	}
	if value, ok := ouo.mutation.AddedTotalcost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotalcost,
		})
	}
	if value, ok := ouo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldAddress,
		})
	}
	if value, ok := ouo.mutation.IsDeclined(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldIsDeclined,
		})
	}
	if value, ok := ouo.mutation.PaymentOk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldPaymentOk,
		})
	}
	if value, ok := ouo.mutation.IsAccepted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldIsAccepted,
		})
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
