// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// LabelInfo defined
type LabelInfo struct {
	//
	LabelInfoId null.Int `xorm:"int(11) pk notnull autoincr 'label_info_id'" json:"label_info_id" xml:"label_info_id"`
	//
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" xml:"parent_id"`
	//
	LabelName null.String `xorm:"varchar(100) 'label_name'" json:"label_name" xml:"label_name"`
	//
	LabelRemark null.String `xorm:"varchar(500) 'label_remark'" json:"label_remark" xml:"label_remark"`
	//
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" xml:"pt_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined LabelInfo
func (m *LabelInfo) TableName() string {
	return "label_info"
}
