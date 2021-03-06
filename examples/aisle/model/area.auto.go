// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Area defined
type Area struct {
	//
	AreaId null.Int `xorm:"int(11) pk notnull autoincr 'area_id'" json:"area_id" xml:"area_id"`
	//
	AreaName null.String `xorm:"varchar(20) 'area_name'" json:"area_name" xml:"area_name"`
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
	AreaCode null.String `xorm:"varchar(30) 'area_code'" json:"area_code" xml:"area_code"`
	//
	WcId null.Int `xorm:"int(11) 'wc_id'" json:"wc_id" xml:"wc_id"`
}

// TableName table name of defined Area
func (m *Area) TableName() string {
	return "area"
}
