// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// RegularDictionary defined
type RegularDictionary struct {
	//
	PkDicId null.Int `xorm:"int(11) pk notnull autoincr 'pk_dic_id'" json:"pk_dic_id" xml:"pk_dic_id"`
	//
	RegularDicId null.Float `xorm:"float(11,2) 'regular_dic_id'" json:"regular_dic_id" xml:"regular_dic_id"`
	//
	DicParent null.Float `xorm:"float(11,2) 'dic_parent'" json:"dic_parent" xml:"dic_parent"`
	//
	DicId null.Float `xorm:"float(11,2) 'dic_id'" json:"dic_id" xml:"dic_id"`
	//
	DicName null.String `xorm:"varchar(500) 'dic_name'" json:"dic_name" xml:"dic_name"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
}

// TableName table name of defined RegularDictionary
func (m *RegularDictionary) TableName() string {
	return "regular_dictionary"
}
