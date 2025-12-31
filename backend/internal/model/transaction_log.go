package model

import (
	"database/sql"
	"time"
)

type TransactionLog struct {
	ID             int64        `json:"id" gorm:"primaryKey;autoIncrement;column:id" ui:"sortable"`
	SerialNumber   string       `json:"serial_number" gorm:"column:serial_number;size:50" ui:"visible;filterable;sortable"`
	DeviceID       string       `json:"device_id" gorm:"column:device_id;size:100" ui:"filterable;sortable"`
	TrxID          string       `json:"trx_id" gorm:"column:trx_id;size:80;not null;unique" ui:"visible;filterable;sortable"`
	TrxType        string       `json:"trx_type" gorm:"column:trx_type;size:255" ui:"visible;filterable;sortable"`
	Amount         int64        `json:"amount" gorm:"column:amount" ui:"visible;filterable;sortable"`
	TrxDate        sql.NullTime `json:"trx_date" gorm:"column:trx_date" ui:"visible;filterable;sortable"`
	TrackKSNIndex  string       `json:"track_ksn_index" gorm:"column:track_ksn_index;size:255"`
	AmountKSNIndex string       `json:"amount_ksn_index" gorm:"column:amount_ksn_index;size:255"`
	EMVKSNIndex    string       `json:"emv_ksn_index" gorm:"column:emv_ksn_index;size:255"`
	PINKSNIndex    string       `json:"pin_ksn_index" gorm:"column:pin_ksn_index;size:255"`
	TID            string       `json:"tid" gorm:"column:tid;size:20" ui:"visible;filterable;sortable"`
	MID            string       `json:"mid" gorm:"column:mid;size:20" ui:"visible;filterable;sortable"`
	ResponseCode   string       `json:"response_code" gorm:"column:response_code;size:10" ui:"visible;filterable;sortable"`
	RequestData    string       `json:"request_data" gorm:"column:request_data;type:longtext"`
	ResponseData   string       `json:"response_data" gorm:"column:response_data;type:longtext"`
	BatchNum       string       `json:"batch_num" gorm:"column:batch_num;size:50" ui:"visible;filterable;sortable"`
	ApprovalCode   string       `json:"approval_code" gorm:"column:approval_code;size:50" ui:"visible;filterable;sortable"`
	InvoiceNum     string       `json:"invoice_num" gorm:"column:invoice_num;size:50" ui:"visible;filterable;sortable"`
	RRN            string       `json:"rrn" gorm:"column:rrn;size:50" ui:"visible;filterable;sortable"`
	AccountNumber  string       `json:"account_number" gorm:"column:account_number;size:255" ui:"visible;filterable;sortable"`
	AccountName    string       `json:"account_name" gorm:"column:account_name;size:255" ui:"visible;filterable;sortable"`
	AccountBank    string       `json:"account_bank" gorm:"column:account_bank;size:255" ui:"visible;filterable;sortable"`
	Codebic        string       `json:"codebic" gorm:"column:codebic;size:255"`
	SpecialTrx     bool         `json:"special_trx" gorm:"column:special_trx" ui:"visible;filterable;sortable"`
	PaidTime       sql.NullTime `json:"paid_time" gorm:"column:paid_time" ui:"visible;filterable;sortable"`
	SendToOdoo     bool         `json:"send_to_odoo" gorm:"column:send_to_odoo" ui:"visible;filterable;sortable"`
	SendToOdooAt   sql.NullTime `json:"send_to_odoo_at" gorm:"column:send_to_odoo_at" ui:"visible;filterable;sortable"`
	CreatedAt      time.Time    `json:"created_at" gorm:"column:created_at" ui:"visible;filterable;sortable"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"column:updated_at" ui:"visible;filterable;sortable"`
}

func (TransactionLog) TableName() string {
	return "transaction_logs"
}
