// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketModelMateriel defined
type MarketModelMateriel struct {
	//
	MMMId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_m_id'" json:"m_m_m_id" xml:"m_m_m_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	MmId null.Int `xorm:"int(11) 'mm_id'" json:"mm_id" xml:"mm_id"`
	//
	PlanNum null.Int `xorm:"int(11) 'plan_num'" json:"plan_num" xml:"plan_num"`
	//
	ReceiveNum null.Int `xorm:"int(11) 'receive_num'" json:"receive_num" xml:"receive_num"`
	//
	UseNum null.Int `xorm:"int(11) 'use_num'" json:"use_num" xml:"use_num"`
	//
	ReturnNum null.Int `xorm:"int(11) 'return_num'" json:"return_num" xml:"return_num"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	MaterielId null.Int `xorm:"int(11) 'materiel_id'" json:"materiel_id" xml:"materiel_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined MarketModelMateriel
func (m *MarketModelMateriel) TableName() string {
	return "market_model_materiel"
}
