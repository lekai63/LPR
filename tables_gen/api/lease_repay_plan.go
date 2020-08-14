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

func configLeaseRepayPlanRouter(router *httprouter.Router) {
	router.GET("/leaserepayplan", GetAllLeaseRepayPlan)
	router.POST("/leaserepayplan", AddLeaseRepayPlan)
	router.GET("/leaserepayplan/:argId", GetLeaseRepayPlan)
	router.PUT("/leaserepayplan/:argId", UpdateLeaseRepayPlan)
	router.DELETE("/leaserepayplan/:argId", DeleteLeaseRepayPlan)
}

func configGinLeaseRepayPlanRouter(router gin.IRoutes) {
	router.GET("/leaserepayplan", ConverHttprouterToGin(GetAllLeaseRepayPlan))
	router.POST("/leaserepayplan", ConverHttprouterToGin(AddLeaseRepayPlan))
	router.GET("/leaserepayplan/:argId", ConverHttprouterToGin(GetLeaseRepayPlan))
	router.PUT("/leaserepayplan/:argId", ConverHttprouterToGin(UpdateLeaseRepayPlan))
	router.DELETE("/leaserepayplan/:argId", ConverHttprouterToGin(DeleteLeaseRepayPlan))
}

// GetAllLeaseRepayPlan is a function to get a slice of record(s) from lease_repay_plan table in the fzzl database
// @Summary Get list of LeaseRepayPlan
// @Tags LeaseRepayPlan
// @Description GetAllLeaseRepayPlan is a handler to get a slice of record(s) from lease_repay_plan table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.LeaseRepayPlan}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leaserepayplan [get]
// http "http://localhost:8080/leaserepayplan?page=0&pagesize=20" X-Api-User:user123
func GetAllLeaseRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "lease_repay_plan", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLeaseRepayPlan(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLeaseRepayPlan is a function to get a single record from the lease_repay_plan table in the fzzl database
// @Summary Get record from table LeaseRepayPlan by  argId
// @Tags LeaseRepayPlan
// @ID argId
// @Description GetLeaseRepayPlan is a function to get a single record from the lease_repay_plan table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.LeaseRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /leaserepayplan/{argId} [get]
// http "http://localhost:8080/leaserepayplan/1" X-Api-User:user123
func GetLeaseRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_repay_plan", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLeaseRepayPlan(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLeaseRepayPlan add to add a single record to lease_repay_plan table in the fzzl database
// @Summary Add an record to lease_repay_plan table
// @Description add to add a single record to lease_repay_plan table in the fzzl database
// @Tags LeaseRepayPlan
// @Accept  json
// @Produce  json
// @Param LeaseRepayPlan body model.LeaseRepayPlan true "Add LeaseRepayPlan"
// @Success 200 {object} model.LeaseRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leaserepayplan [post]
// echo '{"period": 70,"plan_amount": 78,"actual_date": "2032-06-16T21:40:39.011280264+08:00","actual_amount": 66,"created_at": "2170-08-30T00:56:07.420982522+08:00","updated_at": "2166-11-06T17:35:51.032828976+08:00","id": 12,"plan_date": "2113-02-01T04:32:14.216889545+08:00","plan_principal": 74,"plan_interest": 71,"actual_principal": 65,"actual_interest": 37,"lease_contract_id": 20}' | http POST "http://localhost:8080/leaserepayplan" X-Api-User:user123
func AddLeaseRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	leaserepayplan := &model.LeaseRepayPlan{}

	if err := readJSON(r, leaserepayplan); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leaserepayplan.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leaserepayplan.Prepare()

	if err := leaserepayplan.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_repay_plan", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	leaserepayplan, _, err = dao.AddLeaseRepayPlan(ctx, leaserepayplan)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leaserepayplan)
}

// UpdateLeaseRepayPlan Update a single record from lease_repay_plan table in the fzzl database
// @Summary Update an record in table lease_repay_plan
// @Description Update a single record from lease_repay_plan table in the fzzl database
// @Tags LeaseRepayPlan
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  LeaseRepayPlan body model.LeaseRepayPlan true "Update LeaseRepayPlan record"
// @Success 200 {object} model.LeaseRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leaserepayplan/{argId} [patch]
// echo '{"period": 70,"plan_amount": 78,"actual_date": "2032-06-16T21:40:39.011280264+08:00","actual_amount": 66,"created_at": "2170-08-30T00:56:07.420982522+08:00","updated_at": "2166-11-06T17:35:51.032828976+08:00","id": 12,"plan_date": "2113-02-01T04:32:14.216889545+08:00","plan_principal": 74,"plan_interest": 71,"actual_principal": 65,"actual_interest": 37,"lease_contract_id": 20}' | http PUT "http://localhost:8080/leaserepayplan/1"  X-Api-User:user123
func UpdateLeaseRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leaserepayplan := &model.LeaseRepayPlan{}
	if err := readJSON(r, leaserepayplan); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leaserepayplan.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leaserepayplan.Prepare()

	if err := leaserepayplan.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_repay_plan", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leaserepayplan, _, err = dao.UpdateLeaseRepayPlan(ctx,
		argId,
		leaserepayplan)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leaserepayplan)
}

// DeleteLeaseRepayPlan Delete a single record from lease_repay_plan table in the fzzl database
// @Summary Delete a record from lease_repay_plan
// @Description Delete a single record from lease_repay_plan table in the fzzl database
// @Tags LeaseRepayPlan
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.LeaseRepayPlan
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /leaserepayplan/{argId} [delete]
// http DELETE "http://localhost:8080/leaserepayplan/1" X-Api-User:user123
func DeleteLeaseRepayPlan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_repay_plan", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLeaseRepayPlan(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
