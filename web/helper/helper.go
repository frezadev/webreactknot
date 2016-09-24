package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/toolkit"
)

var (
	DebugMode bool
)

type Payloads struct {
	Skip   int
	Take   int
	Sort   []Sorting
	Filter *FilterJS `json:"filter"`
	Misc   toolkit.M `json:"misc"`
}

type Sorting struct {
	Field string
	Dir   string
}

type FilterJS struct {
	Filters []*Filter `json:"filters"`
	Logic   string
}

type Filter struct {
	Field string      `json:"field"`
	Op    string      `json:"operator"`
	Value interface{} `json:"value"`
}

func (s *Payloads) ParseFilter() []*dbox.Filter {

	var filters []*dbox.Filter
	datelist := []string{
		"timestamp",
		"dateinfo.dateid",
		"startdate",
	}

	if s != nil {
		for _, each := range s.Filter.Filters {
			field := strings.ToLower(each.Field)
			switch each.Op {
			case "gte":
				var value interface{} = each.Value
				if toolkit.TypeName(value) == "string" {
					if value.(string) != "" {
						if toolkit.HasMember(datelist, field) {
							b, err := time.Parse("2006-01-02T15:04:05.000Z", value.(string))
							t, _ := time.Parse("2006-01-02 15:04:05", b.UTC().Format("2006-01-02")+" 00:00:00")
							if err != nil {
								fmt.Println(err.Error())
							} else {
								value = t
							}
						}
						filters = append(filters, dbox.Gte(field, value))
					}
				} else {
					filters = append(filters, dbox.Gte(field, value))
				}
			case "gt":
				var value interface{} = each.Value
				if toolkit.TypeName(value) == "string" {
					if value.(string) != "" {
						if toolkit.HasMember(datelist, field) {
							b, err := time.Parse("2006-01-02T15:04:05.000Z", value.(string))
							t, _ := time.Parse("2006-01-02 15:04:05", b.UTC().Format("2006-01-02")+" 00:00:00")
							if err != nil {
								fmt.Println(err.Error())
							} else {
								value = t
							}
						}
						filters = append(filters, dbox.Gt(field, value))
					}
				} else {
					filters = append(filters, dbox.Gt(field, value))
				}
			case "lte":
				var value interface{} = each.Value

				if toolkit.TypeName(value) == "string" {
					if value.(string) != "" {
						if toolkit.HasMember(datelist, field) {
							b, err := time.Parse("2006-01-02T15:04:05.000Z", value.(string))
							t, _ := time.Parse("2006-01-02 15:04:05", b.UTC().Format("2006-01-02")+" 23:59:59")
							if err != nil {
								fmt.Println(err.Error())
							} else {
								value = t
							}
						}
						filters = append(filters, dbox.Lte(field, value))
					}
				} else {
					filters = append(filters, dbox.Lte(field, value))
				}
			case "lt":
				var value interface{} = each.Value

				if toolkit.TypeName(value) == "string" {
					if value.(string) != "" {
						if toolkit.HasMember(datelist, field) {
							b, err := time.Parse("2006-01-02T15:04:05.000Z", value.(string))
							t, _ := time.Parse("2006-01-02 15:04:05", b.UTC().Format("2006-01-02")+" 23:59:59")
							if err != nil {
								fmt.Println(err.Error())
							} else {
								value = t
							}
						}
						filters = append(filters, dbox.Lt(field, value))
					}
				} else {
					filters = append(filters, dbox.Lt(field, value))
				}
			case "eq":
				value := each.Value

				if field == "turbine" && value.(string) == "" {
					continue
				} else if field == "isvalidtimeduration" && value.(bool) == true {
					continue
				} else if field == "projectid" && value.(string) == "" {
					continue
				}

				filters = append(filters, dbox.Eq(field, value))
			case "neq":
				value := each.Value
				filters = append(filters, dbox.Ne(field, value))
			case "in":
				value := each.Value
				if (field == "turbineid" && toolkit.SliceLen(value) == 0) ||
					field == "turbine" && toolkit.SliceLen(value) == 0 {
					continue
				}
				filters = append(filters, dbox.In(field, value.([]interface{})...))
			}
		}
	}

	return filters
}

func HandleError(err error, optionalArgs ...interface{}) bool {
	if err != nil {
		fmt.Printf("error occured: %s", err.Error())

		if len(optionalArgs) > 0 {
			optionalArgs[0].(func(bool))(false)
		}

		return false
	}

	if len(optionalArgs) > 0 {
		optionalArgs[0].(func(bool))(true)
	}

	return true
}

func CreateResult(success bool, data interface{}, message string) map[string]interface{} {
	if !success {
		fmt.Println("ERROR! ", message)
		if DebugMode {
			panic(message)
		}
	}

	return map[string]interface{}{
		"data":    data,
		"success": success,
		"message": message,
	}
}
