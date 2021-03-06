// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// WechatMessage defined
type WechatMessage struct {
	//
	T3230 null.Int `xorm:"int(11) pk notnull autoincr 't_323_0'" json:"t_323_0" xml:"t_323_0"`
	//
	Title null.String `xorm:"varchar(500) 'title'" json:"title" xml:"title"`
	//
	Content null.String `xorm:"varchar(2000) 'content'" json:"content" xml:"content"`
	//
	IfNotice null.Int `xorm:"int(11) 'if_notice'" json:"if_notice" xml:"if_notice"`
	//
	IfNoticecn null.String `xorm:"varchar 'if_noticecn'" json:"if_noticecn" xml:"if_noticecn"`
	//
	IfSee null.Int `xorm:"int(11) 'if_see'" json:"if_see" xml:"if_see"`
	//
	IfSeecn null.String `xorm:"varchar 'if_seecn'" json:"if_seecn" xml:"if_seecn"`
	//
	NoticeUser null.Int `xorm:"int(11) 'notice_user'" json:"notice_user" xml:"notice_user"`
	//
	NoticeName null.String `xorm:"varchar(500) 'notice_name'" json:"notice_name" xml:"notice_name"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined WechatMessage
func (m *WechatMessage) TableName() string {
	return "wechat_message"
}
