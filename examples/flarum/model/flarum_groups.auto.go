// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// FlarumGroups defined
type FlarumGroups struct {
	//
	ID null.Int `xorm:"int(10) pk notnull autoincr 'id'" json:"id" xml:"id"`
	//
	NameSingular null.String `xorm:"varchar(100) notnull 'name_singular'" json:"name_singular" xml:"name_singular"`
	//
	NamePlural null.String `xorm:"varchar(100) notnull 'name_plural'" json:"name_plural" xml:"name_plural"`
	//
	Color null.String `xorm:"varchar(20) 'color'" json:"color" xml:"color"`
	//
	Icon null.String `xorm:"varchar(100) 'icon'" json:"icon" xml:"icon"`
	// Creator
	CreateBy null.String `xorm:"varchar(36) comment('Creator') 'create_by'" json:"create_by" xml:"create_by"`
	// Creation time
	CreateTime null.Time `xorm:"datetime comment('Creation time') 'create_time'" json:"create_time" xml:"create_time"`
	// Last updated by
	UpdateBy null.String `xorm:"varchar(36) comment('Last updated by') 'update_by'" json:"update_by" xml:"update_by"`
	// Last update time
	UpdateTime null.Time `xorm:"datetime comment('Last update time') 'update_time'" json:"update_time" xml:"update_time"`
	// Delete tag
	DelFlag null.Int `xorm:"notnull comment('Delete tag') 'del_flag'" json:"del_flag" xml:"del_flag"`
	// Remark
	Remark null.String `xorm:"varchar(200) comment('Remark') 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined FlarumGroups
func (m *FlarumGroups) TableName() string {
	return "flarum_groups"
}
