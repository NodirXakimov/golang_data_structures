package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var countOfRequests int

const (
	token = "588931888:AAGNi2aQTidr27cMFThfAaimKiTOn7HF5Q0"
)

func main() {
	readyURL := fmt.Sprintf("https://api.telegram.org/bot%v/%s", token, "getUpdates")
	resp, err := http.Get(readyURL)
	if err != nil {
		log.Fatalln("Error on getting data from API: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))
	fmt.Printf("Type of the body: %T\n", body)

	r := gin.Default()
	r.GET("/", sendMessage)
	r.Run()
}

func sendMessage(c *gin.Context) {
	countOfRequests++
	message := fmt.Sprintf("Hello every one in group: %v", countOfRequests)
	message = fmt.Sprintf("https://api.telegram.org/bot588931888:AAGNi2aQTidr27cMFThfAaimKiTOn7HF5Q0/sendMessage?chat_id=-1001608326886&text=%s", message)
	resp, err := http.Get(message)
	if err != nil {
		log.Fatalln("Error on the sending message to the group: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":           "message sent",
		"response":          resp.Body,
		"count_of_requests": countOfRequests,
	})
	return
}
