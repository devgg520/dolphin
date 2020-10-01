// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Level defined
type Level struct {
	//
	LevelId null.Int `xorm:"int(11) pk notnull autoincr 'level_id'" json:"level_id" xml:"level_id"`
	//
	LevelName null.String `xorm:"varchar(1000) 'level_name'" json:"level_name" xml:"level_name"`
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
	LevelDesc null.String `xorm:"varchar(1000) 'level_desc'" json:"level_desc" xml:"level_desc"`
}

// TableName table name of defined Level
func (m *Level) TableName() string {
	return "level"
}
