package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/event-booking/helpers"
	"github.com/kamil-budzik/event-booking/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId := helpers.GetEventId(ctx)

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register user for event.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registered",
	})

}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId := helpers.GetEventId(ctx)

	var event models.Event
	event.ID = eventId

	err := event.CancelRegistration(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled",
	})

}
