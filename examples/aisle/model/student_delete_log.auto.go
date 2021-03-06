// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentDeleteLog defined
type StudentDeleteLog struct {
	//
	SDLId null.Int `xorm:"int(11) pk notnull autoincr 's_d_l_id'" json:"s_d_l_id" xml:"s_d_l_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	StuName null.String `xorm:"varchar(15) 'stu_name'" json:"stu_name" xml:"stu_name"`
	//
	LjStu null.Int `xorm:"int(11) 'lj_stu'" json:"lj_stu" xml:"lj_stu"`
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
	StuPhone null.String `xorm:"varchar(20) 'stu_phone'" json:"stu_phone" xml:"stu_phone"`
}

// TableName table name of defined StudentDeleteLog
func (m *StudentDeleteLog) TableName() string {
	return "student_delete_log"
}
