// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CsVisit defined
type CsVisit struct {
	//
	CsVisitId null.Int `xorm:"int(11) pk notnull autoincr 'cs_visit_id'" json:"cs_visit_id" xml:"cs_visit_id"`
	//
	VisitContent null.String `xorm:"varchar(300) 'visit_content'" json:"visit_content" xml:"visit_content"`
	//
	VisitStartTime null.Time `xorm:"datetime 'visit_start_time'" json:"visit_start_time" xml:"visit_start_time"`
	//
	VisitEndTime null.Time `xorm:"datetime 'visit_end_time'" json:"visit_end_time" xml:"visit_end_time"`
	//
	NextVisitTime null.Time `xorm:"datetime 'next_visit_time'" json:"next_visit_time" xml:"next_visit_time"`
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
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	VisituserId null.Int `xorm:"int(11) 'visituser_id'" json:"visituser_id" xml:"visituser_id"`
}

// TableName table name of defined CsVisit
func (m *CsVisit) TableName() string {
	return "cs_visit"
}
