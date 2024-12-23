package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/event-booking/helpers"
	"github.com/kamil-budzik/event-booking/models"
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
	userId := ctx.GetInt64("userId")

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	event.UserId = userId

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
	userId := ctx.GetInt64("userId")
	event := helpers.GetEventById(ctx, id)

	var updatedEvent models.Event
	err := ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
	}

	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event",
		})
		return
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
	userId := ctx.GetInt64("userId")
	event := helpers.GetEventById(ctx, id)

	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event",
		})
		return
	}

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
