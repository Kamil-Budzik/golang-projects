package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/event-booking/models"
)

func GetEventId(ctx *gin.Context) int64 {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return 0
	}

	return id
}

func GetEventById(ctx *gin.Context, id int64) *models.Event {
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event",
			"err":     err,
		})
	}

	return event
}
