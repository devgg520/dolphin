// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MenuPicture defined
type MenuPicture struct {
	//
	T2850 null.Int `xorm:"int(11) pk notnull autoincr 't_285_0'" json:"t_285_0" xml:"t_285_0"`
	//
	MenupicName null.String `xorm:"varchar(500) notnull 'menupic_name'" json:"menupic_name" xml:"menupic_name"`
	//
	MenupicPicture null.Int `xorm:"int(11) notnull 'menupic_picture'" json:"menupic_picture" xml:"menupic_picture"`
	//
	ProId null.Int `xorm:"int(11) 'pro_id'" json:"pro_id" xml:"pro_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined MenuPicture
func (m *MenuPicture) TableName() string {
	return "menu_picture"
}
