package routes

import (
	"go/by/example/restful/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(
			http.StatusBadGateway,
			gin.H{"message": "An error ocured during getting data fron db", "error": err},
		)
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadGateway,
			gin.H{"message": "An error ocured during parsing id", "error": err},
		)
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "An error ocured during fetching data"},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "An error ocured during cheking formats", "error": err},
		)
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "An error ocured during creating", "error": err},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"message": "Event Created!", "event": event},
	)
}

func updateEvent(context *gin.Context) {
	//1) Checking the correctness of id
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadGateway,
			gin.H{"message": "An error ocured during parsing id", "error": err},
		)
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "An error ocured during fetching data"},
		)
		return
	}

	//2) Checking if the format of object what we just got
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not paste request data", "error": err})
		return
	}

	//3) Calling update func
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not update event", "error": err})
		return
	}
	context.JSON(http.StatusOK,
		gin.H{"message": "Event Updated succsesfuly", "Id": eventId, "r": updateEvent})
	return

}

func deleteEvent(context *gin.Context) {

	//1)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		  context.JSON(
        http.StatusBadRequest,
        gin.H{"message": "Could not pass id", "error": err},
      )
    return
	}
  deleteEvent, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "An error ocured during fetching data"},
		)
		return
	}
  //2)
  err = deleteEvent.Delete()
  if err != nil {
	  context.JSON( 
      http.StatusBadRequest,
      gin.H{"message": "An error ocured during removing object", "error": err},
    )
    return
	}
	//3)
   context.JSON( 
      http.StatusAccepted,
      gin.H{"message": "Removing object sucsseded", "eventId": eventId},
    )

}
