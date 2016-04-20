package controllers

import (
	"bapi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about report
type ReportController struct {
	beego.Controller
}

// @Title create
// @Description create report
// @Param	body		body 	models.Report	true		"The report content"
// @Success 200 {string} models.Report.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ReportController) Post() {
	var ob models.ReportObject
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)

	reportid := models.AddOne(ob)

	o.Data["json"] = map[string]string{"ReportId": reportid}
	o.ServeJSON()
}

// @Title Get
// @Description find report by reportid
// @Param	reportId		path 	string	true		"the reportid you want to get"
// @Success 200 {report} models.Report
// @Failure 403 :reportId is empty
// @router /:reportId [get]
func (o *ReportController) Get() {
	reportId := o.Ctx.Input.Param(":reportId")
	//fmt.Println("reportId:", reportId)
	if reportId != "" {
		ob, err := models.GetOne(reportId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all reports
// @Success 200 {report} models.Report
// @Failure 403 :reportId is empty
// @router / [get]
func (o *ReportController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title delete
// @Description delete the report
// @Param	reportId		path 	string	true		"The reportId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 reportId is empty
// @router /:reportId [delete]
func (o *ReportController) Delete() {
	objectId := o.Ctx.Input.Param(":reportId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
