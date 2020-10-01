// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// FamilyShareNumbe defined
type FamilyShareNumbe struct {
	//
	FSNId null.Int `xorm:"int(11) pk notnull autoincr 'f_s_n_id'" json:"f_s_n_id" xml:"f_s_n_id"`
	//
	FsNum null.String `xorm:"varchar(500) 'fs_num'" json:"fs_num" xml:"fs_num"`
	//
	FsYbNum null.Float `xorm:"float(10,2) 'fs_yb_num'" json:"fs_yb_num" xml:"fs_yb_num"`
	//
	FsYbHour null.Float `xorm:"float(10,2) 'fs_yb_hour'" json:"fs_yb_hour" xml:"fs_yb_hour"`
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

// TableName table name of defined FamilyShareNumbe
func (m *FamilyShareNumbe) TableName() string {
	return "family_share_numbe"
}
