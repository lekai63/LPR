package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "lease_repay_plan" => http://localhost:9033/admin/info/lease_repay_plan
// "bank_loan_contract" => http://localhost:9033/admin/info/bank_loan_contract
// "bank_repay_plan" => http://localhost:9033/admin/info/bank_repay_plan
// "lease_contract" => http://localhost:9033/admin/info/lease_contract
// "lessee_info" => http://localhost:9033/admin/info/lessee_info
// "shareholder_loan_contract" => http://localhost:9033/admin/info/shareholder_loan_contract
// "shareholder_loan_repaid_record" => http://localhost:9033/admin/info/shareholder_loan_repaid_record
//
// "bank_loan_contract" => http://localhost:9033/admin/info/bank_loan_contract
// "bank_repay_plan" => http://localhost:9033/admin/info/bank_repay_plan
// "lease_repay_plan" => http://localhost:9033/admin/info/lease_repay_plan
// "shareholder_loan_contract" => http://localhost:9033/admin/info/shareholder_loan_contract
// "shareholder_loan_repaid_record" => http://localhost:9033/admin/info/shareholder_loan_repaid_record
//
// example end
//
var Generators = map[string]table.Generator{

	// tables
	"lease_contract":          GetLeaseContractTable,
	"lessee_info":             GetLesseeInfoTable,
	"lease_repay_plan":        GetLeaseRepayPlanTable,
	"lease_repay_plan_import": GetLeaseRepayPlanImportTable,

	"bank_loan_contract":     GetBankLoanContractTable,
	"bank_repay_plan":        GetBankRepayPlanTable,
	"bank_repay_plan_import": GetBankRepayPlanImportTable,

	"shareholder_loan_contract":      GetShareholderLoanContractTable,
	"shareholder_loan_repaid_record": GetShareholderLoanRepaidRecordTable,

	// reports

	// generators end
}
