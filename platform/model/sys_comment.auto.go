// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SysComment defined 评论表
type SysComment struct {
	// 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// 主题ID
	TopicId null.String `xorm:"varchar(36) notnull comment('主题ID') 'topic_id'" json:"topic_id" form:"topic_id" xml:"topic_id"`
	// 主题类型
	TopicType null.String `xorm:"varchar(36) notnull comment('主题类型') 'topic_type'" json:"topic_type" form:"topic_type" xml:"topic_type"`
	// 评论内容
	Content null.String `xorm:"varchar(36) comment('评论内容') 'content'" json:"content" form:"content" xml:"content"`
	// 评论用户id
	FromUid null.String `xorm:"varchar(36) notnull comment('评论用户id') 'from_uid'" json:"from_uid" form:"from_uid" xml:"from_uid"`
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

// TableName table name of defined SysComment
func (m *SysComment) TableName() string {
	return "sys_comment"
}
