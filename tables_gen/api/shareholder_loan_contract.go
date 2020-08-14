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

func configShareholderLoanContractRouter(router *httprouter.Router) {
	router.GET("/shareholderloancontract", GetAllShareholderLoanContract)
	router.POST("/shareholderloancontract", AddShareholderLoanContract)
	router.GET("/shareholderloancontract/:argId", GetShareholderLoanContract)
	router.PUT("/shareholderloancontract/:argId", UpdateShareholderLoanContract)
	router.DELETE("/shareholderloancontract/:argId", DeleteShareholderLoanContract)
}

func configGinShareholderLoanContractRouter(router gin.IRoutes) {
	router.GET("/shareholderloancontract", ConverHttprouterToGin(GetAllShareholderLoanContract))
	router.POST("/shareholderloancontract", ConverHttprouterToGin(AddShareholderLoanContract))
	router.GET("/shareholderloancontract/:argId", ConverHttprouterToGin(GetShareholderLoanContract))
	router.PUT("/shareholderloancontract/:argId", ConverHttprouterToGin(UpdateShareholderLoanContract))
	router.DELETE("/shareholderloancontract/:argId", ConverHttprouterToGin(DeleteShareholderLoanContract))
}

// GetAllShareholderLoanContract is a function to get a slice of record(s) from shareholder_loan_contract table in the fzzl database
// @Summary Get list of ShareholderLoanContract
// @Tags ShareholderLoanContract
// @Description GetAllShareholderLoanContract is a handler to get a slice of record(s) from shareholder_loan_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ShareholderLoanContract}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloancontract [get]
// http "http://localhost:8080/shareholderloancontract?page=0&pagesize=20" X-Api-User:user123
func GetAllShareholderLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "shareholder_loan_contract", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllShareholderLoanContract(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetShareholderLoanContract is a function to get a single record from the shareholder_loan_contract table in the fzzl database
// @Summary Get record from table ShareholderLoanContract by  argId
// @Tags ShareholderLoanContract
// @ID argId
// @Description GetShareholderLoanContract is a function to get a single record from the shareholder_loan_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.ShareholderLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /shareholderloancontract/{argId} [get]
// http "http://localhost:8080/shareholderloancontract/1" X-Api-User:user123
func GetShareholderLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_contract", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetShareholderLoanContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddShareholderLoanContract add to add a single record to shareholder_loan_contract table in the fzzl database
// @Summary Add an record to shareholder_loan_contract table
// @Description add to add a single record to shareholder_loan_contract table in the fzzl database
// @Tags ShareholderLoanContract
// @Accept  json
// @Produce  json
// @Param ShareholderLoanContract body model.ShareholderLoanContract true "Add ShareholderLoanContract"
// @Success 200 {object} model.ShareholderLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloancontract [post]
// echo '{"abbreviation": "quQgIdJkfZqspZPDmVIKyvKcH","loan_rate": 20,"loan_contract_no": "OlJnOsAonuTcnsTrmnZNBvybr","all_repaid_principal": 99,"created_at": "2068-11-05T12:33:11.594084599+08:00","updated_at": "2275-07-26T05:57:44.150039016+08:00","id": 64,"creditor": "gwGudmyaueKhKTdXvhGcmFbWO","loan_principal": 38,"loan_start_date": "2187-04-27T17:57:19.022824796+08:00","loan_end_date": "2270-10-17T09:08:02.995601011+08:00","all_repaid_interest": 89,"is_finished": false}' | http POST "http://localhost:8080/shareholderloancontract" X-Api-User:user123
func AddShareholderLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	shareholderloancontract := &model.ShareholderLoanContract{}

	if err := readJSON(r, shareholderloancontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shareholderloancontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shareholderloancontract.Prepare()

	if err := shareholderloancontract.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_contract", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	shareholderloancontract, _, err = dao.AddShareholderLoanContract(ctx, shareholderloancontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shareholderloancontract)
}

// UpdateShareholderLoanContract Update a single record from shareholder_loan_contract table in the fzzl database
// @Summary Update an record in table shareholder_loan_contract
// @Description Update a single record from shareholder_loan_contract table in the fzzl database
// @Tags ShareholderLoanContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  ShareholderLoanContract body model.ShareholderLoanContract true "Update ShareholderLoanContract record"
// @Success 200 {object} model.ShareholderLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloancontract/{argId} [patch]
// echo '{"abbreviation": "quQgIdJkfZqspZPDmVIKyvKcH","loan_rate": 20,"loan_contract_no": "OlJnOsAonuTcnsTrmnZNBvybr","all_repaid_principal": 99,"created_at": "2068-11-05T12:33:11.594084599+08:00","updated_at": "2275-07-26T05:57:44.150039016+08:00","id": 64,"creditor": "gwGudmyaueKhKTdXvhGcmFbWO","loan_principal": 38,"loan_start_date": "2187-04-27T17:57:19.022824796+08:00","loan_end_date": "2270-10-17T09:08:02.995601011+08:00","all_repaid_interest": 89,"is_finished": false}' | http PUT "http://localhost:8080/shareholderloancontract/1"  X-Api-User:user123
func UpdateShareholderLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shareholderloancontract := &model.ShareholderLoanContract{}
	if err := readJSON(r, shareholderloancontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shareholderloancontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shareholderloancontract.Prepare()

	if err := shareholderloancontract.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_contract", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shareholderloancontract, _, err = dao.UpdateShareholderLoanContract(ctx,
		argId,
		shareholderloancontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shareholderloancontract)
}

// DeleteShareholderLoanContract Delete a single record from shareholder_loan_contract table in the fzzl database
// @Summary Delete a record from shareholder_loan_contract
// @Description Delete a single record from shareholder_loan_contract table in the fzzl database
// @Tags ShareholderLoanContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.ShareholderLoanContract
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /shareholderloancontract/{argId} [delete]
// http DELETE "http://localhost:8080/shareholderloancontract/1" X-Api-User:user123
func DeleteShareholderLoanContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_contract", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteShareholderLoanContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
