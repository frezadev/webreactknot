package controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eaciit/knot/knot.v1"
)

type App struct {
	Server *knot.Server
}

var (
	LayoutFile  = "layout.html"
	AppBasePath = func(dir string, err error) string { return dir }(os.Getwd())
	DATA_PATH   = filepath.Join(AppBasePath, "data")
	CONFIG_PATH = filepath.Join(AppBasePath, "config")
)

func init() {
	fmt.Println("Base Path ===> ", AppBasePath)

	if DATA_PATH != "" {
		fmt.Println("DATA_PATH ===> ", DATA_PATH)
		fmt.Println("CONFIG_PATH ===> ", CONFIG_PATH)
	}
}

type Payload struct {
	DateStart time.Time
	DateEnd   time.Time
	Project   string
	Turbine   string
	Skip      int
	Take      int
	// Sort                []Sorting
	Filter              toolkit.M
	IsValidTimeDuration bool
}