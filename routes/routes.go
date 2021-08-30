package routes

import (
	"api-echo/modules/employee"
	"api-echo/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Main for setting routes
func Main(e *echo.Echo) {
	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			employeeHandler := employee.EmployeeController{}
			employee := v1.Group("/employee")
			{
				employee.GET("/index", employeeHandler.GetData)
				employee.POST("/create", employeeHandler.Create)
				employee.POST("/update/:id", employeeHandler.Update)
				employee.GET("/delete/:id", employeeHandler.Delete)
			}
		}
	}

	e.GET("/", func(ctx echo.Context) error {
		return ctx.HTML(http.StatusOK, "<div style='text-align:center'><b>WELCOME TO "+services.GetConfiguration("server.name")+"</b></div>")

	})
	e.GET("/swagger/*", services.WrapHandler)
}
