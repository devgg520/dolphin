// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassScheduleTask defined
type ClassScheduleTask struct {
	//
	CSTId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_t_id'" json:"c_s_t_id" xml:"c_s_t_id"`
	//
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" xml:"cs_id"`
	//
	FeedContent null.String `xorm:"'feed_content'" json:"feed_content" xml:"feed_content"`
	//
	FeedRequier null.String `xorm:"'feed_requier'" json:"feed_requier" xml:"feed_requier"`
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
	CheckState null.Int `xorm:"int(11) default(54) 'check_state'" json:"check_state" xml:"check_state"`
	//
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" xml:"check_time"`
	//
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" xml:"check_user"`
	//
	CheckNoReason null.String `xorm:"varchar(5000) 'check_no_reason'" json:"check_no_reason" xml:"check_no_reason"`
	//
	AddUserType null.String `xorm:"varchar(1000) 'add_user_type'" json:"add_user_type" xml:"add_user_type"`
	//
	IfSendStu null.Int `xorm:"int(11) default(3) 'if_send_stu'" json:"if_send_stu" xml:"if_send_stu"`
	//
	IfSendTa null.Int `xorm:"int(11) default(3) 'if_send_ta'" json:"if_send_ta" xml:"if_send_ta"`
}

// TableName table name of defined ClassScheduleTask
func (m *ClassScheduleTask) TableName() string {
	return "class_schedule_task"
}