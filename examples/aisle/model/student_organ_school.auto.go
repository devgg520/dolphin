// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentOrganSchool defined
type StudentOrganSchool struct {
	//
	SOSId null.Int `xorm:"int(11) pk notnull autoincr 's_o_s_id'" json:"s_o_s_id" xml:"s_o_s_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	StudentId null.Int `xorm:"int(11) 'student_id'" json:"student_id" xml:"student_id"`
	//
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" xml:"os_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	SyId null.Int `xorm:"int(11) 'sy_id'" json:"sy_id" xml:"sy_id"`
	//
	StuGxYy null.Int `xorm:"int(11) 'stu_gx_yy'" json:"stu_gx_yy" xml:"stu_gx_yy"`
	//
	SfPlzx null.Int `xorm:"int(11) 'sf_plzx'" json:"sf_plzx" xml:"sf_plzx"`
}

// TableName table name of defined StudentOrganSchool
func (m *StudentOrganSchool) TableName() string {
	return "student_organ_school"
}
