package v1

import (
	"log"
	"net/http"

	"github.com/dentist/api/models"
	"github.com/dentist/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// CreateClient ...
// @Summary CreateClient
// @Description Api for creating a new client
// @Tags client
// @Accept json
// @Produce json
// @Param Client body models.ReqClient true "CreateClient"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/client [post]
func (h *handlerV1) CreateClient(c *gin.Context) {
	var req models.Client
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	Id := uuid.NewString()
	response, err := h.storage.Client().CreateClient(&repo.Client{
		Id:            Id,
		Name:          req.Name,
		LastName:      req.LastName,
		FatherName:    req.FatherName,
		PhoneNumber:   req.PhoneNumber,
		Address:       req.Address,
		BirthDate:     req.BirthDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : "Failed to create client",
		})
		h.logger.Error("Failed to create client")
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

// GetClient
// @Summary GetClient
// @Description Api for get client
// @Tags client
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/client [get]
func (h *handlerV1) GetClient(c *gin.Context) {
	id := c.Query("id")

	response, err := h.storage.Client().GetClient(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to get client",
		})
		h.logger.Error("Failed to get client")
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateClient
// @Summary UpdateClient
// @Description Api for update client
// @Tags client
// @Accept json
// @Produce json
// @Param Client body models.Client true "UpdateClient"
// @Success 200 {object} models.Client
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/client [put]
func (h *handlerV1) UpdateClient(c *gin.Context) {
	var client models.Client
	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		h.logger.Error("Error binding json update client")
		return
	}

	response, err := h.storage.Client().UpdateClient(&repo.Client{
		Id: client.Id,
		Name: client.Name,
		LastName: client.LastName,
		FatherName: client.FatherName,
		PhoneNumber: client.PhoneNumber,
		Address: client.Address,
		BirthDate: client.BirthDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to update client",
		})
		h.logger.Error("Failed to update client")
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteClient
// @Summary DeleteClient
// @Description Api for delete client
// @Tags client
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} bool
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/client [delete]
func (h *handlerV1) DeleteClient(c *gin.Context) {
	id := c.Query("id")
	response, err := h.storage.Client().DeleteClient(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to delete client",
		})
		h.logger.Error("Failed to delete client")
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllClients
// @Summary GetAllClients
// @Description Api for get all clients
// @Tags client
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.AllClients
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/clients [get]
func (h *handlerV1) GetAllClients(c * gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	response, err := h.storage.Client().GetAllClients(&repo.GetAllClient{
		Page: cast.ToInt(page),
		Limit: cast.ToInt(limit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to get all clients",
		})
		h.logger.Error("Failed to get all clients")
		return
	}
	if len(response.Clients) == 0 {
		c.JSON(http.StatusOK, models.AllClients{
			Clients: []*models.Client{},
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllClientsCount
// @Summary GetAllClientsCount
// @Description Api for get all clients count
// @Tags client
// @Accept json
// @Produce json
// @Success 200 {object} int
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/count [get]
func (h *handlerV1) GetAllClientsCount(c *gin.Context) {
	resp, err := h.storage.Client().GetAllClientsCount()
	if err != nil {
		log.Println("Failed to get all clietns count")
		return 
	}
	
	c.JSON(http.StatusOK, resp)
}


// SearchingClients
// @Summary SearchingClients
// @Description Api for searching clients
// @Tags client
// @Accept json
// @Produce json
// @Param str query string true "SearchClients"
// @Success 200 {object} models.AllClients
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /v1/search [get]
func(h *handlerV1) SearchClients(c *gin.Context) {
	str := c.Query("str")
	response, err := h.storage.Client().SearchClients(str)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : "Failed to search clients",
		})
		log.Println("Failed to search clients", err)
		return
	}

	c.JSON(http.StatusOK, response)
}