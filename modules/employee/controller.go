package employee

import (
	"log"
	"net/http"
	"strconv"

	shared "api-echo/services/models"

	"github.com/labstack/echo/v4"
)

// EmployeeController struct controller
type EmployeeController struct {
}

// GetData get list data employee
// @Tags API EMPLOYEE
// @Summary Request get data
// @Description Get data employee (Validation: Need testing)
// @Accept json
// @Produce json
// @Success 200 {object} models.JSONResponse{Data=Employee{}}
// @Failure 400 {object} models.JSONResponse{}
// @Failure 401 {object} models.JSONResponse{}
// @Failure 404 {object} models.JSONResponse{}
// @Failure 500 {object} models.JSONResponse{}
// @Router /api/v1/employee/index [get]
func (ctx *EmployeeController) GetData(c echo.Context) error {
	data, err := GetAll()
	if err != nil {
		log.Println("[Error] EmployeeController.GetData : ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	response := shared.JSONResponse{
		Status:  "Success",
		Data:    data,
		Message: "Success get employee",
	}
	return c.JSON(http.StatusOK, response)
}

// Create create employee
// @Tags API EMPLOYEE
// @Summary Request for create data
// @Description Create data employee (Validation: Need testing)
// @Accept json
// @Param body body EmployeeAdd true "Body"
// @Produce json
// @Success 200 {object} models.JSONResponse{}
// @Failure 400 {object} models.JSONResponse{}
// @Failure 401 {object} models.JSONResponse{}
// @Failure 404 {object} models.JSONResponse{}
// @Failure 500 {object} models.JSONResponse{}
// @Router /api/v1/employee/create [post]
func (ctx *EmployeeController) Create(c echo.Context) error {
	dataReq := EmployeeAdd{}
	internalerror, err := ValidationCreate(&dataReq, c)
	if err != nil {
		if internalerror {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		log.Println("[Error] EmployeeController.Create : ", err)
		return nil
	}
	_, err = Create(dataReq)
	if err != nil {
		log.Println("[Error] EmployeeController.Create : ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	response := shared.JSONResponse{
		Status:  "Success",
		Message: "Success create employee",
	}
	return c.JSON(http.StatusOK, response)
}

// Update edit employe
// @Tags API EMPLOYEE
// @Summary Request for update data
// @Description Update data employee (Validation: Need testing)
// @Accept json
// @Param id path int true "id"
// @Param body body EmployeeAdd true "Body"
// @Produce json
// @Success 200 {object} models.JSONResponse{}
// @Failure 400 {object} models.JSONResponse{}
// @Failure 401 {object} models.JSONResponse{}
// @Failure 404 {object} models.JSONResponse{}
// @Failure 500 {object} models.JSONResponse{}
// @Router /api/v1/employee/update/{id} [post]
func (ctx *EmployeeController) Update(c echo.Context) error {
	dataReq := EmployeeAdd{}
	internalerror, err := ValidationUpdate(&dataReq, c)
	if err != nil {
		if internalerror {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		log.Println("[Error] EmployeeController.Update : ", err)
		return nil
	}
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	_, err = Update(dataReq, idInt)
	if err != nil {
		log.Println("[Error] EmployeeController.Update : ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	response := shared.JSONResponse{
		Status:  "Success",
		Message: "Success update employee",
	}
	return c.JSON(http.StatusOK, response)
}

// Delete delete employee
// @Tags API EMPLOYEE
// @Summary Request for delete data
// @Description Delete data employee (Validation: Need testing)
// @Accept json
// @Param id path int true "id"
// @Produce json
// @Success 200 {object} models.JSONResponse{}
// @Failure 400 {object} models.JSONResponse{}
// @Failure 401 {object} models.JSONResponse{}
// @Failure 404 {object} models.JSONResponse{}
// @Failure 500 {object} models.JSONResponse{}
// @Router /api/v1/employee/delete/{id} [get]
func (ctx *EmployeeController) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := Delete(idInt)
	if err != nil {
		log.Println("[Error] EmployeeController.Delete : ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	response := shared.JSONResponse{
		Status:  "Success",
		Message: "Success delete employee",
	}
	return c.JSON(http.StatusOK, response)
}
