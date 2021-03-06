// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ButtonPicture defined
type ButtonPicture struct {
	//
	T2230 null.Int `xorm:"int(11) pk notnull autoincr 't_223_0'" json:"t_223_0" xml:"t_223_0"`
	//
	BtnpicName null.String `xorm:"varchar(500) notnull 'btnpic_name'" json:"btnpic_name" xml:"btnpic_name"`
	//
	BtnpicType null.Float `xorm:"float(11,2) notnull 'btnpic_type'" json:"btnpic_type" xml:"btnpic_type"`
	//
	BtnpicPicture null.Int `xorm:"int(11) notnull 'btnpic_picture'" json:"btnpic_picture" xml:"btnpic_picture"`
	//
	ProId null.Int `xorm:"int(11) 'pro_id'" json:"pro_id" xml:"pro_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined ButtonPicture
func (m *ButtonPicture) TableName() string {
	return "button_picture"
}
