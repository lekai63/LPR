package api

import (
	"net/http"

	"github.com/lekai63/lpr/tables_gen/dao"
	"github.com/lekai63/lpr/tables_gen/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configBankLoanContractRouter(router *httprouter.Router) {
	router.GET("/bankloancontract", GetAllBankLoanContract)
	router.POST("/bankloancontract", AddBankLoanContract)
	router.GET("/bankloancontract/:argId", GetBankLoanContract)
	router.PUT("/bankloancontract/:argId", UpdateBankLoanContract)
	router.DELETE("/bankloancontract/:argId", DeleteBankLoanContract)
}

func configGinBankLoanContractRouter(router gin.IRoutes) {
	router.GET("/bankloancontract", ConverHttprouterToGin(GetAllBankLoanContract))
	router.POST("/bankloancontract", ConverHttprouterToGin(AddBankLoanContract))
	router.GET("/bankloancontract/:argId", ConverHttprouterToGin(GetBankLoanContract))
	router.PUT("/bankloancontract/:argId", ConverHttprouterToGin(UpdateBankLoanContract))
	router.DELETE("/bankloancontract/:argId", ConverHttprouterToGin(DeleteBankLoanContract))
}

// GetAllBankLoanContract is a function to get a slice of record(s) from bank_loan_contract table in the fzzl database
// @Summary Get list of BankLoanContract
// @Tags BankLoanContract
// @Description GetAllBankLoanContract is a handler to get a slice of record(s) from bank_loan_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BankLoanContract}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankloancontract [get]
// http "http://localhost:8080/bankloancontract?page=0&pagesize=20" X-Api-User:user123
func GetAllBankLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "bank_loan_contract", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBankLoanContract(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBankLoanContract is a function to get a single record from the bank_loan_contract table in the fzzl database
// @Summary Get record from table BankLoanContract by  argId
// @Tags BankLoanContract
// @ID argId
// @Description GetBankLoanContract is a function to get a single record from the bank_loan_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.BankLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /bankloancontract/{argId} [get]
// http "http://localhost:8080/bankloancontract/1" X-Api-User:user123
func GetBankLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_loan_contract", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBankLoanContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBankLoanContract add to add a single record to bank_loan_contract table in the fzzl database
// @Summary Add an record to bank_loan_contract table
// @Description add to add a single record to bank_loan_contract table in the fzzl database
// @Tags BankLoanContract
// @Accept  json
// @Produce  json
// @Param BankLoanContract body model.BankLoanContract true "Add BankLoanContract"
// @Success 200 {object} model.BankLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankloancontract [post]
// echo '{"updated_at": "2256-03-05T04:48:12.728069936Z","bank_branch": "nSpcnebeejEROBXQPGkGRudDW","loan_method": "osceXdoUpXtMDDvoZbMnBRPjv","contract_end_date": "2160-05-28T15:06:14.640345944Z","bank_account": "GULpfKwHPBPCOGGhFpXfiEJPg","contract_start_date": "2158-10-16T01:57:11.075148375Z","actual_start_date": "2248-01-16T13:01:58.291957537Z","is_lpr": true,"current_reprice_day": "2070-06-08T03:29:31.098285897Z","current_lpr": 43,"bank_contract_no": "NjvsSUtingPieqcliYdCvjVgH","bank_contract_name": "AsDYDIJHPlSFJERjcrGgncbin","interest_calc_method": "SDexZKvrRawdhYmTKBGkuIATP","contact_tel": "jQoPgKUglteidwSNYThwGOLwk","current_rate": 25,"all_repaid_interest": 98,"is_finished": true,"loan_principal": 56,"lpr_plus": 51,"next_reprice_day": "2124-10-05T13:36:53.442867265Z","all_repaid_principal": 14,"contact_person": "ekqcTkbYNOuDaTGUoxueCYGZD","id": 15,"lease_contract_id": 35,"bank_name": "CTgYtxVhiQCWowmsXncKKAGOw","created_at": "2098-08-23T10:17:04.730191942Z"}' | http POST "http://localhost:8080/bankloancontract" X-Api-User:user123
func AddBankLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	bankloancontract := &model.BankLoanContract{}

	if err := readJSON(r, bankloancontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := bankloancontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	bankloancontract.Prepare()

	if err := bankloancontract.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_loan_contract", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	bankloancontract, _, err = dao.AddBankLoanContract(ctx, bankloancontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, bankloancontract)
}

// UpdateBankLoanContract Update a single record from bank_loan_contract table in the fzzl database
// @Summary Update an record in table bank_loan_contract
// @Description Update a single record from bank_loan_contract table in the fzzl database
// @Tags BankLoanContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  BankLoanContract body model.BankLoanContract true "Update BankLoanContract record"
// @Success 200 {object} model.BankLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankloancontract/{argId} [patch]
// echo '{"updated_at": "2256-03-05T04:48:12.728069936Z","bank_branch": "nSpcnebeejEROBXQPGkGRudDW","loan_method": "osceXdoUpXtMDDvoZbMnBRPjv","contract_end_date": "2160-05-28T15:06:14.640345944Z","bank_account": "GULpfKwHPBPCOGGhFpXfiEJPg","contract_start_date": "2158-10-16T01:57:11.075148375Z","actual_start_date": "2248-01-16T13:01:58.291957537Z","is_lpr": true,"current_reprice_day": "2070-06-08T03:29:31.098285897Z","current_lpr": 43,"bank_contract_no": "NjvsSUtingPieqcliYdCvjVgH","bank_contract_name": "AsDYDIJHPlSFJERjcrGgncbin","interest_calc_method": "SDexZKvrRawdhYmTKBGkuIATP","contact_tel": "jQoPgKUglteidwSNYThwGOLwk","current_rate": 25,"all_repaid_interest": 98,"is_finished": true,"loan_principal": 56,"lpr_plus": 51,"next_reprice_day": "2124-10-05T13:36:53.442867265Z","all_repaid_principal": 14,"contact_person": "ekqcTkbYNOuDaTGUoxueCYGZD","id": 15,"lease_contract_id": 35,"bank_name": "CTgYtxVhiQCWowmsXncKKAGOw","created_at": "2098-08-23T10:17:04.730191942Z"}' | http PUT "http://localhost:8080/bankloancontract/1"  X-Api-User:user123
func UpdateBankLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	bankloancontract := &model.BankLoanContract{}
	if err := readJSON(r, bankloancontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := bankloancontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	bankloancontract.Prepare()

	if err := bankloancontract.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_loan_contract", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	bankloancontract, _, err = dao.UpdateBankLoanContract(ctx,
		argId,
		bankloancontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, bankloancontract)
}

// DeleteBankLoanContract Delete a single record from bank_loan_contract table in the fzzl database
// @Summary Delete a record from bank_loan_contract
// @Description Delete a single record from bank_loan_contract table in the fzzl database
// @Tags BankLoanContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.BankLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /bankloancontract/{argId} [delete]
// http DELETE "http://localhost:8080/bankloancontract/1" X-Api-User:user123
func DeleteBankLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_loan_contract", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBankLoanContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
