// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassTypeStageCourse defined
type ClassTypeStageCourse struct {
	//
	CTSCId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_s_c_id'" json:"c_t_s_c_id" xml:"c_t_s_c_id"`
	//
	CtsId null.Int `xorm:"int(11) 'cts_id'" json:"cts_id" xml:"cts_id"`
	//
	CourseId null.Int `xorm:"int(11) 'course_id'" json:"course_id" xml:"course_id"`
	//
	CtscHour null.Int `xorm:"int(11) 'ctsc_hour'" json:"ctsc_hour" xml:"ctsc_hour"`
	//
	CtscMinute null.Int `xorm:"int(11) 'ctsc_minute'" json:"ctsc_minute" xml:"ctsc_minute"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	Price null.Float `xorm:"float(10,2) 'price'" json:"price" xml:"price"`
	//
	AllPrice null.Float `xorm:"float(10,2) 'all_price'" json:"all_price" xml:"all_price"`
	//
	CtId null.Int `xorm:"int(11) 'ct_id'" json:"ct_id" xml:"ct_id"`
}

// TableName table name of defined ClassTypeStageCourse
func (m *ClassTypeStageCourse) TableName() string {
	return "class_type_stage_course"
}
