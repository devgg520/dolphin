// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// SchoolOfLimit defined
type SchoolOfLimit struct {
	//
	SOLId null.Int `xorm:"int(11) pk notnull autoincr 's_o_l_id'" json:"s_o_l_id" xml:"s_o_l_id"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" xml:"pk_sch"`
	//
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" xml:"start_date"`
	//
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" xml:"end_date"`
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
	OfYjType null.Int `xorm:"int(11) 'of_yj_type'" json:"of_yj_type" xml:"of_yj_type"`
	//
	OfBase decimal.Decimal `xorm:"decimal(50,3) 'of_base'" json:"of_base" xml:"of_base"`
	//
	SolName null.String `xorm:"varchar(20) 'sol_name'" json:"sol_name" xml:"sol_name"`
	//
	PkYyb null.Int `xorm:"int(11) 'pk_yyb'" json:"pk_yyb" xml:"pk_yyb"`
}

// TableName table name of defined SchoolOfLimit
func (m *SchoolOfLimit) TableName() string {
	return "school_of_limit"
}
