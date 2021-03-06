// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// LanguagePackCity defined
type LanguagePackCity struct {
	//
	LPCId null.Int `xorm:"int(11) pk notnull autoincr 'l_p_c_id'" json:"l_p_c_id" xml:"l_p_c_id"`
	//
	LpId null.Int `xorm:"int(11) 'lp_id'" json:"lp_id" xml:"lp_id"`
	//
	CityId null.Int `xorm:"int(11) 'city_id'" json:"city_id" xml:"city_id"`
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

// TableName table name of defined LanguagePackCity
func (m *LanguagePackCity) TableName() string {
	return "language_pack_city"
}
