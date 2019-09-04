// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
Package job is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	Job
*/
package job

import context "context"

import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"

import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobORM struct {
	JobId   string `gorm:"primary_key"`
	JobName string `gorm:"type:varchar(512)"`
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
	for _, f := range updateMask.Paths {
		if f == prefix+"JobId" {
			patchee.JobId = patcher.JobId
			continue
		}
		if f == prefix+"JobName" {
			patchee.JobName = patcher.JobName
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
