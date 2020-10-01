// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Invoice defined
type Invoice struct {
	//
	InvoiceId null.Int `xorm:"int(11) pk notnull autoincr 'invoice_id'" json:"invoice_id" xml:"invoice_id"`
	//
	InvoiceNumber null.String `xorm:"varchar(100) 'invoice_number'" json:"invoice_number" xml:"invoice_number"`
	//
	InvoiceMoney null.Int `xorm:"int(11) 'invoice_money'" json:"invoice_money" xml:"invoice_money"`
	//
	InvoiceCompany null.String `xorm:"varchar(100) 'invoice_company'" json:"invoice_company" xml:"invoice_company"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	InvoiceType null.Int `xorm:"int(11) 'invoice_type'" json:"invoice_type" xml:"invoice_type"`
	//
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" xml:"fee_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	BillType null.Int `xorm:"int(11) 'bill_type'" json:"bill_type" xml:"bill_type"`
	//
	OrId null.Int `xorm:"int(11) 'or_id'" json:"or_id" xml:"or_id"`
}

// TableName table name of defined Invoice
func (m *Invoice) TableName() string {
	return "invoice"
}
