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

func configBankRepayPlanRouter(router *httprouter.Router) {
	router.GET("/bankrepayplan", GetAllBankRepayPlan)
	router.POST("/bankrepayplan", AddBankRepayPlan)
	router.GET("/bankrepayplan/:argId", GetBankRepayPlan)
	router.PUT("/bankrepayplan/:argId", UpdateBankRepayPlan)
	router.DELETE("/bankrepayplan/:argId", DeleteBankRepayPlan)
}

func configGinBankRepayPlanRouter(router gin.IRoutes) {
	router.GET("/bankrepayplan", ConverHttprouterToGin(GetAllBankRepayPlan))
	router.POST("/bankrepayplan", ConverHttprouterToGin(AddBankRepayPlan))
	router.GET("/bankrepayplan/:argId", ConverHttprouterToGin(GetBankRepayPlan))
	router.PUT("/bankrepayplan/:argId", ConverHttprouterToGin(UpdateBankRepayPlan))
	router.DELETE("/bankrepayplan/:argId", ConverHttprouterToGin(DeleteBankRepayPlan))
}

// GetAllBankRepayPlan is a function to get a slice of record(s) from bank_repay_plan table in the fzzl database
// @Summary Get list of BankRepayPlan
// @Tags BankRepayPlan
// @Description GetAllBankRepayPlan is a handler to get a slice of record(s) from bank_repay_plan table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BankRepayPlan}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankrepayplan [get]
// http "http://localhost:8080/bankrepayplan?page=0&pagesize=20" X-Api-User:user123
func GetAllBankRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "bank_repay_plan", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBankRepayPlan(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBankRepayPlan is a function to get a single record from the bank_repay_plan table in the fzzl database
// @Summary Get record from table BankRepayPlan by  argId
// @Tags BankRepayPlan
// @ID argId
// @Description GetBankRepayPlan is a function to get a single record from the bank_repay_plan table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.BankRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /bankrepayplan/{argId} [get]
// http "http://localhost:8080/bankrepayplan/1" X-Api-User:user123
func GetBankRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_repay_plan", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBankRepayPlan(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBankRepayPlan add to add a single record to bank_repay_plan table in the fzzl database
// @Summary Add an record to bank_repay_plan table
// @Description add to add a single record to bank_repay_plan table in the fzzl database
// @Tags BankRepayPlan
// @Accept  json
// @Produce  json
// @Param BankRepayPlan body model.BankRepayPlan true "Add BankRepayPlan"
// @Success 200 {object} model.BankRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankrepayplan [post]
// echo '{"plan_amount": 91,"plan_interest": 66,"actual_amount": 26,"actual_principal": 26,"updated_at": "2300-01-18T08:02:37.147327156Z","id": 67,"bank_loan_contract_id": 16,"plan_date": "2159-08-03T14:14:14.922878857Z","plan_principal": 90,"actual_date": "2084-06-07T08:57:10.753963539Z","actual_interest": 35,"created_at": "2219-01-24T11:56:18.51500847Z"}' | http POST "http://localhost:8080/bankrepayplan" X-Api-User:user123
func AddBankRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	bankrepayplan := &model.BankRepayPlan{}

	if err := readJSON(r, bankrepayplan); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := bankrepayplan.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	bankrepayplan.Prepare()

	if err := bankrepayplan.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_repay_plan", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	bankrepayplan, _, err = dao.AddBankRepayPlan(ctx, bankrepayplan)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, bankrepayplan)
}

// UpdateBankRepayPlan Update a single record from bank_repay_plan table in the fzzl database
// @Summary Update an record in table bank_repay_plan
// @Description Update a single record from bank_repay_plan table in the fzzl database
// @Tags BankRepayPlan
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  BankRepayPlan body model.BankRepayPlan true "Update BankRepayPlan record"
// @Success 200 {object} model.BankRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /bankrepayplan/{argId} [patch]
// echo '{"plan_amount": 91,"plan_interest": 66,"actual_amount": 26,"actual_principal": 26,"updated_at": "2300-01-18T08:02:37.147327156Z","id": 67,"bank_loan_contract_id": 16,"plan_date": "2159-08-03T14:14:14.922878857Z","plan_principal": 90,"actual_date": "2084-06-07T08:57:10.753963539Z","actual_interest": 35,"created_at": "2219-01-24T11:56:18.51500847Z"}' | http PUT "http://localhost:8080/bankrepayplan/1"  X-Api-User:user123
func UpdateBankRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	bankrepayplan := &model.BankRepayPlan{}
	if err := readJSON(r, bankrepayplan); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := bankrepayplan.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	bankrepayplan.Prepare()

	if err := bankrepayplan.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_repay_plan", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	bankrepayplan, _, err = dao.UpdateBankRepayPlan(ctx,
		argId,
		bankrepayplan)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, bankrepayplan)
}

// DeleteBankRepayPlan Delete a single record from bank_repay_plan table in the fzzl database
// @Summary Delete a record from bank_repay_plan
// @Description Delete a single record from bank_repay_plan table in the fzzl database
// @Tags BankRepayPlan
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.BankRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /bankrepayplan/{argId} [delete]
// http DELETE "http://localhost:8080/bankrepayplan/1" X-Api-User:user123
func DeleteBankRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "bank_repay_plan", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBankRepayPlan(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
