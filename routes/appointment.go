package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spyder7370/SMH-Service/models"
)

func BookAppointmentRoute(c *gin.Context) {
	response := &models.Response{}
	request := models.AppointmentRequest{}
	body := c.Request.Body
	defer body.Close()
	requestBytes, err := io.ReadAll(body)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = &models.Error{
			Code:        "INVALID_REQUEST_BYTES",
			Description: err.Error(),
			StackTrace:  errors.WithStack(err).Error(),
		}
		c.JSON(response.Status, response)
		return
	}
	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = &models.Error{
			Code:        "INVALID_JSON_REQUEST",
			Description: err.Error(),
			StackTrace:  errors.WithStack(err).Error(),
		}
		c.JSON(response.Status, response)
		return
	}
}
func BookAppointment(appointmentRequest *models.AppointmentRequest) *models.Response {

	return nil
}

func InitRoutes() {
	httpServer := gin.Default()
	httpServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	httpServer.POST("/bookAppointment", BookAppointmentRoute)
	httpServer.Run()
}
