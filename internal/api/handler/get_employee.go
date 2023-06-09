package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
)

// GetEmployee implements openapi.ServerInterface
func (h *Handlers) GetEmployee(c *gin.Context, params openapi.GetEmployeeParams) {
	employee, err := h.staffApp.GetEmployee(c, params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := openapi.GetEmployeeResponse{
		ID:      employee.ID,
		Name:    employee.Name,
		PhotoID: employee.PhotoID,
		Meta: openapi.Meta{
			AdditionalProperties: employee.Meta,
		},
	}

	c.JSON(http.StatusOK, resp)
}
