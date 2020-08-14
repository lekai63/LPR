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

func configLeaseContractRouter(router *httprouter.Router) {
	router.GET("/leasecontract", GetAllLeaseContract)
	router.POST("/leasecontract", AddLeaseContract)
	router.GET("/leasecontract/:argId", GetLeaseContract)
	router.PUT("/leasecontract/:argId", UpdateLeaseContract)
	router.DELETE("/leasecontract/:argId", DeleteLeaseContract)
}

func configGinLeaseContractRouter(router gin.IRoutes) {
	router.GET("/leasecontract", ConverHttprouterToGin(GetAllLeaseContract))
	router.POST("/leasecontract", ConverHttprouterToGin(AddLeaseContract))
	router.GET("/leasecontract/:argId", ConverHttprouterToGin(GetLeaseContract))
	router.PUT("/leasecontract/:argId", ConverHttprouterToGin(UpdateLeaseContract))
	router.DELETE("/leasecontract/:argId", ConverHttprouterToGin(DeleteLeaseContract))
}

// GetAllLeaseContract is a function to get a slice of record(s) from lease_contract table in the fzzl database
// @Summary Get list of LeaseContract
// @Tags LeaseContract
// @Description GetAllLeaseContract is a handler to get a slice of record(s) from lease_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.LeaseContract}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leasecontract [get]
// http "http://localhost:8080/leasecontract?page=0&pagesize=20" X-Api-User:user123
func GetAllLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "lease_contract", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLeaseContract(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLeaseContract is a function to get a single record from the lease_contract table in the fzzl database
// @Summary Get record from table LeaseContract by  argId
// @Tags LeaseContract
// @ID argId
// @Description GetLeaseContract is a function to get a single record from the lease_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /leasecontract/{argId} [get]
// http "http://localhost:8080/leasecontract/1" X-Api-User:user123
func GetLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLeaseContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLeaseContract add to add a single record to lease_contract table in the fzzl database
// @Summary Add an record to lease_contract table
// @Description add to add a single record to lease_contract table in the fzzl database
// @Tags LeaseContract
// @Accept  json
// @Produce  json
// @Param LeaseContract body model.LeaseContract true "Add LeaseContract"
// @Success 200 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leasecontract [post]
// echo '{"current_reprice_day": "2227-12-28T05:18:17.629928242+08:00","current_rate": 68,"next_reprice_day": "2308-04-03T11:08:59.327243366+08:00","fee": 7,"margin": 23,"contract_principal": 15,"term_month": 94,"is_lpr": true,"lessee_info_id": 74,"id": 50,"lessee": "becSGjPBEjbxtgDtsurletkNh","end_date": "2283-12-12T11:14:37.790091817+08:00","current_lpr": 78,"updated_at": "2300-05-29T03:35:08.442019234+08:00","received_principal": 79,"received_interest": 80,"is_finished": false,"contract_no": "DdSPGdATETpDeTIGOQyueatkZ","abbreviation": "fDoDOTTJySAnNossJrSycoRFN","start_date": "2297-12-02T06:11:37.732038748+08:00","actual_principal": 7,"irr": 4,"subject_matter": "WyQVjTushAnpbmsthUZgGlHoU","lpr_plus": 18,"created_at": "2105-07-22T15:23:38.494368994+08:00"}' | http POST "http://localhost:8080/leasecontract" X-Api-User:user123
func AddLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	leasecontract := &model.LeaseContract{}

	if err := readJSON(r, leasecontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leasecontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leasecontract.Prepare()

	if err := leasecontract.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	leasecontract, _, err = dao.AddLeaseContract(ctx, leasecontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leasecontract)
}

// UpdateLeaseContract Update a single record from lease_contract table in the fzzl database
// @Summary Update an record in table lease_contract
// @Description Update a single record from lease_contract table in the fzzl database
// @Tags LeaseContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  LeaseContract body model.LeaseContract true "Update LeaseContract record"
// @Success 200 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leasecontract/{argId} [patch]
// echo '{"current_reprice_day": "2227-12-28T05:18:17.629928242+08:00","current_rate": 68,"next_reprice_day": "2308-04-03T11:08:59.327243366+08:00","fee": 7,"margin": 23,"contract_principal": 15,"term_month": 94,"is_lpr": true,"lessee_info_id": 74,"id": 50,"lessee": "becSGjPBEjbxtgDtsurletkNh","end_date": "2283-12-12T11:14:37.790091817+08:00","current_lpr": 78,"updated_at": "2300-05-29T03:35:08.442019234+08:00","received_principal": 79,"received_interest": 80,"is_finished": false,"contract_no": "DdSPGdATETpDeTIGOQyueatkZ","abbreviation": "fDoDOTTJySAnNossJrSycoRFN","start_date": "2297-12-02T06:11:37.732038748+08:00","actual_principal": 7,"irr": 4,"subject_matter": "WyQVjTushAnpbmsthUZgGlHoU","lpr_plus": 18,"created_at": "2105-07-22T15:23:38.494368994+08:00"}' | http PUT "http://localhost:8080/leasecontract/1"  X-Api-User:user123
func UpdateLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leasecontract := &model.LeaseContract{}
	if err := readJSON(r, leasecontract); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := leasecontract.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	leasecontract.Prepare()

	if err := leasecontract.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	leasecontract, _, err = dao.UpdateLeaseContract(ctx,
		argId,
		leasecontract)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, leasecontract)
}

// DeleteLeaseContract Delete a single record from lease_contract table in the fzzl database
// @Summary Delete a record from lease_contract
// @Description Delete a single record from lease_contract table in the fzzl database
// @Tags LeaseContract
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /leasecontract/{argId} [delete]
// http DELETE "http://localhost:8080/leasecontract/1" X-Api-User:user123
func DeleteLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLeaseContract(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
