package main

import (
	. "github.com/frezadev/webreactknot/library/core"
	web "github.com/frezadev/webreactknot/web"
	. "github.com/frezadev/webreactknot/web/controller"
	"fmt"
	"net/http"

	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

func main() {
	config := new(Configuration)

	config.ID = "port"
	port, err := config.GetPort()
	if err != nil {
		toolkit.Printf("Error get port : %s \n", err.Error())
	}

	ConfigPath = CONFIG_PATH

	flagAddress := toolkit.Sprintf("localhost:%v", toolkit.ToString(port))
	appConf := knot.AppContainerConfig{Address: flagAddress}
	otherRoutes := map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			prefix := web.AppName

			sessionid := r.Session("sessionid", "")
			if sessionid == "" {
				http.Redirect(r.Writer, r.Request, fmt.Sprintf("/%s/page/login", prefix), 301)
			} else {
				http.Redirect(r.Writer, r.Request, fmt.Sprintf("/%s/page/dashboard", prefix), 301)
			}

			return true
		},
	}

	knot.DefaultOutputType = knot.OutputTemplate
	knot.StartContainerWithFn(&appConf, otherRoutes)
}
