// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Complaint defined
type Complaint struct {
	//
	ComplaintId null.Int `xorm:"int(11) pk notnull autoincr 'complaint_id'" json:"complaint_id" xml:"complaint_id"`
	//
	StudentId null.Int `xorm:"int(11) 'student_id'" json:"student_id" xml:"student_id"`
	//
	ComplaintReason null.String `xorm:"varchar(500) 'complaint_reason'" json:"complaint_reason" xml:"complaint_reason"`
	//
	ComplaintDate null.Time `xorm:"datetime 'complaint_date'" json:"complaint_date" xml:"complaint_date"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	IfSolve null.Int `xorm:"int(11) 'if_solve'" json:"if_solve" xml:"if_solve"`
	//
	Solution null.String `xorm:"varchar(500) 'solution'" json:"solution" xml:"solution"`
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
	SolveResult null.Int `xorm:"int(11) 'solve_result'" json:"solve_result" xml:"solve_result"`
	//
	SolutionDate null.Time `xorm:"datetime 'solution_date'" json:"solution_date" xml:"solution_date"`
	//
	ReceiveUser null.Int `xorm:"int(11) 'receive_user'" json:"receive_user" xml:"receive_user"`
	//
	ComState null.Int `xorm:"int(11) 'com_state'" json:"com_state" xml:"com_state"`
	//
	ComType null.Int `xorm:"int(11) 'com_type'" json:"com_type" xml:"com_type"`
	//
	ContentType null.Int `xorm:"int(11) 'content_type'" json:"content_type" xml:"content_type"`
	//
	ParentalFeedback null.String `xorm:"varchar(2000) 'parental_feedback'" json:"parental_feedback" xml:"parental_feedback"`
	//
	ComplaintSchool null.Int `xorm:"int(11) 'complaint_school'" json:"complaint_school" xml:"complaint_school"`
	//
	Complaint null.String `xorm:"varchar(500) 'complaint'" json:"complaint" xml:"complaint"`
	//
	ComplaintPhone null.String `xorm:"varchar(200) 'complaint_phone'" json:"complaint_phone" xml:"complaint_phone"`
	//
	ComplaintForm null.Int `xorm:"int(11) 'complaint_form'" json:"complaint_form" xml:"complaint_form"`
}

// TableName table name of defined Complaint
func (m *Complaint) TableName() string {
	return "complaint"
}
