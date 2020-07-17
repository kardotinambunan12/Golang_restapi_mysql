package main

import (
	"./customlogger"
	"./routes"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	//route index
	e := routes.Index()
	e.Validator = &CustomValidator{validator: validator.New()}

	logger := customlogger.GetInstance("SYSTEM")
	logger.Println("Starting Application")

	// http
	e.Logger.Fatal(e.Start(":4000"))
	// https
	// e.Logger.Fatal(e.StartAutoTLS(":443"))
}
