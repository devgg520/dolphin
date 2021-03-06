// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// DateTable defined
type DateTable struct {
	//
	DateTableId null.Int `xorm:"int(11) pk notnull autoincr 'date_table_id'" json:"date_table_id" xml:"date_table_id"`
	//
	DtDate null.Time `xorm:"datetime 'dt_date'" json:"dt_date" xml:"dt_date"`
	//
	DtDesc null.String `xorm:"varchar(1000) 'dt_desc'" json:"dt_desc" xml:"dt_desc"`
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
	Week null.Int `xorm:"int(11) 'week'" json:"week" xml:"week"`
	//
	Isholiday null.Int `xorm:"int(11) 'isholiday'" json:"isholiday" xml:"isholiday"`
}

// TableName table name of defined DateTable
func (m *DateTable) TableName() string {
	return "date_table"
}
