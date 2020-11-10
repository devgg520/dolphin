// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysAreaTemplateDetail defined 区域信息模板明细
type SysAreaTemplateDetail struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" xml:"id"`
	// 区域名称
	Name null.String `xorm:"varchar(200) notnull comment('区域名称') 'name'" json:"name" xml:"name"`
	// 区域信息模板id
	TempId null.String `xorm:"varchar(36) comment('区域信息模板id') 'temp_id'" json:"temp_id" xml:"temp_id"`
	// 值
	Value null.String `xorm:"varchar(36) comment('值') 'value'" json:"value" xml:"value"`
	// 类型 0:数值项 1:单选项 2:文字项 3:列表项
	Type null.Int `xorm:"comment('类型 0:数值项 1:单选项 2:文字项 3:列表项') 'type'" json:"type" xml:"type"`
	// 优先级
	Priority null.Int `xorm:"comment('优先级') 'priority'" json:"priority" xml:"priority"`
	// 内容
	Content null.String `xorm:"text comment('内容') 'content'" json:"content" xml:"content"`
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

// TableName table name of defined SysAreaTemplateDetail
func (m *SysAreaTemplateDetail) TableName() string {
	return "sys_area_template_detail"
}
