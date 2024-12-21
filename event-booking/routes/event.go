package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/event-booking/helpers"
	"github.com/kamil-budzik/event-booking/models"
	"github.com/kamil-budzik/event-booking/utils"
)

func getEvent(ctx *gin.Context) {
	id := helpers.GetEventId(ctx)

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
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

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

func updateEvent(ctx *gin.Context) {
	id := helpers.GetEventId(ctx)
	event := helpers.GetEventById(ctx, id)
	_ = event

	var updatedEvent models.Event
	err := ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event. Try again later!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated!",
		"event":   updatedEvent,
	})

}

func deleteEvent(ctx *gin.Context) {
	id := helpers.GetEventId(ctx)
	event := helpers.GetEventById(ctx, id)

	err := event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete the event.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted!",
		"event":   event,
	})

}
