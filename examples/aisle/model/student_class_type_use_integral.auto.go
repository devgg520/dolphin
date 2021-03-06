// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentClassTypeUseIntegral defined
type StudentClassTypeUseIntegral struct {
	//
	SCTUIId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_u_i_id'" json:"s_c_t_u_i_id" xml:"s_c_t_u_i_id"`
	//
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" xml:"sct_id"`
	//
	UseIntegral null.Float `xorm:"float(11,2) 'use_integral'" json:"use_integral" xml:"use_integral"`
	//
	DyMoney null.Float `xorm:"float(11,2) 'dy_money'" json:"dy_money" xml:"dy_money"`
	//
	StuIntegral null.Int `xorm:"int(11) 'stu_integral'" json:"stu_integral" xml:"stu_integral"`
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

// TableName table name of defined StudentClassTypeUseIntegral
func (m *StudentClassTypeUseIntegral) TableName() string {
	return "student_class_type_use_integral"
}
