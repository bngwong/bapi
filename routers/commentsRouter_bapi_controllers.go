package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["bapi/controllers:ReportController"] = append(beego.GlobalControllerRouter["bapi/controllers:ReportController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["bapi/controllers:ReportController"] = append(beego.GlobalControllerRouter["bapi/controllers:ReportController"],
		beego.ControllerComments{
			"Get",
			`/:reportId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["bapi/controllers:ReportController"] = append(beego.GlobalControllerRouter["bapi/controllers:ReportController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["bapi/controllers:ReportController"] = append(beego.GlobalControllerRouter["bapi/controllers:ReportController"],
		beego.ControllerComments{
			"Delete",
			`/:reportId`,
			[]string{"delete"},
			nil})

}
