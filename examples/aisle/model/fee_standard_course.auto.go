// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// FeeStandardCourse defined
type FeeStandardCourse struct {
	//
	FSCId null.Int `xorm:"int(11) pk notnull autoincr 'f_s_c_id'" json:"f_s_c_id" xml:"f_s_c_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	FsId null.Int `xorm:"int(11) 'fs_id'" json:"fs_id" xml:"fs_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	CourseId null.Int `xorm:"int(11) 'course_id'" json:"course_id" xml:"course_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined FeeStandardCourse
func (m *FeeStandardCourse) TableName() string {
	return "fee_standard_course"
}