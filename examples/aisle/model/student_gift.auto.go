// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentGift defined
type StudentGift struct {
	//
	SGId null.Int `xorm:"int(11) pk notnull autoincr 's_g_id'" json:"s_g_id" xml:"s_g_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	GiftId null.Int `xorm:"int(11) 'gift_id'" json:"gift_id" xml:"gift_id"`
	//
	BuyWay null.Int `xorm:"int(11) 'buy_way'" json:"buy_way" xml:"buy_way"`
	//
	ResourcePrice null.Float `xorm:"float(11,2) 'resource_price'" json:"resource_price" xml:"resource_price"`
	//
	BuyPrice null.Float `xorm:"float(11,2) 'buy_price'" json:"buy_price" xml:"buy_price"`
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
	//
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" xml:"of_id"`
	//
	RefundMoney null.Float `xorm:"float(10,2) 'refund_money'" json:"refund_money" xml:"refund_money"`
	//
	RefundPrice null.Float `xorm:"float(10,2) 'refund_price'" json:"refund_price" xml:"refund_price"`
	//
	SgNum null.Float `xorm:"float(10,2) 'sg_num'" json:"sg_num" xml:"sg_num"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	GiftExchangeType null.Int `xorm:"int(11) 'gift_exchange_type'" json:"gift_exchange_type" xml:"gift_exchange_type"`
	//
	OggId null.Int `xorm:"int(11) 'ogg_id'" json:"ogg_id" xml:"ogg_id"`
	//
	UsePoint null.Float `xorm:"float(10,2) 'use_point'" json:"use_point" xml:"use_point"`
	//
	StuExGift null.Int `xorm:"int(11) 'stu_ex_gift'" json:"stu_ex_gift" xml:"stu_ex_gift"`
	//
	ZsGiveMonth null.Time `xorm:"datetime 'zs_give_month'" json:"zs_give_month" xml:"zs_give_month"`
}

// TableName table name of defined StudentGift
func (m *StudentGift) TableName() string {
	return "student_gift"
}
