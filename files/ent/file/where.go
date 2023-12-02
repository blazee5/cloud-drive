// Code generated by ent, DO NOT EDIT.

package file

import (
	"entgo.io/ent/dialect/sql"
	"github.com/blazee5/cloud-drive/files/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUserID, v))
}

// ContentType applies equality check predicate on the "content_type" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldContentType, v))
}

// DownloadCount applies equality check predicate on the "download_count" field. It's identical to DownloadCountEQ.
func DownloadCount(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDownloadCount, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldName, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldUserID, v))
}

// ContentTypeEQ applies the EQ predicate on the "content_type" field.
func ContentTypeEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldContentType, v))
}

// ContentTypeNEQ applies the NEQ predicate on the "content_type" field.
func ContentTypeNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldContentType, v))
}

// ContentTypeIn applies the In predicate on the "content_type" field.
func ContentTypeIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldContentType, vs...))
}

// ContentTypeNotIn applies the NotIn predicate on the "content_type" field.
func ContentTypeNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldContentType, vs...))
}

// ContentTypeGT applies the GT predicate on the "content_type" field.
func ContentTypeGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldContentType, v))
}

// ContentTypeGTE applies the GTE predicate on the "content_type" field.
func ContentTypeGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldContentType, v))
}

// ContentTypeLT applies the LT predicate on the "content_type" field.
func ContentTypeLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldContentType, v))
}

// ContentTypeLTE applies the LTE predicate on the "content_type" field.
func ContentTypeLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldContentType, v))
}

// ContentTypeContains applies the Contains predicate on the "content_type" field.
func ContentTypeContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldContentType, v))
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "content_type" field.
func ContentTypeHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldContentType, v))
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "content_type" field.
func ContentTypeHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldContentType, v))
}

// ContentTypeEqualFold applies the EqualFold predicate on the "content_type" field.
func ContentTypeEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldContentType, v))
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "content_type" field.
func ContentTypeContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldContentType, v))
}

// DownloadCountEQ applies the EQ predicate on the "download_count" field.
func DownloadCountEQ(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDownloadCount, v))
}

// DownloadCountNEQ applies the NEQ predicate on the "download_count" field.
func DownloadCountNEQ(v int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldDownloadCount, v))
}

// DownloadCountIn applies the In predicate on the "download_count" field.
func DownloadCountIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldDownloadCount, vs...))
}

// DownloadCountNotIn applies the NotIn predicate on the "download_count" field.
func DownloadCountNotIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldDownloadCount, vs...))
}

// DownloadCountGT applies the GT predicate on the "download_count" field.
func DownloadCountGT(v int) predicate.File {
	return predicate.File(sql.FieldGT(FieldDownloadCount, v))
}

// DownloadCountGTE applies the GTE predicate on the "download_count" field.
func DownloadCountGTE(v int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldDownloadCount, v))
}

// DownloadCountLT applies the LT predicate on the "download_count" field.
func DownloadCountLT(v int) predicate.File {
	return predicate.File(sql.FieldLT(FieldDownloadCount, v))
}

// DownloadCountLTE applies the LTE predicate on the "download_count" field.
func DownloadCountLTE(v int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldDownloadCount, v))
}

// DownloadCountIsNil applies the IsNil predicate on the "download_count" field.
func DownloadCountIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldDownloadCount))
}

// DownloadCountNotNil applies the NotNil predicate on the "download_count" field.
func DownloadCountNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldDownloadCount))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(sql.NotPredicates(p))
}
