package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Activity struct {
	Activity string `json:"activity"`
}

func InitializeHandlers(router *gin.Engine) {
	router.GET("/name/:name", getActivity)
}

func getActivity(c *gin.Context) {
	name := c.Param("name")
	response, err := getName(name)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return
	}
	c.JSON(http.StatusOK, response.Activity)
}

func getName(name string) (*Activity, error) {
	url := "https://www.boredapi.com/api/activity"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var activity Activity
	if err := json.NewDecoder(response.Body).Decode(&activity); err != nil {
		return nil, err
	}
	activity.Activity = name + ", " + activity.Activity

	return &activity, nil
}

func main() {

	router := gin.Default()
	InitializeHandlers(router)
	router.Run(":8080")

}
