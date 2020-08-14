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

func configShareholderLoanRepaidRecordRouter(router *httprouter.Router) {
	router.GET("/shareholderloanrepaidrecord", GetAllShareholderLoanRepaidRecord)
	router.POST("/shareholderloanrepaidrecord", AddShareholderLoanRepaidRecord)
	router.GET("/shareholderloanrepaidrecord/:argId", GetShareholderLoanRepaidRecord)
	router.PUT("/shareholderloanrepaidrecord/:argId", UpdateShareholderLoanRepaidRecord)
	router.DELETE("/shareholderloanrepaidrecord/:argId", DeleteShareholderLoanRepaidRecord)
}

func configGinShareholderLoanRepaidRecordRouter(router gin.IRoutes) {
	router.GET("/shareholderloanrepaidrecord", ConverHttprouterToGin(GetAllShareholderLoanRepaidRecord))
	router.POST("/shareholderloanrepaidrecord", ConverHttprouterToGin(AddShareholderLoanRepaidRecord))
	router.GET("/shareholderloanrepaidrecord/:argId", ConverHttprouterToGin(GetShareholderLoanRepaidRecord))
	router.PUT("/shareholderloanrepaidrecord/:argId", ConverHttprouterToGin(UpdateShareholderLoanRepaidRecord))
	router.DELETE("/shareholderloanrepaidrecord/:argId", ConverHttprouterToGin(DeleteShareholderLoanRepaidRecord))
}

// GetAllShareholderLoanRepaidRecord is a function to get a slice of record(s) from shareholder_loan_repaid_record table in the fzzl database
// @Summary Get list of ShareholderLoanRepaidRecord
// @Tags ShareholderLoanRepaidRecord
// @Description GetAllShareholderLoanRepaidRecord is a handler to get a slice of record(s) from shareholder_loan_repaid_record table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ShareholderLoanRepaidRecord}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloanrepaidrecord [get]
// http "http://localhost:8080/shareholderloanrepaidrecord?page=0&pagesize=20" X-Api-User:user123
func GetAllShareholderLoanRepaidRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "shareholder_loan_repaid_record", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllShareholderLoanRepaidRecord(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetShareholderLoanRepaidRecord is a function to get a single record from the shareholder_loan_repaid_record table in the fzzl database
// @Summary Get record from table ShareholderLoanRepaidRecord by  argId
// @Tags ShareholderLoanRepaidRecord
// @ID argId
// @Description GetShareholderLoanRepaidRecord is a function to get a single record from the shareholder_loan_repaid_record table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.ShareholderLoanRepaidRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /shareholderloanrepaidrecord/{argId} [get]
// http "http://localhost:8080/shareholderloanrepaidrecord/1" X-Api-User:user123
func GetShareholderLoanRepaidRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_repaid_record", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetShareholderLoanRepaidRecord(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddShareholderLoanRepaidRecord add to add a single record to shareholder_loan_repaid_record table in the fzzl database
// @Summary Add an record to shareholder_loan_repaid_record table
// @Description add to add a single record to shareholder_loan_repaid_record table in the fzzl database
// @Tags ShareholderLoanRepaidRecord
// @Accept  json
// @Produce  json
// @Param ShareholderLoanRepaidRecord body model.ShareholderLoanRepaidRecord true "Add ShareholderLoanRepaidRecord"
// @Success 200 {object} model.ShareholderLoanRepaidRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloanrepaidrecord [post]
// echo '{"created_at": "2260-09-27T05:42:50.384349782+08:00","updated_at": "2025-02-10T10:24:54.026179139+08:00","id": 83,"repaid_date": "2059-07-14T11:05:52.232246349+08:00","repaid_amount": 91,"repaid_principal": 43,"repaid_interest": 52,"shareholder_loan_contract_id": 15}' | http POST "http://localhost:8080/shareholderloanrepaidrecord" X-Api-User:user123
func AddShareholderLoanRepaidRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	shareholderloanrepaidrecord := &model.ShareholderLoanRepaidRecord{}

	if err := readJSON(r, shareholderloanrepaidrecord); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shareholderloanrepaidrecord.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shareholderloanrepaidrecord.Prepare()

	if err := shareholderloanrepaidrecord.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_repaid_record", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	shareholderloanrepaidrecord, _, err = dao.AddShareholderLoanRepaidRecord(ctx, shareholderloanrepaidrecord)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shareholderloanrepaidrecord)
}

// UpdateShareholderLoanRepaidRecord Update a single record from shareholder_loan_repaid_record table in the fzzl database
// @Summary Update an record in table shareholder_loan_repaid_record
// @Description Update a single record from shareholder_loan_repaid_record table in the fzzl database
// @Tags ShareholderLoanRepaidRecord
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  ShareholderLoanRepaidRecord body model.ShareholderLoanRepaidRecord true "Update ShareholderLoanRepaidRecord record"
// @Success 200 {object} model.ShareholderLoanRepaidRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shareholderloanrepaidrecord/{argId} [patch]
// echo '{"created_at": "2260-09-27T05:42:50.384349782+08:00","updated_at": "2025-02-10T10:24:54.026179139+08:00","id": 83,"repaid_date": "2059-07-14T11:05:52.232246349+08:00","repaid_amount": 91,"repaid_principal": 43,"repaid_interest": 52,"shareholder_loan_contract_id": 15}' | http PUT "http://localhost:8080/shareholderloanrepaidrecord/1"  X-Api-User:user123
func UpdateShareholderLoanRepaidRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shareholderloanrepaidrecord := &model.ShareholderLoanRepaidRecord{}
	if err := readJSON(r, shareholderloanrepaidrecord); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shareholderloanrepaidrecord.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shareholderloanrepaidrecord.Prepare()

	if err := shareholderloanrepaidrecord.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_repaid_record", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shareholderloanrepaidrecord, _, err = dao.UpdateShareholderLoanRepaidRecord(ctx,
		argId,
		shareholderloanrepaidrecord)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shareholderloanrepaidrecord)
}

// DeleteShareholderLoanRepaidRecord Delete a single record from shareholder_loan_repaid_record table in the fzzl database
// @Summary Delete a record from shareholder_loan_repaid_record
// @Description Delete a single record from shareholder_loan_repaid_record table in the fzzl database
// @Tags ShareholderLoanRepaidRecord
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.ShareholderLoanRepaidRecord
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /shareholderloanrepaidrecord/{argId} [delete]
// http DELETE "http://localhost:8080/shareholderloanrepaidrecord/1" X-Api-User:user123
func DeleteShareholderLoanRepaidRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shareholder_loan_repaid_record", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteShareholderLoanRepaidRecord(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
