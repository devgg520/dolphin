// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// FailStudentVisit defined
type FailStudentVisit struct {
	//
	FSVId null.Int `xorm:"int(11) pk notnull autoincr 'f_s_v_id'" json:"f_s_v_id" xml:"f_s_v_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	VisitContent null.String `xorm:"varchar(500) 'visit_content'" json:"visit_content" xml:"visit_content"`
	//
	VisitTime null.Time `xorm:"datetime 'visit_time'" json:"visit_time" xml:"visit_time"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined FailStudentVisit
func (m *FailStudentVisit) TableName() string {
	return "fail_student_visit"
}
