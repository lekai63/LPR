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
	router.GET("/leasecontract/:argId/:argContractNo", GetLeaseContract)
	router.PUT("/leasecontract/:argId/:argContractNo", UpdateLeaseContract)
	router.DELETE("/leasecontract/:argId/:argContractNo", DeleteLeaseContract)
}

func configGinLeaseContractRouter(router gin.IRoutes) {
	router.GET("/leasecontract", ConverHttprouterToGin(GetAllLeaseContract))
	router.POST("/leasecontract", ConverHttprouterToGin(AddLeaseContract))
	router.GET("/leasecontract/:argId/:argContractNo", ConverHttprouterToGin(GetLeaseContract))
	router.PUT("/leasecontract/:argId/:argContractNo", ConverHttprouterToGin(UpdateLeaseContract))
	router.DELETE("/leasecontract/:argId/:argContractNo", ConverHttprouterToGin(DeleteLeaseContract))
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
// @Summary Get record from table LeaseContract by  argId  argContractNo
// @Tags LeaseContract
// @ID argId
// @ID argContractNo
// @Description GetLeaseContract is a function to get a single record from the lease_contract table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  argContractNo path string true "contract_no"
// @Success 200 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /leasecontract/{argId}/{argContractNo} [get]
// http "http://localhost:8080/leasecontract/1/hello world" X-Api-User:user123
func GetLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argContractNo, err := parseString(ps, "argContractNo")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLeaseContract(ctx, argId, argContractNo)
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
// echo '{"received_interest": 63,"updated_at": "2073-08-05T16:50:19.752232494Z","lessee": "DtBRdsNpuxuCfSCeSeZojsWko","fee": 43,"is_lpr": false,"current_reprice_day": "2053-10-01T14:35:24.24712212Z","next_reprice_day": "2081-12-27T06:32:29.985436509Z","current_lpr": 85,"current_rate": 83,"created_at": "2084-08-14T04:28:02.105836245Z","contract_no": "VHkHRSxXdVZZMBwCPjNEGYtKT","abbreviation": "UonadIOWRcvluLPmSQMuMrZcO","contract_principal": 44,"term_month": 50,"irr": 89,"id": 75,"start_date": "2248-08-31T07:45:41.919305751Z","actual_principal": 51,"subject_matter": "UEVwhfhItQudcfKqvJwByiFUy","received_principal": 56,"end_date": "2080-03-15T23:08:12.24363701Z","margin": 60,"lpr_plus": 71,"is_finished": true,"lessee_info_id": 46}' | http POST "http://localhost:8080/leasecontract" X-Api-User:user123
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
// @Param  argId path int true "id"// @Param  argContractNo path string true "contract_no"
// @Param  LeaseContract body model.LeaseContract true "Update LeaseContract record"
// @Success 200 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /leasecontract/{argId}/{argContractNo} [patch]
// echo '{"received_interest": 63,"updated_at": "2073-08-05T16:50:19.752232494Z","lessee": "DtBRdsNpuxuCfSCeSeZojsWko","fee": 43,"is_lpr": false,"current_reprice_day": "2053-10-01T14:35:24.24712212Z","next_reprice_day": "2081-12-27T06:32:29.985436509Z","current_lpr": 85,"current_rate": 83,"created_at": "2084-08-14T04:28:02.105836245Z","contract_no": "VHkHRSxXdVZZMBwCPjNEGYtKT","abbreviation": "UonadIOWRcvluLPmSQMuMrZcO","contract_principal": 44,"term_month": 50,"irr": 89,"id": 75,"start_date": "2248-08-31T07:45:41.919305751Z","actual_principal": 51,"subject_matter": "UEVwhfhItQudcfKqvJwByiFUy","received_principal": 56,"end_date": "2080-03-15T23:08:12.24363701Z","margin": 60,"lpr_plus": 71,"is_finished": true,"lessee_info_id": 46}' | http PUT "http://localhost:8080/leasecontract/1/hello world"  X-Api-User:user123
func UpdateLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argContractNo, err := parseString(ps, "argContractNo")
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
		argId, argContractNo,
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
// @Param  argId path int true "id"// @Param  argContractNo path string true "contract_no"
// @Success 204 {object} model.LeaseContract
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /leasecontract/{argId}/{argContractNo} [delete]
// http DELETE "http://localhost:8080/leasecontract/1/hello world" X-Api-User:user123
func DeleteLeaseContract(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	argContractNo, err := parseString(ps, "argContractNo")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lease_contract", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLeaseContract(ctx, argId, argContractNo)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
