// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	Sallary
	Attachment
	Job
*/
package model

import context "context"
import strings "strings"
import time "time"

import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import pq1 "github.com/lib/pq"
import ptypes1 "github.com/golang/protobuf/ptypes"

import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SallaryORM struct {
	Currency           string
	MaxSallaryJobJobId *string
	MinSallaryJobJobId *string
	Value              uint64
}

// TableName overrides the default tablename generated by GORM
func (SallaryORM) TableName() string {
	return "sallaries"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Sallary) ToORM(ctx context.Context) (SallaryORM, error) {
	to := SallaryORM{}
	var err error
	if prehook, ok := interface{}(m).(SallaryWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Value = m.Value
	to.Currency = m.Currency
	if posthook, ok := interface{}(m).(SallaryWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SallaryORM) ToPB(ctx context.Context) (Sallary, error) {
	to := Sallary{}
	var err error
	if prehook, ok := interface{}(m).(SallaryWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Value = m.Value
	to.Currency = m.Currency
	if posthook, ok := interface{}(m).(SallaryWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Sallary the arg will be the target, the caller the one being converted from

// SallaryBeforeToORM called before default ToORM code
type SallaryWithBeforeToORM interface {
	BeforeToORM(context.Context, *SallaryORM) error
}

// SallaryAfterToORM called after default ToORM code
type SallaryWithAfterToORM interface {
	AfterToORM(context.Context, *SallaryORM) error
}

// SallaryBeforeToPB called before default ToPB code
type SallaryWithBeforeToPB interface {
	BeforeToPB(context.Context, *Sallary) error
}

// SallaryAfterToPB called after default ToPB code
type SallaryWithAfterToPB interface {
	AfterToPB(context.Context, *Sallary) error
}

type AttachmentORM struct {
	CreatedAt *time.Time
	JobJobId  *string
	Type      string
	Value     string
}

// TableName overrides the default tablename generated by GORM
func (AttachmentORM) TableName() string {
	return "attachments"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Attachment) ToORM(ctx context.Context) (AttachmentORM, error) {
	to := AttachmentORM{}
	var err error
	if prehook, ok := interface{}(m).(AttachmentWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Value = m.Value
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if posthook, ok := interface{}(m).(AttachmentWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AttachmentORM) ToPB(ctx context.Context) (Attachment, error) {
	to := Attachment{}
	var err error
	if prehook, ok := interface{}(m).(AttachmentWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Value = m.Value
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if posthook, ok := interface{}(m).(AttachmentWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Attachment the arg will be the target, the caller the one being converted from

// AttachmentBeforeToORM called before default ToORM code
type AttachmentWithBeforeToORM interface {
	BeforeToORM(context.Context, *AttachmentORM) error
}

// AttachmentAfterToORM called after default ToORM code
type AttachmentWithAfterToORM interface {
	AfterToORM(context.Context, *AttachmentORM) error
}

// AttachmentBeforeToPB called before default ToPB code
type AttachmentWithBeforeToPB interface {
	BeforeToPB(context.Context, *Attachment) error
}

// AttachmentAfterToPB called after default ToPB code
type AttachmentWithAfterToPB interface {
	AfterToPB(context.Context, *Attachment) error
}

type JobORM struct {
	CreatedAt      *time.Time
	JobAttachment  []*AttachmentORM `gorm:"foreignkey:JobJobId;association_foreignkey:JobId"`
	JobCategory    string
	JobDescription string `gorm:"type:varchar(512)"`
	JobId          string `gorm:"primary_key"`
	JobName        string `gorm:"type:varchar(512)"`
	JobStatus      int32
	JobTag         pq1.StringArray `gorm:"type:text[]"`
	JobType        int32
	Location       string
	MaxSallary     *SallaryORM     `gorm:"foreignkey:MaxSallaryJobJobId;association_foreignkey:JobId"`
	MinSallary     *SallaryORM     `gorm:"foreignkey:MinSallaryJobJobId;association_foreignkey:JobId"`
	SkillsRequired pq1.StringArray `gorm:"type:text[]"`
	UpdatedAt      *time.Time
	UsersApplied   pq1.StringArray `gorm:"type:text[]"`
	Views          int64
}

// TableName overrides the default tablename generated by GORM
func (JobORM) TableName() string {
	return "jobs"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Job) ToORM(ctx context.Context) (JobORM, error) {
	to := JobORM{}
	var err error
	if prehook, ok := interface{}(m).(JobWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.JobName = m.JobName
	to.JobDescription = m.JobDescription
	to.JobCategory = m.JobCategory
	to.Location = m.Location
	to.Views = m.Views
	if m.JobTag != nil {
		to.JobTag = make(pq1.StringArray, len(m.JobTag))
		copy(to.JobTag, m.JobTag)
	}
	if m.SkillsRequired != nil {
		to.SkillsRequired = make(pq1.StringArray, len(m.SkillsRequired))
		copy(to.SkillsRequired, m.SkillsRequired)
	}
	if m.UsersApplied != nil {
		to.UsersApplied = make(pq1.StringArray, len(m.UsersApplied))
		copy(to.UsersApplied, m.UsersApplied)
	}
	for _, v := range m.JobAttachment {
		if v != nil {
			if tempJobAttachment, cErr := v.ToORM(ctx); cErr == nil {
				to.JobAttachment = append(to.JobAttachment, &tempJobAttachment)
			} else {
				return to, cErr
			}
		} else {
			to.JobAttachment = append(to.JobAttachment, nil)
		}
	}
	to.JobType = int32(m.JobType)
	to.JobStatus = int32(m.JobStatus)
	if m.MinSallary != nil {
		tempMinSallary, err := m.MinSallary.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.MinSallary = &tempMinSallary
	}
	if m.MaxSallary != nil {
		tempMaxSallary, err := m.MaxSallary.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.MaxSallary = &tempMaxSallary
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(JobWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *JobORM) ToPB(ctx context.Context) (Job, error) {
	to := Job{}
	var err error
	if prehook, ok := interface{}(m).(JobWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.JobName = m.JobName
	to.JobDescription = m.JobDescription
	to.JobCategory = m.JobCategory
	to.Location = m.Location
	to.Views = m.Views
	if m.JobTag != nil {
		to.JobTag = make(pq1.StringArray, len(m.JobTag))
		copy(to.JobTag, m.JobTag)
	}
	if m.SkillsRequired != nil {
		to.SkillsRequired = make(pq1.StringArray, len(m.SkillsRequired))
		copy(to.SkillsRequired, m.SkillsRequired)
	}
	if m.UsersApplied != nil {
		to.UsersApplied = make(pq1.StringArray, len(m.UsersApplied))
		copy(to.UsersApplied, m.UsersApplied)
	}
	for _, v := range m.JobAttachment {
		if v != nil {
			if tempJobAttachment, cErr := v.ToPB(ctx); cErr == nil {
				to.JobAttachment = append(to.JobAttachment, &tempJobAttachment)
			} else {
				return to, cErr
			}
		} else {
			to.JobAttachment = append(to.JobAttachment, nil)
		}
	}
	to.JobType = Job_JobType(m.JobType)
	to.JobStatus = Job_JobStatus(m.JobStatus)
	if m.MinSallary != nil {
		tempMinSallary, err := m.MinSallary.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.MinSallary = &tempMinSallary
	}
	if m.MaxSallary != nil {
		tempMaxSallary, err := m.MaxSallary.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.MaxSallary = &tempMaxSallary
	}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if posthook, ok := interface{}(m).(JobWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Job the arg will be the target, the caller the one being converted from

// JobBeforeToORM called before default ToORM code
type JobWithBeforeToORM interface {
	BeforeToORM(context.Context, *JobORM) error
}

// JobAfterToORM called after default ToORM code
type JobWithAfterToORM interface {
	AfterToORM(context.Context, *JobORM) error
}

// JobBeforeToPB called before default ToPB code
type JobWithBeforeToPB interface {
	BeforeToPB(context.Context, *Job) error
}

// JobAfterToPB called after default ToPB code
type JobWithAfterToPB interface {
	AfterToPB(context.Context, *Job) error
}

// DefaultCreateSallary executes a basic gorm create call
func DefaultCreateSallary(ctx context.Context, in *Sallary, db *gorm1.DB) (*Sallary, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SallaryORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SallaryORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SallaryORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type SallaryORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskSallary patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSallary(ctx context.Context, patchee *Sallary, patcher *Sallary, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Sallary, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
		if f == prefix+"Currency" {
			patchee.Currency = patcher.Currency
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListSallary executes a gorm list call
func DefaultListSallary(ctx context.Context, db *gorm1.DB) ([]*Sallary, error) {
	in := Sallary{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SallaryORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &SallaryORM{}, &Sallary{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SallaryORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []SallaryORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SallaryORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Sallary{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SallaryORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type SallaryORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type SallaryORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]SallaryORM) error
}

// DefaultCreateAttachment executes a basic gorm create call
func DefaultCreateAttachment(ctx context.Context, in *Attachment, db *gorm1.DB) (*Attachment, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AttachmentORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AttachmentORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AttachmentORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AttachmentORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskAttachment patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAttachment(ctx context.Context, patchee *Attachment, patcher *Attachment, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Attachment, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAttachment executes a gorm list call
func DefaultListAttachment(ctx context.Context, db *gorm1.DB) ([]*Attachment, error) {
	in := Attachment{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AttachmentORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &AttachmentORM{}, &Attachment{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AttachmentORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []AttachmentORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AttachmentORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Attachment{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AttachmentORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AttachmentORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AttachmentORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]AttachmentORM) error
}

// DefaultCreateJob executes a basic gorm create call
func DefaultCreateJob(ctx context.Context, in *Job, db *gorm1.DB) (*Job, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JobORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JobORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type JobORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JobORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskJob patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskJob(ctx context.Context, patchee *Job, patcher *Job, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Job, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	var updatedMinSallary bool
	var updatedMaxSallary bool
	for i, f := range updateMask.Paths {
		if f == prefix+"JobId" {
			patchee.JobId = patcher.JobId
			continue
		}
		if f == prefix+"JobName" {
			patchee.JobName = patcher.JobName
			continue
		}
		if f == prefix+"JobDescription" {
			patchee.JobDescription = patcher.JobDescription
			continue
		}
		if f == prefix+"JobCategory" {
			patchee.JobCategory = patcher.JobCategory
			continue
		}
		if f == prefix+"Location" {
			patchee.Location = patcher.Location
			continue
		}
		if f == prefix+"Views" {
			patchee.Views = patcher.Views
			continue
		}
		if f == prefix+"JobTag" {
			patchee.JobTag = patcher.JobTag
			continue
		}
		if f == prefix+"SkillsRequired" {
			patchee.SkillsRequired = patcher.SkillsRequired
			continue
		}
		if f == prefix+"UsersApplied" {
			patchee.UsersApplied = patcher.UsersApplied
			continue
		}
		if f == prefix+"JobAttachment" {
			patchee.JobAttachment = patcher.JobAttachment
			continue
		}
		if f == prefix+"JobType" {
			patchee.JobType = patcher.JobType
			continue
		}
		if f == prefix+"JobStatus" {
			patchee.JobStatus = patcher.JobStatus
			continue
		}
		if !updatedMinSallary && strings.HasPrefix(f, prefix+"MinSallary.") {
			updatedMinSallary = true
			if patcher.MinSallary == nil {
				patchee.MinSallary = nil
				continue
			}
			if patchee.MinSallary == nil {
				patchee.MinSallary = &Sallary{}
			}
			if o, err := DefaultApplyFieldMaskSallary(ctx, patchee.MinSallary, patcher.MinSallary, &field_mask1.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"MinSallary.", db); err != nil {
				return nil, err
			} else {
				patchee.MinSallary = o
			}
			continue
		}
		if f == prefix+"MinSallary" {
			updatedMinSallary = true
			patchee.MinSallary = patcher.MinSallary
			continue
		}
		if !updatedMaxSallary && strings.HasPrefix(f, prefix+"MaxSallary.") {
			updatedMaxSallary = true
			if patcher.MaxSallary == nil {
				patchee.MaxSallary = nil
				continue
			}
			if patchee.MaxSallary == nil {
				patchee.MaxSallary = &Sallary{}
			}
			if o, err := DefaultApplyFieldMaskSallary(ctx, patchee.MaxSallary, patcher.MaxSallary, &field_mask1.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"MaxSallary.", db); err != nil {
				return nil, err
			} else {
				patchee.MaxSallary = o
			}
			continue
		}
		if f == prefix+"MaxSallary" {
			updatedMaxSallary = true
			patchee.MaxSallary = patcher.MaxSallary
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListJob executes a gorm list call
func DefaultListJob(ctx context.Context, db *gorm1.DB) ([]*Job, error) {
	in := Job{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JobORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &JobORM{}, &Job{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JobORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("job_id")
	ormResponse := []JobORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JobORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Job{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type JobORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JobORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JobORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]JobORM) error
}
