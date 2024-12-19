package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/event-booking/models"
)

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		fmt.Println(err)
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find event with given ID.",
		})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. Try again later!",
		})
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	event.UserId = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. Try again later!",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}
