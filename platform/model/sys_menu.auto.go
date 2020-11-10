// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysMenu defined 系统菜单
type SysMenu struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" xml:"id"`
	// 名称
	Name null.String `xorm:"varchar(36) notnull comment('名称') 'name'" json:"name" xml:"name"`
	// 编码
	Code null.String `xorm:"varchar(36) notnull comment('编码') 'code'" json:"code" xml:"code"`
	// 父菜单ID，一级菜单为null
	Parent null.String `xorm:"varchar(36) comment('父菜单ID，一级菜单为null') 'parent'" json:"parent" xml:"parent"`
	// 继承关系
	Inheritance null.String `xorm:"varchar(500) comment('继承关系') 'inheritance'" json:"inheritance" xml:"inheritance"`
	// 菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	URL null.String `xorm:"varchar(200) notnull comment('菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀&#43;目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)') 'url'" json:"url" xml:"url"`
	// 菜单组件
	Component null.String `xorm:"varchar(100) comment('菜单组件') 'component'" json:"component" xml:"component"`
	// 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
	Perms null.String `xorm:"varchar(500) comment('授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)') 'perms'" json:"perms" xml:"perms"`
	// 类型 0：目录 1：菜单 2：按钮
	Type null.Int `xorm:"notnull comment('类型 0：目录 1：菜单 2：按钮') 'type'" json:"type" xml:"type"`
	// 菜单图标
	Icon null.String `xorm:"varchar(50) comment('菜单图标') 'icon'" json:"icon" xml:"icon"`
	// 排序
	Order null.Int `xorm:"notnull comment('排序') 'order'" json:"order" xml:"order"`
	// 是否隐藏
	Hidden null.Int `xorm:"notnull comment('是否隐藏') 'hidden'" json:"hidden" xml:"hidden"`
	// 创建人
	CreateBy null.String `xorm:"varchar(36) notnull comment('创建人') 'create_by'" json:"create_by" xml:"create_by"`
	// 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" xml:"create_time"`
	// 最后更新人
	UpdateBy null.String `xorm:"varchar(36) notnull comment('最后更新人') 'update_by'" json:"update_by" xml:"update_by"`
	// 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" xml:"update_time"`
	// 删除标记
	DelFlag null.Int `xorm:"notnull comment('删除标记') 'del_flag'" json:"del_flag" xml:"del_flag"`
	// 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" xml:"remark"`
}

// TableName table name of defined SysMenu
func (m *SysMenu) TableName() string {
	return "sys_menu"
}
