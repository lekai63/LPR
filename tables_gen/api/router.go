package api

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/lekai63/lpr/tables_gen/dao"
	"github.com/lekai63/lpr/tables_gen/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_             = time.Second // import time.Second for unknown usage in api
	crudEndpoints map[string]*CrudAPI
)

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"create_url"`
	RetrieveOneURL  string           `json:"retrieve_one_url"`
	RetrieveManyURL string           `json:"retrieve_many_url"`
	UpdateURL       string           `json:"update_url"`
	DeleteURL       string           `json:"delete_url"`
	FetchDDLURL     string           `json:"fetch_ddl_url"`
	TableInfo       *model.TableInfo `json:"table_info"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigRouter configure http.Handler router
func ConfigRouter() http.Handler {
	router := httprouter.New()
	configBankLoanContractRouter(router)
	configBankRepayPlanRouter(router)
	configLeaseContractRouter(router)
	configLeaseRepayPlanRouter(router)
	configLesseeInfoRouter(router)
	configShareholderLoanContractRouter(router)
	configShareholderLoanRepaidRecordRouter(router)

	router.GET("/ddl/:argID", GetDdl)
	router.GET("/ddl", GetDdlEndpoints)
	return router
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinBankLoanContractRouter(router)
	configGinBankRepayPlanRouter(router)
	configGinLeaseContractRouter(router)
	configGinLeaseRepayPlanRouter(router)
	configGinLesseeInfoRouter(router)
	configGinShareholderLoanContractRouter(router)
	configGinShareholderLoanRepaidRecordRouter(router)

	router.GET("/ddl/:argID", ConverHttprouterToGin(GetDdl))
	router.GET("/ddl", ConverHttprouterToGin(GetDdlEndpoints))
	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

// GetDdl is a function to get table info for a table in the fzzl database
// @Summary Get table info for a table in the fzzl database by argID
// @Tags TableInfo
// @ID argID
// @Description GetDdl is a function to get table info for a table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} api.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:8080/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		returnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	writeJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the fzzl database
// @Summary Gets a list of ddl endpoints available for tables in the fzzl database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the fzzl database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.CrudAPI
// @Router /ddl [get]
// http "http://localhost:8080/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "bank_loan_contract",
		CreateURL:       "/bankloancontract",
		RetrieveOneURL:  "/bankloancontract",
		RetrieveManyURL: "/bankloancontract",
		UpdateURL:       "/bankloancontract",
		DeleteURL:       "/bankloancontract",
		FetchDDLURL:     "/ddl/bank_loan_contract",
	}

	tmp.TableInfo, _ = model.GetTableInfo("bank_loan_contract")
	crudEndpoints["bank_loan_contract"] = tmp

	tmp = &CrudAPI{
		Name:            "bank_repay_plan",
		CreateURL:       "/bankrepayplan",
		RetrieveOneURL:  "/bankrepayplan",
		RetrieveManyURL: "/bankrepayplan",
		UpdateURL:       "/bankrepayplan",
		DeleteURL:       "/bankrepayplan",
		FetchDDLURL:     "/ddl/bank_repay_plan",
	}

	tmp.TableInfo, _ = model.GetTableInfo("bank_repay_plan")
	crudEndpoints["bank_repay_plan"] = tmp

	tmp = &CrudAPI{
		Name:            "lease_contract",
		CreateURL:       "/leasecontract",
		RetrieveOneURL:  "/leasecontract",
		RetrieveManyURL: "/leasecontract",
		UpdateURL:       "/leasecontract",
		DeleteURL:       "/leasecontract",
		FetchDDLURL:     "/ddl/lease_contract",
	}

	tmp.TableInfo, _ = model.GetTableInfo("lease_contract")
	crudEndpoints["lease_contract"] = tmp

	tmp = &CrudAPI{
		Name:            "lease_repay_plan",
		CreateURL:       "/leaserepayplan",
		RetrieveOneURL:  "/leaserepayplan",
		RetrieveManyURL: "/leaserepayplan",
		UpdateURL:       "/leaserepayplan",
		DeleteURL:       "/leaserepayplan",
		FetchDDLURL:     "/ddl/lease_repay_plan",
	}

	tmp.TableInfo, _ = model.GetTableInfo("lease_repay_plan")
	crudEndpoints["lease_repay_plan"] = tmp

	tmp = &CrudAPI{
		Name:            "lessee_info",
		CreateURL:       "/lesseeinfo",
		RetrieveOneURL:  "/lesseeinfo",
		RetrieveManyURL: "/lesseeinfo",
		UpdateURL:       "/lesseeinfo",
		DeleteURL:       "/lesseeinfo",
		FetchDDLURL:     "/ddl/lessee_info",
	}

	tmp.TableInfo, _ = model.GetTableInfo("lessee_info")
	crudEndpoints["lessee_info"] = tmp

	tmp = &CrudAPI{
		Name:            "shareholder_loan_contract",
		CreateURL:       "/shareholderloancontract",
		RetrieveOneURL:  "/shareholderloancontract",
		RetrieveManyURL: "/shareholderloancontract",
		UpdateURL:       "/shareholderloancontract",
		DeleteURL:       "/shareholderloancontract",
		FetchDDLURL:     "/ddl/shareholder_loan_contract",
	}

	tmp.TableInfo, _ = model.GetTableInfo("shareholder_loan_contract")
	crudEndpoints["shareholder_loan_contract"] = tmp

	tmp = &CrudAPI{
		Name:            "shareholder_loan_repaid_record",
		CreateURL:       "/shareholderloanrepaidrecord",
		RetrieveOneURL:  "/shareholderloanrepaidrecord",
		RetrieveManyURL: "/shareholderloanrepaidrecord",
		UpdateURL:       "/shareholderloanrepaidrecord",
		DeleteURL:       "/shareholderloanrepaidrecord",
		FetchDDLURL:     "/ddl/shareholder_loan_repaid_record",
	}

	tmp.TableInfo, _ = model.GetTableInfo("shareholder_loan_repaid_record")
	crudEndpoints["shareholder_loan_repaid_record"] = tmp

}
