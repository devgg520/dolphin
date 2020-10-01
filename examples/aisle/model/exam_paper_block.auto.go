// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ExamPaperBlock defined
type ExamPaperBlock struct {
	//
	EPBId null.Int `xorm:"int(11) pk notnull autoincr 'e_p_b_id'" json:"e_p_b_id" xml:"e_p_b_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	EpId null.Int `xorm:"int(11) 'ep_id'" json:"ep_id" xml:"ep_id"`
	//
	BlockName null.String `xorm:"varchar(50) 'block_name'" json:"block_name" xml:"block_name"`
	//
	BlockTimeLength null.Int `xorm:"int(11) 'block_time_length'" json:"block_time_length" xml:"block_time_length"`
	//
	BlockQuesNum null.Int `xorm:"int(11) 'block_ques_num'" json:"block_ques_num" xml:"block_ques_num"`
	//
	BlockScore null.Int `xorm:"int(11) 'block_score'" json:"block_score" xml:"block_score"`
	//
	BlockOrder null.Int `xorm:"int(11) 'block_order'" json:"block_order" xml:"block_order"`
	//
	BlockRemark null.String `xorm:"varchar(200) 'block_remark'" json:"block_remark" xml:"block_remark"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined ExamPaperBlock
func (m *ExamPaperBlock) TableName() string {
	return "exam_paper_block"
}
