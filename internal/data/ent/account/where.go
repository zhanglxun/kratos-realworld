// Code generated by ent, DO NOT EDIT.

package account

import (
	"go-server/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// Account applies equality check predicate on the "account" field. It's identical to AccountEQ.
func Account(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccount, v))
}

// Pwd applies equality check predicate on the "pwd" field. It's identical to PwdEQ.
func Pwd(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPwd, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldMobile, v))
}

// CreateID applies equality check predicate on the "create_id" field. It's identical to CreateIDEQ.
func CreateID(v int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreateID, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreateTime, v))
}

// ModifyID applies equality check predicate on the "modify_id" field. It's identical to ModifyIDEQ.
func ModifyID(v int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldModifyID, v))
}

// ModifyTime applies equality check predicate on the "modify_time" field. It's identical to ModifyTimeEQ.
func ModifyTime(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldModifyTime, v))
}

// AccountEQ applies the EQ predicate on the "account" field.
func AccountEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccount, v))
}

// AccountNEQ applies the NEQ predicate on the "account" field.
func AccountNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAccount, v))
}

// AccountIn applies the In predicate on the "account" field.
func AccountIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldAccount, vs...))
}

// AccountNotIn applies the NotIn predicate on the "account" field.
func AccountNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldAccount, vs...))
}

// AccountGT applies the GT predicate on the "account" field.
func AccountGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldAccount, v))
}

// AccountGTE applies the GTE predicate on the "account" field.
func AccountGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldAccount, v))
}

// AccountLT applies the LT predicate on the "account" field.
func AccountLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldAccount, v))
}

// AccountLTE applies the LTE predicate on the "account" field.
func AccountLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldAccount, v))
}

// AccountContains applies the Contains predicate on the "account" field.
func AccountContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldAccount, v))
}

// AccountHasPrefix applies the HasPrefix predicate on the "account" field.
func AccountHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldAccount, v))
}

// AccountHasSuffix applies the HasSuffix predicate on the "account" field.
func AccountHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldAccount, v))
}

// AccountIsNil applies the IsNil predicate on the "account" field.
func AccountIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldAccount))
}

// AccountNotNil applies the NotNil predicate on the "account" field.
func AccountNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldAccount))
}

// AccountEqualFold applies the EqualFold predicate on the "account" field.
func AccountEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldAccount, v))
}

// AccountContainsFold applies the ContainsFold predicate on the "account" field.
func AccountContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldAccount, v))
}

// PwdEQ applies the EQ predicate on the "pwd" field.
func PwdEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPwd, v))
}

// PwdNEQ applies the NEQ predicate on the "pwd" field.
func PwdNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPwd, v))
}

// PwdIn applies the In predicate on the "pwd" field.
func PwdIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldPwd, vs...))
}

// PwdNotIn applies the NotIn predicate on the "pwd" field.
func PwdNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldPwd, vs...))
}

// PwdGT applies the GT predicate on the "pwd" field.
func PwdGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldPwd, v))
}

// PwdGTE applies the GTE predicate on the "pwd" field.
func PwdGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldPwd, v))
}

// PwdLT applies the LT predicate on the "pwd" field.
func PwdLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldPwd, v))
}

// PwdLTE applies the LTE predicate on the "pwd" field.
func PwdLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldPwd, v))
}

// PwdContains applies the Contains predicate on the "pwd" field.
func PwdContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldPwd, v))
}

// PwdHasPrefix applies the HasPrefix predicate on the "pwd" field.
func PwdHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldPwd, v))
}

// PwdHasSuffix applies the HasSuffix predicate on the "pwd" field.
func PwdHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldPwd, v))
}

// PwdIsNil applies the IsNil predicate on the "pwd" field.
func PwdIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldPwd))
}

// PwdNotNil applies the NotNil predicate on the "pwd" field.
func PwdNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldPwd))
}

// PwdEqualFold applies the EqualFold predicate on the "pwd" field.
func PwdEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldPwd, v))
}

// PwdContainsFold applies the ContainsFold predicate on the "pwd" field.
func PwdContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldPwd, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldNickname))
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldNickname))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldNickname, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldEmail, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldMobile, v))
}

// CreateIDEQ applies the EQ predicate on the "create_id" field.
func CreateIDEQ(v int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreateID, v))
}

// CreateIDNEQ applies the NEQ predicate on the "create_id" field.
func CreateIDNEQ(v int64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCreateID, v))
}

// CreateIDIn applies the In predicate on the "create_id" field.
func CreateIDIn(vs ...int64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCreateID, vs...))
}

// CreateIDNotIn applies the NotIn predicate on the "create_id" field.
func CreateIDNotIn(vs ...int64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCreateID, vs...))
}

// CreateIDGT applies the GT predicate on the "create_id" field.
func CreateIDGT(v int64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCreateID, v))
}

// CreateIDGTE applies the GTE predicate on the "create_id" field.
func CreateIDGTE(v int64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCreateID, v))
}

// CreateIDLT applies the LT predicate on the "create_id" field.
func CreateIDLT(v int64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCreateID, v))
}

// CreateIDLTE applies the LTE predicate on the "create_id" field.
func CreateIDLTE(v int64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCreateID, v))
}

// CreateIDIsNil applies the IsNil predicate on the "create_id" field.
func CreateIDIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldCreateID))
}

// CreateIDNotNil applies the NotNil predicate on the "create_id" field.
func CreateIDNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldCreateID))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldCreateTime))
}

// ModifyIDEQ applies the EQ predicate on the "modify_id" field.
func ModifyIDEQ(v int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldModifyID, v))
}

// ModifyIDNEQ applies the NEQ predicate on the "modify_id" field.
func ModifyIDNEQ(v int64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldModifyID, v))
}

// ModifyIDIn applies the In predicate on the "modify_id" field.
func ModifyIDIn(vs ...int64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldModifyID, vs...))
}

// ModifyIDNotIn applies the NotIn predicate on the "modify_id" field.
func ModifyIDNotIn(vs ...int64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldModifyID, vs...))
}

// ModifyIDGT applies the GT predicate on the "modify_id" field.
func ModifyIDGT(v int64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldModifyID, v))
}

// ModifyIDGTE applies the GTE predicate on the "modify_id" field.
func ModifyIDGTE(v int64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldModifyID, v))
}

// ModifyIDLT applies the LT predicate on the "modify_id" field.
func ModifyIDLT(v int64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldModifyID, v))
}

// ModifyIDLTE applies the LTE predicate on the "modify_id" field.
func ModifyIDLTE(v int64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldModifyID, v))
}

// ModifyIDIsNil applies the IsNil predicate on the "modify_id" field.
func ModifyIDIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldModifyID))
}

// ModifyIDNotNil applies the NotNil predicate on the "modify_id" field.
func ModifyIDNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldModifyID))
}

// ModifyTimeEQ applies the EQ predicate on the "modify_time" field.
func ModifyTimeEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldModifyTime, v))
}

// ModifyTimeNEQ applies the NEQ predicate on the "modify_time" field.
func ModifyTimeNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldModifyTime, v))
}

// ModifyTimeIn applies the In predicate on the "modify_time" field.
func ModifyTimeIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldModifyTime, vs...))
}

// ModifyTimeNotIn applies the NotIn predicate on the "modify_time" field.
func ModifyTimeNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldModifyTime, vs...))
}

// ModifyTimeGT applies the GT predicate on the "modify_time" field.
func ModifyTimeGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldModifyTime, v))
}

// ModifyTimeGTE applies the GTE predicate on the "modify_time" field.
func ModifyTimeGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldModifyTime, v))
}

// ModifyTimeLT applies the LT predicate on the "modify_time" field.
func ModifyTimeLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldModifyTime, v))
}

// ModifyTimeLTE applies the LTE predicate on the "modify_time" field.
func ModifyTimeLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldModifyTime, v))
}

// ModifyTimeIsNil applies the IsNil predicate on the "modify_time" field.
func ModifyTimeIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldModifyTime))
}

// ModifyTimeNotNil applies the NotNil predicate on the "modify_time" field.
func ModifyTimeNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldModifyTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
