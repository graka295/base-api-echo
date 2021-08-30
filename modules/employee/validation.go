package employee

import (
	"api-echo/services"
	shared "api-echo/services/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// ValidationCreate function for validate Create return error, badrequest
func ValidationCreate(datareq *EmployeeAdd, c echo.Context) (bool, error) {
	if err := c.Bind(&datareq); err != nil {
		log.Println("[Error] Employee.ValidationCreate : ", err)
		return true, err
	}

	//validate
	errmessage, validate, err := services.Validate(datareq)
	if err != nil {
		log.Println("[Error] Employee.ValidationCreate : ", err)
		return true, err
	}
	if validate {
		response := shared.JSONResponse{
			Status:     "Failed",
			Validation: errmessage,
		}
		c.JSON(http.StatusBadRequest, response)
		return false, errors.New("Bad request")
	}

	errorMap := map[string]string{}
	errorValdiate := false
	// validate get by name
	_, err = GetByName(datareq.Nama)
	if !gorm.IsRecordNotFoundError(err) {
		statusMessage := services.Message("validateMessage.exist", map[string]interface{}{
			"Name": "Name",
		})
		errorMap["Name"] = statusMessage
		errorValdiate = true
	}
	if errorValdiate {
		response := shared.JSONResponse{
			Status:     "Failed",
			Validation: errorMap,
		}
		c.JSON(http.StatusBadRequest, response)
		return false, errors.New("Bad request")
	}

	return false, nil
}

// ValidationUpdate function for validate Create return error, badrequest
func ValidationUpdate(datareq *EmployeeAdd, c echo.Context) (bool, error) {
	if err := c.Bind(&datareq); err != nil {
		log.Println("[Error] Employee.ValidationUpdate : ", err)
		return true, err
	}
	//validate
	errmessage, validate, err := services.Validate(datareq)
	if err != nil {
		log.Println("[Error] Employee.ValidationUpdate : ", err)
		return true, err
	}
	if validate {
		response := shared.JSONResponse{
			Status:     "Failed",
			Validation: errmessage,
		}
		c.JSON(http.StatusBadRequest, response)
		return false, errors.New("Bad request")
	}

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	errorMap := map[string]string{}
	errorValdiate := false
	// validate get by name
	_, err = GetByID(idInt)
	if gorm.IsRecordNotFoundError(err) {
		statusMessage := services.Message("validateMessage.notfound", map[string]interface{}{
			"Name": "Pegawai",
		})
		errorMap["Name"] = statusMessage
		errorValdiate = true
	} else {
		dataName, err := GetByName(datareq.Nama)
		if !gorm.IsRecordNotFoundError(err) && dataName.ID != idInt {
			statusMessage := services.Message("validateMessage.exist", map[string]interface{}{
				"Name": "Name",
			})
			errorMap["Name"] = statusMessage
			errorValdiate = true
		}
	}
	if errorValdiate {
		response := shared.JSONResponse{
			Status:     "Failed",
			Validation: errorMap,
		}
		c.JSON(http.StatusBadRequest, response)
		return false, errors.New("Bad request")
	}

	return false, nil
}
