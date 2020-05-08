// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysDataPermissionDetail defined 数据权限表明细
type SysDataPermissionDetail struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk 'id'" json:"id" xml:"id"`
	// 数据权限表ID
	DataPermissionId null.String `xorm:"varchar(36) 'data_permission_id'" json:"data_permission_id" xml:"data_permission_id"`
	// 角色ID
	RoleId null.String `xorm:"varchar(36) 'role_id'" json:"role_id" xml:"role_id"`
	// 权限规则
	Rule null.String `xorm:"varchar(1000) notnull 'rule'" json:"rule" xml:"rule"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull 'create_by'" json:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull 'update_by'" json:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull 'del_flag'" json:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined SysDataPermissionDetail
func (m *SysDataPermissionDetail) TableName() string {
	return "sys_data_permission_detail"
}