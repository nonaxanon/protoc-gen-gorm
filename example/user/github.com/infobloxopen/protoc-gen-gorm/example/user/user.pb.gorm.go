package user

import (
	context "context"
	fmt "fmt"
	ptypes "github.com/golang/protobuf/ptypes"
	auth "github.com/infobloxopen/atlas-app-toolkit/auth"
	gateway "github.com/infobloxopen/atlas-app-toolkit/gateway"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	resource "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	strings "strings"
	time "time"
)

type UserORM struct {
	Birthday   *time.Time
	Num        uint32
	AccountID  string
	CreditCard *CreditCardORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	Id         string         `gorm:"type:uuid;primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Decode(&User{}, m.Id); err != nil {
		return to, err
	} else if v != nil {
		to.Id = v.(string)
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.Birthday != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(m.Birthday); err != nil {
			return to, err
		}
		to.Birthday = &t
	}
	to.Num = m.Num
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Encode(&User{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.Birthday != nil {
		if to.Birthday, err = ptypes.TimestampProto(*m.Birthday); err != nil {
			return to, err
		}
	}
	to.Num = m.Num
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

type CreditCardORM struct {
	Number    string
	UserId    *string
	AccountID string
	Id        int64 `gorm:"type:integer;primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// TableName overrides the default tablename generated by GORM
func (CreditCardORM) TableName() string {
	return "credit_cards"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *CreditCard) ToORM(ctx context.Context) (CreditCardORM, error) {
	to := CreditCardORM{}
	var err error
	if prehook, ok := interface{}(m).(CreditCardWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.DecodeInt64(&CreditCard{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	to.Number = m.Number
	if m.UserId != nil {
		if v, err := resource.Decode(&User{}, m.UserId); err != nil {
			return to, err
		} else if v != nil {
			vv := v.(string)
			to.UserId = &vv
		}
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(CreditCardWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CreditCardORM) ToPB(ctx context.Context) (CreditCard, error) {
	to := CreditCard{}
	var err error
	if prehook, ok := interface{}(m).(CreditCardWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Encode(&CreditCard{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	if m.UserId != nil {
		if v, err := resource.Encode(&User{}, *m.UserId); err != nil {
			return to, err
		} else {
			to.UserId = v
		}
	}
	if posthook, ok := interface{}(m).(CreditCardWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type CreditCard the arg will be the target, the caller the one being converted from

// CreditCardBeforeToORM called before default ToORM code
type CreditCardWithBeforeToORM interface {
	BeforeToORM(context.Context, *CreditCardORM) error
}

// CreditCardAfterToORM called after default ToORM code
type CreditCardWithAfterToORM interface {
	AfterToORM(context.Context, *CreditCardORM) error
}

// CreditCardBeforeToPB called before default ToPB code
type CreditCardWithBeforeToPB interface {
	BeforeToPB(context.Context, *CreditCard) error
}

// CreditCardAfterToPB called after default ToPB code
type CreditCardWithAfterToPB interface {
	AfterToPB(context.Context, *CreditCard) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm.DB) error
}

// DefaultStrictUpdateUser clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	var count int64
	lockedRow := &UserORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	filterCreditCard := CreditCardORM{}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError.EmptyIdError
	}
	filterCreditCard.UserId = new(string)
	*filterCreditCard.UserId = ormObj.Id
	if err = db.Where(filterCreditCard).Delete(CreditCardORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask.FieldMask, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUser executes a bulk gorm update call with patch behavior
func DefaultPatchSetUser(ctx context.Context, objects []*User, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*User, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*User, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUser(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedId bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedBirthday bool
	var updatedCreditCard bool
	for i, f := range updateMask.Paths {
		if !updatedId && strings.HasPrefix(f, prefix+"Id.") {
			if patcher.Id == nil {
				patchee.Id = nil
				continue
			}
			if patchee.Id == nil {
				patchee.Id = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Id."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Id, patchee.Id, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Id" {
			updatedId = true
			patchee.Id = patcher.Id
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedBirthday && strings.HasPrefix(f, prefix+"Birthday.") {
			if patcher.Birthday == nil {
				patchee.Birthday = nil
				continue
			}
			if patchee.Birthday == nil {
				patchee.Birthday = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Birthday."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Birthday, patchee.Birthday, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Birthday" {
			updatedBirthday = true
			patchee.Birthday = patcher.Birthday
			continue
		}
		if f == prefix+"Age" {
			patchee.Age = patcher.Age
			continue
		}
		if f == prefix+"Num" {
			patchee.Num = patcher.Num
			continue
		}
		if !updatedCreditCard && strings.HasPrefix(f, prefix+"CreditCard.") {
			if patcher.CreditCard == nil {
				patchee.CreditCard = nil
				continue
			}
			if patchee.CreditCard == nil {
				patchee.CreditCard = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreditCard."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreditCard, patchee.CreditCard, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreditCard" {
			updatedCreditCard = true
			patchee.CreditCard = patcher.CreditCard
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm.DB) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserORM) error
}

// DefaultCreateCreditCard executes a basic gorm create call
func DefaultCreateCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CreditCardORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &CreditCardORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CreditCardORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CreditCardORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CreditCardORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CreditCardORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CreditCardORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteCreditCardSet(ctx context.Context, in []*CreditCard, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&CreditCardORM{})).(CreditCardORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&CreditCardORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CreditCardORM{})).(CreditCardORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CreditCardORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*CreditCard, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*CreditCard, *gorm.DB) error
}

// DefaultStrictUpdateCreditCard clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateCreditCard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	var count int64
	lockedRow := &CreditCardORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type CreditCardORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchCreditCard executes a basic gorm update call with patch behavior
func DefaultPatchCreditCard(ctx context.Context, in *CreditCard, updateMask *field_mask.FieldMask, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj CreditCard
	var err error
	if hook, ok := interface{}(&pbObj).(CreditCardWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadCreditCard(ctx, &CreditCard{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CreditCardWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskCreditCard(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CreditCardWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateCreditCard(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CreditCardWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CreditCardWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *CreditCard, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CreditCardWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *CreditCard, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CreditCardWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *CreditCard, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CreditCardWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *CreditCard, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetCreditCard executes a bulk gorm update call with patch behavior
func DefaultPatchSetCreditCard(ctx context.Context, objects []*CreditCard, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*CreditCard, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*CreditCard, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchCreditCard(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskCreditCard patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskCreditCard(ctx context.Context, patchee *CreditCard, patcher *CreditCard, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*CreditCard, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedId bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedUserId bool
	for i, f := range updateMask.Paths {
		if !updatedId && strings.HasPrefix(f, prefix+"Id.") {
			if patcher.Id == nil {
				patchee.Id = nil
				continue
			}
			if patchee.Id == nil {
				patchee.Id = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Id."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Id, patchee.Id, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Id" {
			updatedId = true
			patchee.Id = patcher.Id
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"Number" {
			patchee.Number = patcher.Number
			continue
		}
		if !updatedUserId && strings.HasPrefix(f, prefix+"UserId.") {
			if patcher.UserId == nil {
				patchee.UserId = nil
				continue
			}
			if patchee.UserId == nil {
				patchee.UserId = &message{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UserId."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UserId, patchee.UserId, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UserId" {
			updatedUserId = true
			patchee.UserId = patcher.UserId
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListCreditCard executes a gorm list call
func DefaultListCreditCard(ctx context.Context, db *gorm.DB) ([]*CreditCard, error) {
	in := CreditCard{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &CreditCardORM{}, &CreditCard{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CreditCardORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CreditCardORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*CreditCard{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CreditCardORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CreditCardORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]CreditCardORM) error
}
