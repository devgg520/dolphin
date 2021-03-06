// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// LanguagePack defined
type LanguagePack struct {
	//
	LPId null.Int `xorm:"int(11) pk notnull autoincr 'l_p_id'" json:"l_p_id" xml:"l_p_id"`
	//
	LpName null.String `xorm:"varchar(200) 'lp_name'" json:"lp_name" xml:"lp_name"`
	//
	LpBegindate null.Time `xorm:"datetime 'lp_begindate'" json:"lp_begindate" xml:"lp_begindate"`
	//
	LpEnddate null.Time `xorm:"datetime 'lp_enddate'" json:"lp_enddate" xml:"lp_enddate"`
	//
	LpMoney null.Float `xorm:"float(11,2) 'lp_money'" json:"lp_money" xml:"lp_money"`
	//
	LpDesc null.String `xorm:"varchar(500) 'lp_desc'" json:"lp_desc" xml:"lp_desc"`
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
	AchievementNum null.Float `xorm:"float(50,2) 'achievement_num'" json:"achievement_num" xml:"achievement_num"`
	//
	Mzks null.Float `xorm:"float(50,2) 'mzks'" json:"mzks" xml:"mzks"`
	//
	OrderEdition null.Int `xorm:"int(11) 'order_edition'" json:"order_edition" xml:"order_edition"`
	//
	ZdAgreeLife null.Int `xorm:"int(11) 'zd_agree_life'" json:"zd_agree_life" xml:"zd_agree_life"`
}

// TableName table name of defined LanguagePack
func (m *LanguagePack) TableName() string {
	return "language_pack"
}
