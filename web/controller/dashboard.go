package controller

import (
	// . "github.com/frezadev/webreactknot/library/core"
	// . "github.com/frezadev/webreactknot/library/helper"
	// . "github.com/frezadev/webreactknot/library/models"
	"github.com/frezadev/webreactknot/web/helper"

	// "strings"

	"github.com/eaciit/knot/knot.v1"
)

type DashboardController struct {
	App
}

func CreateDashboardController() *DashboardController {
	var controller = new(DashboardController)
	return controller
}

func (m *DashboardController) GetData(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	/*csr, e := DB().Connection.NewQuery().From("rpt_scadasummary").Cursor(nil)

	if e != nil {
		return helper.CreateResult(false, nil, e.Error())
	}
	defer csr.Close()

	data := []tk.M{}
	e = csr.Fetch(&data, 0, false)

	result := []string{}

	for _, val := range data {
		result = append(result, val.GetString("ID"))
	}
	sort.Strings(result)

	if e != nil {
		return helper.CreateResult(false, nil, e.Error())
	}

	return helper.CreateResult(true, result, "success")*/
	return helper.CreateResult(true, nil, "success")
}
