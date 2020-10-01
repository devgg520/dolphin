// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// StudentClassYpk31 defined
type StudentClassYpk31 struct {
	//
	SCYId null.Int `xorm:"int(11) pk notnull autoincr 's_c_y_id'" json:"s_c_y_id" xml:"s_c_y_id"`
	//
	PkYpk null.Int `xorm:"int(11) 'pk_ypk'" json:"pk_ypk" xml:"pk_ypk"`
	//
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" xml:"pk_stu"`
	//
	Kc null.Float `xorm:"float(11,2) 'kc'" json:"kc" xml:"kc"`
	//
	OnePrice decimal.Decimal `xorm:"decimal(11,2) 'one_price'" json:"one_price" xml:"one_price"`
	//
	AllPrice decimal.Decimal `xorm:"decimal(11,2) 'all_price'" json:"all_price" xml:"all_price"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" xml:"pk_sch"`
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
	ClassDate null.Time `xorm:"datetime 'class_date'" json:"class_date" xml:"class_date"`
	//
	PkSct null.Int `xorm:"int(11) 'pk_sct'" json:"pk_sct" xml:"pk_sct"`
	//
	PkCt null.Int `xorm:"int(11) 'pk_ct'" json:"pk_ct" xml:"pk_ct"`
}

// TableName table name of defined StudentClassYpk31
func (m *StudentClassYpk31) TableName() string {
	return "student_class_ypk31"
}
