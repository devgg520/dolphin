// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Partner defined
type Partner struct {
	//
	PartnerId null.Int `xorm:"int(11) pk notnull autoincr 'partner_id'" json:"partner_id" xml:"partner_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	PrName null.String `xorm:"varchar(100) 'pr_name'" json:"pr_name" xml:"pr_name"`
	//
	PrRemark null.String `xorm:"varchar(500) 'pr_remark'" json:"pr_remark" xml:"pr_remark"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	PartnerType null.Int `xorm:"int(11) 'partner_type'" json:"partner_type" xml:"partner_type"`
}

// TableName table name of defined Partner
func (m *Partner) TableName() string {
	return "partner"
}
