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

func configLesseeInfoRouter(router *httprouter.Router) {
	router.GET("/lesseeinfo", GetAllLesseeInfo)
	router.POST("/lesseeinfo", AddLesseeInfo)
	router.GET("/lesseeinfo/:argId", GetLesseeInfo)
	router.PUT("/lesseeinfo/:argId", UpdateLesseeInfo)
	router.DELETE("/lesseeinfo/:argId", DeleteLesseeInfo)
}

func configGinLesseeInfoRouter(router gin.IRoutes) {
	router.GET("/lesseeinfo", ConverHttprouterToGin(GetAllLesseeInfo))
	router.POST("/lesseeinfo", ConverHttprouterToGin(AddLesseeInfo))
	router.GET("/lesseeinfo/:argId", ConverHttprouterToGin(GetLesseeInfo))
	router.PUT("/lesseeinfo/:argId", ConverHttprouterToGin(UpdateLesseeInfo))
	router.DELETE("/lesseeinfo/:argId", ConverHttprouterToGin(DeleteLesseeInfo))
}

// GetAllLesseeInfo is a function to get a slice of record(s) from lessee_info table in the fzzl database
// @Summary Get list of LesseeInfo
// @Tags LesseeInfo
// @Description GetAllLesseeInfo is a handler to get a slice of record(s) from lessee_info table in the fzzl database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.LesseeInfo}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /lesseeinfo [get]
// http "http://localhost:8080/lesseeinfo?page=0&pagesize=20" X-Api-User:user123
func GetAllLesseeInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "lessee_info", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLesseeInfo(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLesseeInfo is a function to get a single record from the lessee_info table in the fzzl database
// @Summary Get record from table LesseeInfo by  argId
// @Tags LesseeInfo
// @ID argId
// @Description GetLesseeInfo is a function to get a single record from the lessee_info table in the fzzl database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.LesseeInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /lesseeinfo/{argId} [get]
// http "http://localhost:8080/lesseeinfo/1" X-Api-User:user123
func GetLesseeInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lessee_info", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLesseeInfo(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLesseeInfo add to add a single record to lessee_info table in the fzzl database
// @Summary Add an record to lessee_info table
// @Description add to add a single record to lessee_info table in the fzzl database
// @Tags LesseeInfo
// @Accept  json
// @Produce  json
// @Param LesseeInfo body model.LesseeInfo true "Add LesseeInfo"
// @Success 200 {object} model.LesseeInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /lesseeinfo [post]
// echo '{"business_license": "ihwaAiJCTJWVjMbPCekNJuKQX","lessee": "QfsIAhXYaegccuQDSUAVQUHvW","short_name": "cMKbLWDTLDFutYXyIqDabRESC","email": "mDleomEDHngLEJdLqjxODCjVN","contact_person": "PoMNgJBmHJQcBomaCFsIYTtOT","contact_tel": "LBXucnnUBGfLZSrjquJnqLCRi","risk_manager": "FvjnCGuYAPnGrbdBGvBEeOfID","id": 39,"customer_manager": "USjonDdCTXwSRIIZFqxqxoiJI","created_at": "2106-02-28T16:54:26.994308429Z","updated_at": "2275-08-12T05:23:30.253323103Z"}' | http POST "http://localhost:8080/lesseeinfo" X-Api-User:user123
func AddLesseeInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	lesseeinfo := &model.LesseeInfo{}

	if err := readJSON(r, lesseeinfo); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := lesseeinfo.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	lesseeinfo.Prepare()

	if err := lesseeinfo.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lessee_info", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	lesseeinfo, _, err = dao.AddLesseeInfo(ctx, lesseeinfo)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, lesseeinfo)
}

// UpdateLesseeInfo Update a single record from lessee_info table in the fzzl database
// @Summary Update an record in table lessee_info
// @Description Update a single record from lessee_info table in the fzzl database
// @Tags LesseeInfo
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  LesseeInfo body model.LesseeInfo true "Update LesseeInfo record"
// @Success 200 {object} model.LesseeInfo
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /lesseeinfo/{argId} [patch]
// echo '{"business_license": "ihwaAiJCTJWVjMbPCekNJuKQX","lessee": "QfsIAhXYaegccuQDSUAVQUHvW","short_name": "cMKbLWDTLDFutYXyIqDabRESC","email": "mDleomEDHngLEJdLqjxODCjVN","contact_person": "PoMNgJBmHJQcBomaCFsIYTtOT","contact_tel": "LBXucnnUBGfLZSrjquJnqLCRi","risk_manager": "FvjnCGuYAPnGrbdBGvBEeOfID","id": 39,"customer_manager": "USjonDdCTXwSRIIZFqxqxoiJI","created_at": "2106-02-28T16:54:26.994308429Z","updated_at": "2275-08-12T05:23:30.253323103Z"}' | http PUT "http://localhost:8080/lesseeinfo/1"  X-Api-User:user123
func UpdateLesseeInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	lesseeinfo := &model.LesseeInfo{}
	if err := readJSON(r, lesseeinfo); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := lesseeinfo.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	lesseeinfo.Prepare()

	if err := lesseeinfo.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "lessee_info", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	lesseeinfo, _, err = dao.UpdateLesseeInfo(ctx,
		argId,
		lesseeinfo)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, lesseeinfo)
}

// DeleteLesseeInfo Delete a single record from lessee_info table in the fzzl database
// @Summary Delete a record from lessee_info
// @Description Delete a single record from lessee_info table in the fzzl database
// @Tags LesseeInfo
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.LesseeInfo
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /lesseeinfo/{argId} [delete]
// http DELETE "http://localhost:8080/lesseeinfo/1" X-Api-User:user123
func DeleteLesseeInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseInt32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "lessee_info", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLesseeInfo(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
