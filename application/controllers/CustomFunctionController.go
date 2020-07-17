package controllers

import (

	"github.com/labstack/echo"	
	logger    "../../customlogger"
)

// ## Define Config Variable Global
var logs 		  			= logger.GetInstance("SYSTEMS -")
// ## End

// ## Define Type Global
type response_json map[string]interface{}
type DB struct {
    Value        interface{}
    Error        error
    RowsAffected int64
}
func AppIndex(c echo.Context) error{
	return c.JSON(200, "High performance, minimalist Go web framework running")
}
// #