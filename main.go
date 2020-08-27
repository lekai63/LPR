package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                      // ui theme
	_ "github.com/GoAdminGroup/themes/sword"                         // ui theme 2

	// _ "github.com/lekai63/lpr/models/drivers/postgres"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"github.com/lekai63/lpr/lpr"
	"github.com/lekai63/lpr/models"
	"github.com/lekai63/lpr/pages"
	"github.com/lekai63/lpr/tables"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	// fmt.Printf("s", eng)
	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	// 增加tools页面
	eng.HTML("GET", "/admin/tools", pages.GetTools)

	r.POST("/admin/tools/", func(c *gin.Context) {
		res := make(gin.H)
		switch c.PostForm("calc") {
		case "insterestAll":
			// todo modify
			lpr.Icbc()
			res["code"] = 0
			res["msg"] = "成功调用"
			res["status_code"] = http.StatusOK
		default:
			res["code"] = 1
			res["msg"] = "调用失败"
			res["status_code"] = http.StatusBadRequest
		}

		c.JSON(http.StatusOK, res)
	})

	// tools分组

	// models.Init(eng.PostgresqlConnection())
	// models.Init()
	models.InitGormv2(eng.PostgresqlConnection())
	_ = r.Run(":8080")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.PostgresqlConnection().Close()
}
