// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Materiel defined
type Materiel struct {
	//
	MaterielId null.Int `xorm:"int(11) pk notnull autoincr 'materiel_id'" json:"materiel_id" xml:"materiel_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	MaterielName null.String `xorm:"varchar(50) 'materiel_name'" json:"materiel_name" xml:"materiel_name"`
	//
	MaterielNum null.Int `xorm:"int(11) 'materiel_num'" json:"materiel_num" xml:"materiel_num"`
	//
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" xml:"organ_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" xml:"os_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined Materiel
func (m *Materiel) TableName() string {
	return "materiel"
}
