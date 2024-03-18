package v1

import (
	"net/http"

	"github.com/dentist/api/models"
	"github.com/dentist/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// CreateAppointment ...
// @Summary CreateAppointment
// @Description Api for creat a new appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Param Appointment body models.ReqAppointment true "CreateAppointment"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointment [post]
func (h *handlerV1) CreateAppointment(c *gin.Context) {
	var req models.Appointment
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		h.logger.Error("Error to bind json create appointment")
		return
	}
	Id := uuid.NewString()
	response, err := h.storage.Appointment().CreateAppointment(&repo.Appointment{
		Id:          Id,
		ClientId:    req.ClientId,
		Date:        req.Date,
		Diagnostics: req.Diagnostics,
		Treatment:   req.Treatment,
		Amount:      req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create appointment",
		})
		h.logger.Error("Failed to create appointment")
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAppointment ...
// @Summary GetAppointment
// @Description Api for get a new appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointment [get]
func (h *handlerV1) GetAppointment(c *gin.Context) {
	id := c.Query("id")

	response, err := h.storage.Appointment().GetAppointment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get appointment",
		})
		h.logger.Error("Failed to get appointment")
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAppointment ...
// @Summary UpdateAppointment
// @Description Api for update appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Param Appointment body models.Appointment true "Appointment"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointment [put]
func (h *handlerV1) UpdateAppointment(c *gin.Context) {
	var req models.Appointment
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.storage.Appointment().UpdateAppointment(&repo.Appointment{
		Id:          req.Id,
		ClientId:    req.ClientId,
		Date:        req.Date,
		Diagnostics: req.Diagnostics,
		Treatment:   req.Treatment,
		Amount:      req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update appointment",
		})
		h.logger.Error("Failed to update appointment")
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAppointment ...
// @Summary DeleteAppointment
// @Description Api for delete appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} bool
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointment [delete]
func (h *handlerV1) DeleteAppointment(c *gin.Context) {
	id := c.Query("id")

	response, err := h.storage.Appointment().DeleteAppointment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete appointment",
		})
		h.logger.Error("Failed to delete appointment")
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllAppointments ...
// @Summary GetAllAppointments
// @Description Api for get all appointments
// @Tags appointment
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.AllAppointments
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointments [get]
func (h *handlerV1) GetAllAppointments(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	response, err := h.storage.Appointment().GetAllAppointments(&repo.GetAllAppointment{
		Page:  cast.ToInt(page),
		Limit: cast.ToInt(limit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get all appointments",
		})
		h.logger.Error("Failed to get all appointments")
		return
	}
	if len(response.Appointment) == 0 {
		c.JSON(http.StatusOK, models.AllAppointments{
			Appointments: []*models.Appointment{},
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAppointmentsWithDate ...
// @Summary GetAppointmentsWithDate
// @Description Api for get appointments with date
// @Tags appointment
// @Accept json
// @Produce json
// @Param integer query string true "integer"
// @Success 200 {object} models.AllAppointments
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointmentsdate [get]
func (h *handlerV1) GetAppointmentsWithDate(c *gin.Context) {
	integer := c.Query("integer")
	response, err := h.storage.Appointment().GetAppointmentsWithDate(cast.ToInt(integer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get appointments with date",
		})
		h.logger.Error("Failed to get appointments with date")
		return
	}
	if len(response.Appointment) == 0 {
		c.JSON(http.StatusOK, models.AllAppointments{
			Appointments: []*models.Appointment{},
		})
		return
	}

	c.JSON(http.StatusOK, response)
}


// GetAppointmentWithClientId...
// @Summary GetAppointmentWithClientId
// @Description Api for get appointment with client id
// @Tags appointment
// @Accept json
// @Produce json
// @Param client_id query string true "client_id"
// @Success 200 {object} models.AllAppointments
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointmentid [get]
func (h *handlerV1) GetAppointmentWithClientId(c *gin.Context) {
	client_id := c.Query("client_id")
	response, err := h.storage.Appointment().GetAppointmentsWithClientId(client_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to get appointment with client id",
		})
		h.logger.Error("Failed to get appointment with client id")
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateAppointmentWithClient ...
// @Summary CreateAppointmentWithClient
// @Description Api for creating a new client
// @Tags appointment
// @Accept json
// @Produce json
// @Param Client body models.ReqNew true "CreateAppointmentWithClient"
// @Success 200 {object} models.New
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/appointmentnew [post]
func (h *handlerV1) CreateAppointmentWithClient(c *gin.Context) {
	var req models.ReqNew
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		h.logger.Error("Failed to binding to json CreateAppointmentWithClient")
		return
	}
	Id := uuid.NewString()
	respClient, err := h.storage.Client().CreateClient(&repo.Client{
		Id:          Id,
		Name:        req.ClientName,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create client with appointment",
		})
		h.logger.Error("Failed to create client with appointment")
		return
	}

	id := uuid.NewString()

	respAppointment, err := h.storage.Appointment().CreateAppointment(&repo.Appointment{
		Id: id,
		ClientId: Id,
		Date: req.Date,
		Diagnostics: req.Diagnostics,
		Treatment: req.Treatment,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create appointment with client",
		})
		h.logger.Error("Failed to create appointment with client")
		return
	}

	response := models.New{
		AppointmentId: respAppointment.Id,
		ClientName: respClient.Name,
		PhoneNumber: respClient.PhoneNumber,
		ClientId: respAppointment.ClientId,
		Date: respAppointment.Date,
		Diagnostics: respAppointment.Diagnostics,
		Treatment: respAppointment.Treatment,
		Amount: respAppointment.Amount,
	}
	

	c.JSON(http.StatusCreated, response)
}
