package web

import (
	. "github.com/frezadev/webreactknot/library/core"
	. "github.com/frezadev/webreactknot/library/models"
	"github.com/frezadev/webreactknot/web/controller"
	"os"
	"path/filepath"
	"runtime"

	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

var (
	AppName  string = "web"
	basePath string = (func(dir string, err error) string { return dir }(os.Getwd()))

// 	server *knot.Server
)

func init() {
	app := knot.NewApp(AppName)
	app.ViewsPath = filepath.Join(basePath, AppName, "view") + toolkit.PathSeparator

	runtime.GOMAXPROCS(4)
	ConfigPath = controller.CONFIG_PATH
	// port := new(Ports)
	// port.ID = "port"
	// if err := port.GetPort(); err != nil {
	// 	toolkit.Printf("Error get port: %s \n", err.Error())
	// }
	// if port.Port == 0 {
	// 	if err := setup.Port(port); err != nil {
	// 		toolkit.Printf("Error set port: %s \n", err.Error())
	// 	}
	// }

	// server.Address = toolkit.Sprintf("localhost:%v", toolkit.ToString(port.Port))
	app.Static("res", filepath.Join(controller.AppBasePath, AppName, "assets"))
	app.Static("image", filepath.Join(controller.AppBasePath, AppName, "assets", "img"))
	app.Register(controller.CreatePageController(AppName))
	app.Register(controller.CreateLoginController())
	app.Register(controller.CreateAccessController())
	app.Register(controller.CreateSessionController())
	app.Register(controller.CreateUserController())
	app.Register(controller.CreateGroupController())
	app.Register(controller.CreateLogController())
	app.Register(controller.CreateDashboardController())

	if err := setAclDatabase(); err != nil {
		toolkit.Printf("Error set acl database : %s \n", err.Error())
	}
	knot.RegisterApp(app)
}

func setAclDatabase() error {
	if err := InitialSetDatabase(); err != nil {
		return err
	}

	if err := PrepareDefaultUser(); err != nil {
		return err
	}
	return nil
}
