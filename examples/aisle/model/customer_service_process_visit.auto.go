// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CustomerServiceProcessVisit defined
type CustomerServiceProcessVisit struct {
	//
	CSPVId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_p_v_id'" json:"c_s_p_v_id" xml:"c_s_p_v_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	VisitDate null.Time `xorm:"datetime 'visit_date'" json:"visit_date" xml:"visit_date"`
	//
	FinishDate null.Time `xorm:"datetime 'finish_date'" json:"finish_date" xml:"finish_date"`
	//
	FeedbackProblems null.String `xorm:"varchar(1000) 'feedback_problems'" json:"feedback_problems" xml:"feedback_problems"`
	//
	UnfinishedReason null.String `xorm:"varchar(1000) 'unfinished_reason'" json:"unfinished_reason" xml:"unfinished_reason"`
	//
	CssId null.Int `xorm:"int(11) 'css_id'" json:"css_id" xml:"css_id"`
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

// TableName table name of defined CustomerServiceProcessVisit
func (m *CustomerServiceProcessVisit) TableName() string {
	return "customer_service_process_visit"
}
