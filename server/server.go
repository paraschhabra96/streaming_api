package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	router := gin.Default()

	//Serve client
	router.LoadHTMLGlob("./webapp/*.html")
	router.Static("/webapp","./webapp")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	//Setup Streaming API
	router.GET("/api/users/stream", getUsers)

	router.Run("localhost:8080")
}

func init() {
	//Load mock data for api response
	data, err := ioutil.ReadFile("./chunks.json")
	if err != nil {
		log.Fatalf("Init failed : %v",err)
	}

	var result []ResponseData
	err = json.Unmarshal(data,&result)
	if err != nil {
		log.Fatalf("Init failed : %v",err)
	}
	mockPayloads = result
}

type ResponseData struct {
	Data map[string]interface{} `json:"data"`
	Delay int `json:"delayMs"`
}

var mockPayloads []ResponseData

// getUsers responds with the list of all the users
func getUsers(c *gin.Context) {

	responseChannel := make(chan map[string]interface{})
	go func() {
		for _,payload := range mockPayloads {
			time.Sleep(time.Duration(payload.Delay)*time.Millisecond)
			responseChannel <- payload.Data
		}
		close(responseChannel)
	}()

	c.Stream(func(w io.Writer) bool {
		data, ok := <- responseChannel
		if ok {
			c.SSEvent("data",data)
		} else {
			c.SSEvent("end",data)
		}
		return ok
	})

}