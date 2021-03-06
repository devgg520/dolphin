// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentCastingTable defined
type StudentCastingTable struct {
	//
	SCTId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_id'" json:"s_c_t_id" xml:"s_c_t_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	StuCastId null.Int `xorm:"int(11) 'stu_cast_id'" json:"stu_cast_id" xml:"stu_cast_id"`
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
}

// TableName table name of defined StudentCastingTable
func (m *StudentCastingTable) TableName() string {
	return "student_casting_table"
}
