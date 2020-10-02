// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Organ defined
type Organ struct {
	//
	OrganId null.Int `xorm:"int(11) pk notnull autoincr 'organ_id'" json:"organ_id" xml:"organ_id"`
	//
	OrganName null.String `xorm:"varchar(50) 'organ_name'" json:"organ_name" xml:"organ_name"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	OrganPinyin null.String `xorm:"varchar(100) 'organ_pinyin'" json:"organ_pinyin" xml:"organ_pinyin"`
	//
	OrganNumber null.String `xorm:"varchar(50) 'organ_number'" json:"organ_number" xml:"organ_number"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" xml:"parent_id"`
	//
	IfBuyOnline null.Int `xorm:"int(11) 'if_buy_online'" json:"if_buy_online" xml:"if_buy_online"`
	//
	OrganTell null.String `xorm:"varchar(20) 'organ_tell'" json:"organ_tell" xml:"organ_tell"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	ShengId null.Int `xorm:"int(11) 'sheng_id'" json:"sheng_id" xml:"sheng_id"`
}

// TableName table name of defined Organ
func (m *Organ) TableName() string {
	return "organ"
}