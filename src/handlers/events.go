package handlers

import (
	"github.com/faruq/restful-api-ppl/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	events   []models.Event
	nextID   int = 1
	eventMux sync.RWMutex
)

func GetAllEvents(c *gin.Context) {
	eventMux.RLock()
	defer eventMux.RUnlock()

	c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Data:   events,
	})
}

func GetEventByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid event ID",
		})
		return
	}

	eventMux.RLock()
	defer eventMux.RUnlock()

	for _, event := range events {
		if event.ID == id {
			c.JSON(http.StatusOK, models.Response{
				Status: "success",
				Data:   event,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Response{
		Status:  "error",
		Message: "Event not found",
	})
}

func CreateEvent(c *gin.Context) {
	var req models.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	if req.Title == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Title is required",
		})
		return
	}

	if req.Date == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Date is required",
		})
		return
	}

	eventMux.Lock()
	defer eventMux.Unlock()

	event := models.Event{
		ID:          nextID,
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
		Location:    req.Location,
		CreatedAt:   time.Now(),
	}
	nextID++
	events = append(events, event)

	c.JSON(http.StatusCreated, models.Response{
		Status: "success",
		Data:   event,
	})
}

func UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid event ID",
		})
		return
	}

	var req models.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	eventMux.Lock()
	defer eventMux.Unlock()

	for i, event := range events {
		if event.ID == id {
			if req.Title != "" {
				events[i].Title = req.Title
			}
			if req.Description != "" {
				events[i].Description = req.Description
			}
			if req.Date != "" {
				events[i].Date = req.Date
			}
			if req.Location != "" {
				events[i].Location = req.Location
			}

			c.JSON(http.StatusOK, models.Response{
				Status: "success",
				Data:   events[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Response{
		Status:  "error",
		Message: "Event not found",
	})
}

func DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid event ID",
		})
		return
	}

	eventMux.Lock()
	defer eventMux.Unlock()

	for i, event := range events {
		if event.ID == id {
			events = append(events[:i], events[i+1:]...)
			c.JSON(http.StatusOK, models.Response{
				Status:  "success",
				Message: "Event deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Response{
		Status:  "error",
		Message: "Event not found",
	})
}
