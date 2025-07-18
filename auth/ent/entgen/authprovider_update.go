// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"auth/ent/entgen/authprovider"
	"auth/ent/entgen/authuser"
	"auth/ent/entgen/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthProviderUpdate is the builder for updating AuthProvider entities.
type AuthProviderUpdate struct {
	config
	hooks    []Hook
	mutation *AuthProviderMutation
}

// Where appends a list predicates to the AuthProviderUpdate builder.
func (apu *AuthProviderUpdate) Where(ps ...predicate.AuthProvider) *AuthProviderUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetUpdatedAt sets the "updated_at" field.
func (apu *AuthProviderUpdate) SetUpdatedAt(t time.Time) *AuthProviderUpdate {
	apu.mutation.SetUpdatedAt(t)
	return apu
}

// SetDeletedAt sets the "deleted_at" field.
func (apu *AuthProviderUpdate) SetDeletedAt(t time.Time) *AuthProviderUpdate {
	apu.mutation.SetDeletedAt(t)
	return apu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableDeletedAt(t *time.Time) *AuthProviderUpdate {
	if t != nil {
		apu.SetDeletedAt(*t)
	}
	return apu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apu *AuthProviderUpdate) ClearDeletedAt() *AuthProviderUpdate {
	apu.mutation.ClearDeletedAt()
	return apu
}

// SetCreatedBy sets the "created_by" field.
func (apu *AuthProviderUpdate) SetCreatedBy(s string) *AuthProviderUpdate {
	apu.mutation.SetCreatedBy(s)
	return apu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableCreatedBy(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetCreatedBy(*s)
	}
	return apu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (apu *AuthProviderUpdate) ClearCreatedBy() *AuthProviderUpdate {
	apu.mutation.ClearCreatedBy()
	return apu
}

// SetUpdatedBy sets the "updated_by" field.
func (apu *AuthProviderUpdate) SetUpdatedBy(s string) *AuthProviderUpdate {
	apu.mutation.SetUpdatedBy(s)
	return apu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableUpdatedBy(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetUpdatedBy(*s)
	}
	return apu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (apu *AuthProviderUpdate) ClearUpdatedBy() *AuthProviderUpdate {
	apu.mutation.ClearUpdatedBy()
	return apu
}

// SetDeletedBy sets the "deleted_by" field.
func (apu *AuthProviderUpdate) SetDeletedBy(s string) *AuthProviderUpdate {
	apu.mutation.SetDeletedBy(s)
	return apu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableDeletedBy(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetDeletedBy(*s)
	}
	return apu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (apu *AuthProviderUpdate) ClearDeletedBy() *AuthProviderUpdate {
	apu.mutation.ClearDeletedBy()
	return apu
}

// SetUserAgent sets the "user_agent" field.
func (apu *AuthProviderUpdate) SetUserAgent(s string) *AuthProviderUpdate {
	apu.mutation.SetUserAgent(s)
	return apu
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableUserAgent(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetUserAgent(*s)
	}
	return apu
}

// ClearUserAgent clears the value of the "user_agent" field.
func (apu *AuthProviderUpdate) ClearUserAgent() *AuthProviderUpdate {
	apu.mutation.ClearUserAgent()
	return apu
}

// SetIPAddress sets the "ip_address" field.
func (apu *AuthProviderUpdate) SetIPAddress(s string) *AuthProviderUpdate {
	apu.mutation.SetIPAddress(s)
	return apu
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableIPAddress(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetIPAddress(*s)
	}
	return apu
}

// ClearIPAddress clears the value of the "ip_address" field.
func (apu *AuthProviderUpdate) ClearIPAddress() *AuthProviderUpdate {
	apu.mutation.ClearIPAddress()
	return apu
}

// SetName sets the "name" field.
func (apu *AuthProviderUpdate) SetName(s string) *AuthProviderUpdate {
	apu.mutation.SetName(s)
	return apu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableName(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetName(*s)
	}
	return apu
}

// SetDisplayName sets the "display_name" field.
func (apu *AuthProviderUpdate) SetDisplayName(s string) *AuthProviderUpdate {
	apu.mutation.SetDisplayName(s)
	return apu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (apu *AuthProviderUpdate) SetNillableDisplayName(s *string) *AuthProviderUpdate {
	if s != nil {
		apu.SetDisplayName(*s)
	}
	return apu
}

// AddUserIDs adds the "users" edge to the AuthUser entity by IDs.
func (apu *AuthProviderUpdate) AddUserIDs(ids ...string) *AuthProviderUpdate {
	apu.mutation.AddUserIDs(ids...)
	return apu
}

// AddUsers adds the "users" edges to the AuthUser entity.
func (apu *AuthProviderUpdate) AddUsers(a ...*AuthUser) *AuthProviderUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.AddUserIDs(ids...)
}

// Mutation returns the AuthProviderMutation object of the builder.
func (apu *AuthProviderUpdate) Mutation() *AuthProviderMutation {
	return apu.mutation
}

// ClearUsers clears all "users" edges to the AuthUser entity.
func (apu *AuthProviderUpdate) ClearUsers() *AuthProviderUpdate {
	apu.mutation.ClearUsers()
	return apu
}

// RemoveUserIDs removes the "users" edge to AuthUser entities by IDs.
func (apu *AuthProviderUpdate) RemoveUserIDs(ids ...string) *AuthProviderUpdate {
	apu.mutation.RemoveUserIDs(ids...)
	return apu
}

// RemoveUsers removes "users" edges to AuthUser entities.
func (apu *AuthProviderUpdate) RemoveUsers(a ...*AuthUser) *AuthProviderUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AuthProviderUpdate) Save(ctx context.Context) (int, error) {
	apu.defaults()
	return withHooks(ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AuthProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AuthProviderUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AuthProviderUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apu *AuthProviderUpdate) defaults() {
	if _, ok := apu.mutation.UpdatedAt(); !ok {
		v := authprovider.UpdateDefaultUpdatedAt()
		apu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AuthProviderUpdate) check() error {
	if v, ok := apu.mutation.CreatedBy(); ok {
		if err := authprovider.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.created_by": %w`, err)}
		}
	}
	if v, ok := apu.mutation.UpdatedBy(); ok {
		if err := authprovider.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.updated_by": %w`, err)}
		}
	}
	if v, ok := apu.mutation.DeletedBy(); ok {
		if err := authprovider.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "deleted_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.deleted_by": %w`, err)}
		}
	}
	if v, ok := apu.mutation.Name(); ok {
		if err := authprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.name": %w`, err)}
		}
	}
	return nil
}

func (apu *AuthProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authprovider.Table, authprovider.Columns, sqlgraph.NewFieldSpec(authprovider.FieldID, field.TypeInt))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.UpdatedAt(); ok {
		_spec.SetField(authprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apu.mutation.DeletedAt(); ok {
		_spec.SetField(authprovider.FieldDeletedAt, field.TypeTime, value)
	}
	if apu.mutation.DeletedAtCleared() {
		_spec.ClearField(authprovider.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := apu.mutation.CreatedBy(); ok {
		_spec.SetField(authprovider.FieldCreatedBy, field.TypeString, value)
	}
	if apu.mutation.CreatedByCleared() {
		_spec.ClearField(authprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := apu.mutation.UpdatedBy(); ok {
		_spec.SetField(authprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if apu.mutation.UpdatedByCleared() {
		_spec.ClearField(authprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := apu.mutation.DeletedBy(); ok {
		_spec.SetField(authprovider.FieldDeletedBy, field.TypeString, value)
	}
	if apu.mutation.DeletedByCleared() {
		_spec.ClearField(authprovider.FieldDeletedBy, field.TypeString)
	}
	if value, ok := apu.mutation.UserAgent(); ok {
		_spec.SetField(authprovider.FieldUserAgent, field.TypeString, value)
	}
	if apu.mutation.UserAgentCleared() {
		_spec.ClearField(authprovider.FieldUserAgent, field.TypeString)
	}
	if value, ok := apu.mutation.IPAddress(); ok {
		_spec.SetField(authprovider.FieldIPAddress, field.TypeString, value)
	}
	if apu.mutation.IPAddressCleared() {
		_spec.ClearField(authprovider.FieldIPAddress, field.TypeString)
	}
	if value, ok := apu.mutation.Name(); ok {
		_spec.SetField(authprovider.FieldName, field.TypeString, value)
	}
	if value, ok := apu.mutation.DisplayName(); ok {
		_spec.SetField(authprovider.FieldDisplayName, field.TypeString, value)
	}
	if apu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !apu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AuthProviderUpdateOne is the builder for updating a single AuthProvider entity.
type AuthProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (apuo *AuthProviderUpdateOne) SetUpdatedAt(t time.Time) *AuthProviderUpdateOne {
	apuo.mutation.SetUpdatedAt(t)
	return apuo
}

// SetDeletedAt sets the "deleted_at" field.
func (apuo *AuthProviderUpdateOne) SetDeletedAt(t time.Time) *AuthProviderUpdateOne {
	apuo.mutation.SetDeletedAt(t)
	return apuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableDeletedAt(t *time.Time) *AuthProviderUpdateOne {
	if t != nil {
		apuo.SetDeletedAt(*t)
	}
	return apuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apuo *AuthProviderUpdateOne) ClearDeletedAt() *AuthProviderUpdateOne {
	apuo.mutation.ClearDeletedAt()
	return apuo
}

// SetCreatedBy sets the "created_by" field.
func (apuo *AuthProviderUpdateOne) SetCreatedBy(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetCreatedBy(s)
	return apuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableCreatedBy(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetCreatedBy(*s)
	}
	return apuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (apuo *AuthProviderUpdateOne) ClearCreatedBy() *AuthProviderUpdateOne {
	apuo.mutation.ClearCreatedBy()
	return apuo
}

// SetUpdatedBy sets the "updated_by" field.
func (apuo *AuthProviderUpdateOne) SetUpdatedBy(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetUpdatedBy(s)
	return apuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableUpdatedBy(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetUpdatedBy(*s)
	}
	return apuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (apuo *AuthProviderUpdateOne) ClearUpdatedBy() *AuthProviderUpdateOne {
	apuo.mutation.ClearUpdatedBy()
	return apuo
}

// SetDeletedBy sets the "deleted_by" field.
func (apuo *AuthProviderUpdateOne) SetDeletedBy(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetDeletedBy(s)
	return apuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableDeletedBy(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetDeletedBy(*s)
	}
	return apuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (apuo *AuthProviderUpdateOne) ClearDeletedBy() *AuthProviderUpdateOne {
	apuo.mutation.ClearDeletedBy()
	return apuo
}

// SetUserAgent sets the "user_agent" field.
func (apuo *AuthProviderUpdateOne) SetUserAgent(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetUserAgent(s)
	return apuo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableUserAgent(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetUserAgent(*s)
	}
	return apuo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (apuo *AuthProviderUpdateOne) ClearUserAgent() *AuthProviderUpdateOne {
	apuo.mutation.ClearUserAgent()
	return apuo
}

// SetIPAddress sets the "ip_address" field.
func (apuo *AuthProviderUpdateOne) SetIPAddress(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetIPAddress(s)
	return apuo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableIPAddress(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetIPAddress(*s)
	}
	return apuo
}

// ClearIPAddress clears the value of the "ip_address" field.
func (apuo *AuthProviderUpdateOne) ClearIPAddress() *AuthProviderUpdateOne {
	apuo.mutation.ClearIPAddress()
	return apuo
}

// SetName sets the "name" field.
func (apuo *AuthProviderUpdateOne) SetName(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetName(s)
	return apuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableName(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetName(*s)
	}
	return apuo
}

// SetDisplayName sets the "display_name" field.
func (apuo *AuthProviderUpdateOne) SetDisplayName(s string) *AuthProviderUpdateOne {
	apuo.mutation.SetDisplayName(s)
	return apuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (apuo *AuthProviderUpdateOne) SetNillableDisplayName(s *string) *AuthProviderUpdateOne {
	if s != nil {
		apuo.SetDisplayName(*s)
	}
	return apuo
}

// AddUserIDs adds the "users" edge to the AuthUser entity by IDs.
func (apuo *AuthProviderUpdateOne) AddUserIDs(ids ...string) *AuthProviderUpdateOne {
	apuo.mutation.AddUserIDs(ids...)
	return apuo
}

// AddUsers adds the "users" edges to the AuthUser entity.
func (apuo *AuthProviderUpdateOne) AddUsers(a ...*AuthUser) *AuthProviderUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.AddUserIDs(ids...)
}

// Mutation returns the AuthProviderMutation object of the builder.
func (apuo *AuthProviderUpdateOne) Mutation() *AuthProviderMutation {
	return apuo.mutation
}

// ClearUsers clears all "users" edges to the AuthUser entity.
func (apuo *AuthProviderUpdateOne) ClearUsers() *AuthProviderUpdateOne {
	apuo.mutation.ClearUsers()
	return apuo
}

// RemoveUserIDs removes the "users" edge to AuthUser entities by IDs.
func (apuo *AuthProviderUpdateOne) RemoveUserIDs(ids ...string) *AuthProviderUpdateOne {
	apuo.mutation.RemoveUserIDs(ids...)
	return apuo
}

// RemoveUsers removes "users" edges to AuthUser entities.
func (apuo *AuthProviderUpdateOne) RemoveUsers(a ...*AuthUser) *AuthProviderUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the AuthProviderUpdate builder.
func (apuo *AuthProviderUpdateOne) Where(ps ...predicate.AuthProvider) *AuthProviderUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AuthProviderUpdateOne) Select(field string, fields ...string) *AuthProviderUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AuthProvider entity.
func (apuo *AuthProviderUpdateOne) Save(ctx context.Context) (*AuthProvider, error) {
	apuo.defaults()
	return withHooks(ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AuthProviderUpdateOne) SaveX(ctx context.Context) *AuthProvider {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AuthProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AuthProviderUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apuo *AuthProviderUpdateOne) defaults() {
	if _, ok := apuo.mutation.UpdatedAt(); !ok {
		v := authprovider.UpdateDefaultUpdatedAt()
		apuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AuthProviderUpdateOne) check() error {
	if v, ok := apuo.mutation.CreatedBy(); ok {
		if err := authprovider.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.created_by": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.UpdatedBy(); ok {
		if err := authprovider.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.updated_by": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.DeletedBy(); ok {
		if err := authprovider.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "deleted_by", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.deleted_by": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.Name(); ok {
		if err := authprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entgen: validator failed for field "AuthProvider.name": %w`, err)}
		}
	}
	return nil
}

func (apuo *AuthProviderUpdateOne) sqlSave(ctx context.Context) (_node *AuthProvider, err error) {
	if err := apuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authprovider.Table, authprovider.Columns, sqlgraph.NewFieldSpec(authprovider.FieldID, field.TypeInt))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgen: missing "AuthProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authprovider.FieldID)
		for _, f := range fields {
			if !authprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != authprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.UpdatedAt(); ok {
		_spec.SetField(authprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apuo.mutation.DeletedAt(); ok {
		_spec.SetField(authprovider.FieldDeletedAt, field.TypeTime, value)
	}
	if apuo.mutation.DeletedAtCleared() {
		_spec.ClearField(authprovider.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := apuo.mutation.CreatedBy(); ok {
		_spec.SetField(authprovider.FieldCreatedBy, field.TypeString, value)
	}
	if apuo.mutation.CreatedByCleared() {
		_spec.ClearField(authprovider.FieldCreatedBy, field.TypeString)
	}
	if value, ok := apuo.mutation.UpdatedBy(); ok {
		_spec.SetField(authprovider.FieldUpdatedBy, field.TypeString, value)
	}
	if apuo.mutation.UpdatedByCleared() {
		_spec.ClearField(authprovider.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := apuo.mutation.DeletedBy(); ok {
		_spec.SetField(authprovider.FieldDeletedBy, field.TypeString, value)
	}
	if apuo.mutation.DeletedByCleared() {
		_spec.ClearField(authprovider.FieldDeletedBy, field.TypeString)
	}
	if value, ok := apuo.mutation.UserAgent(); ok {
		_spec.SetField(authprovider.FieldUserAgent, field.TypeString, value)
	}
	if apuo.mutation.UserAgentCleared() {
		_spec.ClearField(authprovider.FieldUserAgent, field.TypeString)
	}
	if value, ok := apuo.mutation.IPAddress(); ok {
		_spec.SetField(authprovider.FieldIPAddress, field.TypeString, value)
	}
	if apuo.mutation.IPAddressCleared() {
		_spec.ClearField(authprovider.FieldIPAddress, field.TypeString)
	}
	if value, ok := apuo.mutation.Name(); ok {
		_spec.SetField(authprovider.FieldName, field.TypeString, value)
	}
	if value, ok := apuo.mutation.DisplayName(); ok {
		_spec.SetField(authprovider.FieldDisplayName, field.TypeString, value)
	}
	if apuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !apuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   authprovider.UsersTable,
			Columns: []string{authprovider.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthProvider{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}
