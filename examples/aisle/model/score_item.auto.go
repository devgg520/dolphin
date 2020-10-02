// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ScoreItem defined
type ScoreItem struct {
	//
	ScoreItemId null.Int `xorm:"int(11) pk notnull autoincr 'score_item_id'" json:"score_item_id" xml:"score_item_id"`
	//
	SiName null.String `xorm:"varchar(100) 'si_name'" json:"si_name" xml:"si_name"`
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

// TableName table name of defined ScoreItem
func (m *ScoreItem) TableName() string {
	return "score_item"
}