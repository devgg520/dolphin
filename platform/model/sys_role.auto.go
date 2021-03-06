// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysRole defined 角色
type SysRole struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// 名称
	Name null.String `xorm:"varchar(200) notnull comment('名称') 'name'" json:"name" form:"name" xml:"name"`
	// 编码
	Code null.String `xorm:"varchar(36) notnull comment('编码') 'code'" json:"code" form:"code" xml:"code"`
	// 状态 0：禁用 1：正常
	Status null.Int `xorm:"notnull comment('状态 0：禁用 1：正常') 'status'" json:"status" form:"status" xml:"status"`
	// 角色app首页url
	AppIndex null.String `xorm:"varchar(500) comment('角色app首页url') 'app_index'" json:"app_index" form:"app_index" xml:"app_index"`
	// 角色进入后台首页组件
	AdminIndex null.String `xorm:"varchar(500) comment('角色进入后台首页组件') 'admin_index'" json:"admin_index" form:"admin_index" xml:"admin_index"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" form:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" form:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" form:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// TableName table name of defined SysRole
func (m *SysRole) TableName() string {
	return "sys_role"
}
