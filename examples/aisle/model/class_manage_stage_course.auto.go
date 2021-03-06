// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassManageStageCourse defined
type ClassManageStageCourse struct {
	//
	CMSCId null.Int `xorm:"int(11) pk notnull autoincr 'c_m_s_c_id'" json:"c_m_s_c_id" xml:"c_m_s_c_id"`
	//
	CmsId null.Int `xorm:"int(11) 'cms_id'" json:"cms_id" xml:"cms_id"`
	//
	CtscId null.Int `xorm:"int(11) 'ctsc_id'" json:"ctsc_id" xml:"ctsc_id"`
	//
	CourseId null.Int `xorm:"int(11) 'course_id'" json:"course_id" xml:"course_id"`
	//
	CmscHour null.Int `xorm:"int(11) 'cmsc_hour'" json:"cmsc_hour" xml:"cmsc_hour"`
	//
	CmscMinute null.Int `xorm:"int(11) 'cmsc_minute'" json:"cmsc_minute" xml:"cmsc_minute"`
	//
	PlanHour null.Int `xorm:"int(11) 'plan_hour'" json:"plan_hour" xml:"plan_hour"`
	//
	QrHour null.Int `xorm:"int(11) 'qr_hour'" json:"qr_hour" xml:"qr_hour"`
	//
	UseHour null.Int `xorm:"int(11) 'use_hour'" json:"use_hour" xml:"use_hour"`
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
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" xml:"class_id"`
}

// TableName table name of defined ClassManageStageCourse
func (m *ClassManageStageCourse) TableName() string {
	return "class_manage_stage_course"
}
