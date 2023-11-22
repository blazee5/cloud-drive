// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blazee5/cloud-drive/microservices/files/ent/file"
	"github.com/blazee5/cloud-drive/microservices/files/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFile = "File"
)

// FileMutation represents an operation that mutates the File nodes in the graph.
type FileMutation struct {
	config
	op                Op
	typ               string
	id                *int
	name              *string
	user_id           *string
	download_count    *int
	adddownload_count *int
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*File, error)
	predicates        []predicate.File
}

var _ ent.Mutation = (*FileMutation)(nil)

// fileOption allows management of the mutation configuration using functional options.
type fileOption func(*FileMutation)

// newFileMutation creates new mutation for the File entity.
func newFileMutation(c config, op Op, opts ...fileOption) *FileMutation {
	m := &FileMutation{
		config:        c,
		op:            op,
		typ:           TypeFile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFileID sets the ID field of the mutation.
func withFileID(id int) fileOption {
	return func(m *FileMutation) {
		var (
			err   error
			once  sync.Once
			value *File
		)
		m.oldValue = func(ctx context.Context) (*File, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().File.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFile sets the old File of the mutation.
func withFile(node *File) fileOption {
	return func(m *FileMutation) {
		m.oldValue = func(context.Context) (*File, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FileMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FileMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().File.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *FileMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FileMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FileMutation) ResetName() {
	m.name = nil
}

// SetUserID sets the "user_id" field.
func (m *FileMutation) SetUserID(s string) {
	m.user_id = &s
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *FileMutation) UserID() (r string, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldUserID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *FileMutation) ResetUserID() {
	m.user_id = nil
}

// SetDownloadCount sets the "download_count" field.
func (m *FileMutation) SetDownloadCount(i int) {
	m.download_count = &i
	m.adddownload_count = nil
}

// DownloadCount returns the value of the "download_count" field in the mutation.
func (m *FileMutation) DownloadCount() (r int, exists bool) {
	v := m.download_count
	if v == nil {
		return
	}
	return *v, true
}

// OldDownloadCount returns the old "download_count" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldDownloadCount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDownloadCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDownloadCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDownloadCount: %w", err)
	}
	return oldValue.DownloadCount, nil
}

// AddDownloadCount adds i to the "download_count" field.
func (m *FileMutation) AddDownloadCount(i int) {
	if m.adddownload_count != nil {
		*m.adddownload_count += i
	} else {
		m.adddownload_count = &i
	}
}

// AddedDownloadCount returns the value that was added to the "download_count" field in this mutation.
func (m *FileMutation) AddedDownloadCount() (r int, exists bool) {
	v := m.adddownload_count
	if v == nil {
		return
	}
	return *v, true
}

// ClearDownloadCount clears the value of the "download_count" field.
func (m *FileMutation) ClearDownloadCount() {
	m.download_count = nil
	m.adddownload_count = nil
	m.clearedFields[file.FieldDownloadCount] = struct{}{}
}

// DownloadCountCleared returns if the "download_count" field was cleared in this mutation.
func (m *FileMutation) DownloadCountCleared() bool {
	_, ok := m.clearedFields[file.FieldDownloadCount]
	return ok
}

// ResetDownloadCount resets all changes to the "download_count" field.
func (m *FileMutation) ResetDownloadCount() {
	m.download_count = nil
	m.adddownload_count = nil
	delete(m.clearedFields, file.FieldDownloadCount)
}

// Where appends a list predicates to the FileMutation builder.
func (m *FileMutation) Where(ps ...predicate.File) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the FileMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *FileMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.File, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *FileMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *FileMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (File).
func (m *FileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FileMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, file.FieldName)
	}
	if m.user_id != nil {
		fields = append(fields, file.FieldUserID)
	}
	if m.download_count != nil {
		fields = append(fields, file.FieldDownloadCount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case file.FieldName:
		return m.Name()
	case file.FieldUserID:
		return m.UserID()
	case file.FieldDownloadCount:
		return m.DownloadCount()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case file.FieldName:
		return m.OldName(ctx)
	case file.FieldUserID:
		return m.OldUserID(ctx)
	case file.FieldDownloadCount:
		return m.OldDownloadCount(ctx)
	}
	return nil, fmt.Errorf("unknown File field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case file.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case file.FieldUserID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case file.FieldDownloadCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDownloadCount(v)
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FileMutation) AddedFields() []string {
	var fields []string
	if m.adddownload_count != nil {
		fields = append(fields, file.FieldDownloadCount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FileMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case file.FieldDownloadCount:
		return m.AddedDownloadCount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) AddField(name string, value ent.Value) error {
	switch name {
	case file.FieldDownloadCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDownloadCount(v)
		return nil
	}
	return fmt.Errorf("unknown File numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FileMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(file.FieldDownloadCount) {
		fields = append(fields, file.FieldDownloadCount)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FileMutation) ClearField(name string) error {
	switch name {
	case file.FieldDownloadCount:
		m.ClearDownloadCount()
		return nil
	}
	return fmt.Errorf("unknown File nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FileMutation) ResetField(name string) error {
	switch name {
	case file.FieldName:
		m.ResetName()
		return nil
	case file.FieldUserID:
		m.ResetUserID()
		return nil
	case file.FieldDownloadCount:
		m.ResetDownloadCount()
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FileMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FileMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FileMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FileMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FileMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown File unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FileMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown File edge %s", name)
}
