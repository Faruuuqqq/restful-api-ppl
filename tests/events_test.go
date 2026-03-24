package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faruq/restful-api-ppl/src/handlers"
	"github.com/faruq/restful-api-ppl/src/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	api := router.Group("/api")
	{
		events := api.Group("/events")
		{
			events.GET("", handlers.GetAllEvents)
			events.GET("/:id", handlers.GetEventByID)
			events.POST("", handlers.CreateEvent)
			events.PUT("/:id", handlers.UpdateEvent)
			events.DELETE("/:id", handlers.DeleteEvent)
		}
	}

	return router
}

func TestGetAllEvents(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/events", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response.Status)
}

func TestGetEventByID(t *testing.T) {
	router := setupRouter()

	createEvent(router)

	req, _ := http.NewRequest("GET", "/api/events/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response.Status)
}

func TestGetEventByIDNotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/events/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateEvent(t *testing.T) {
	router := setupRouter()

	event := models.EventRequest{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        "2026-04-15",
		Location:    "Test Location",
	}
	body, _ := json.Marshal(event)

	req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Response
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response.Status)
}

func TestCreateEventMissingTitle(t *testing.T) {
	router := setupRouter()

	event := models.EventRequest{
		Description: "Test Description",
		Date:        "2026-04-15",
		Location:    "Test Location",
	}
	body, _ := json.Marshal(event)

	req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateEvent(t *testing.T) {
	router := setupRouter()

	createEvent(router)

	event := models.EventRequest{
		Title: "Updated Event",
	}
	body, _ := json.Marshal(event)

	req, _ := http.NewRequest("PUT", "/api/events/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response.Status)
}

func TestDeleteEvent(t *testing.T) {
	router := setupRouter()

	createEvent(router)

	req, _ := http.NewRequest("DELETE", "/api/events/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response.Status)
}

func createEvent(router *gin.Engine) {
	event := models.EventRequest{
		Title:       "Test Event",
		Description: "Test Description",
		Date:        "2026-04-15",
		Location:    "Test Location",
	}
	body, _ := json.Marshal(event)

	req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}
