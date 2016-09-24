package controller

import (
	. "github.com/frezadev/webreactknot/library/models"

	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type PageController struct {
	App
	Params toolkit.M
}

func CreatePageController(AppName string) *PageController {
	var controller = new(PageController)
	controller.Params = toolkit.M{"AppName": AppName}
	return controller
}

func (w *PageController) GetParams() toolkit.M {
	w.Params.Set("AntiCache", toolkit.RandomString(20))
	return w.Params
}

func (w *PageController) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-index.html"

	return w.GetParams()
}

func (w *PageController) Login(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.ViewName = "page-login.html"

	return w.GetParams()
}

func (w *PageController) Access(r *knot.WebContext) interface{} {
	WriteLog(r.Session("sessionid", ""), "access", r.Request.URL.String())
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-access.html"

	return w.GetParams()
}

func (w *PageController) Group(r *knot.WebContext) interface{} {
	WriteLog(r.Session("sessionid", ""), "access", r.Request.URL.String())
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-group.html"

	return w.GetParams()
}

func (w *PageController) Session(r *knot.WebContext) interface{} {
	WriteLog(r.Session("sessionid", ""), "access", r.Request.URL.String())
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-session.html"

	return w.GetParams()
}

func (w *PageController) Log(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-log.html"

	return w.GetParams()
}

func (w *PageController) User(r *knot.WebContext) interface{} {
	WriteLog(r.Session("sessionid", ""), "access", r.Request.URL.String())
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-user.html"

	return w.GetParams()
}

func (w *PageController) Dashboard(r *knot.WebContext) interface{} {
	WriteLog(r.Session("sessionid", ""), "access", r.Request.URL.String())
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = LayoutFile
	r.Config.ViewName = "page-dashboard.html"

	return w.GetParams()
}
