package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8008"

)

type Batch struct {
	Dataname string `json:"name"`
	DataArray []Data `json:"data"`
}

type Data struct {
	Name  string 	`json:"name"`
	AccX  float64  	`json:"accX"`
	AccY  float64  	`json:"accY"`
	AccZ  float64  	`json:"accZ"`
	GyroX float64  	`json:"gyroX"`
	GyroY float64  	`json:"gyroY"`
	GyroZ float64  	`json:"gyroZ"`
	Time  time.Time	`json:"time"`
}



var GLOBAL_COUNTER = 0

func BatchTaker(c *gin.Context) {
	data := Batch{}
	_ = json.NewDecoder(c.Request.Body).Decode(&data)
	fmt.Println(data)
}

func main () {
	api := gin.New()
	api.POST("/series", BatchTaker)
	err := api.Run(fmt.Sprintf("%s:%s",HOST,PORT ))
	fmt.Println(err)
}
