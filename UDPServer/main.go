package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	HOST = "192.168.100.240"
	PORT = "41234"
	BATCH_CAPACITY = 1
	NAME_OF_ACTION = "Basketball"
	//URL = "http://192.168.100.233:8000/series"
	URL = "http://localhost:8008/series"
	TOKEN = "854362ef8cf672ccb4254cac3e867ae8d7110dd2"
	PAGE_SIZE = 10
	ACC_X = 0
	ACC_Y = 1
	ACC_Z = 2
	GYRO_X = 3
	GYRO_Y = 4
	GYRO_Z = 5
	TIME = 6
)

// TODO ########### БЫЛО ##########
JSON := `
{
	name: "Basketball"(string),
	data: 	{
				accX: [](float64),
				accY: [](float64),
				accZ: [](float64),
				gyroX: [](float64),
				gyroY: [](float64),
				gyroZ: [](float64),
				time: [](timestamp),
			},
}`

// TODO ########### СТАЛО ##########
JSON := `
{
	name: "Basketball"(string),
	data: 	[{
				accX: (float64),
				accY: (float64),
				accZ: (float64),
				gyroX: (float64),
				gyroY: (float64),
				gyroZ: (float64),
				time: (timestamp),
			}],
}`


type Batch struct {
	Dataname string `json:"name"`
	DataArray []Data `json:"data"`
}

type Data struct {
	AccX  float64  	`json:"accX"`
	AccY  float64  	`json:"accY"`
	AccZ  float64  	`json:"accZ"`
	GyroX float64  	`json:"gyroX"`
	GyroY float64  	`json:"gyroY"`
	GyroZ float64  	`json:"gyroZ"`
	Time  time.Time	`json:"time"`
}

var GLOBAL_BATCH = make([]Batch, PAGE_SIZE)

var GLOBAL_COUNTER = 0
var LOCAL_COUNTER = 0


func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 512)
	_, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	} else {

		//fmt.Println("UDP client: ", addr)
		//fmt.Println("Received from UDP client: ", string(buffer), )
		data := strings.Split(string(buffer), ";")


		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].AccX, err = strconv.ParseFloat(data[ACC_X], 64)
		if err != nil {
			fmt.Println(ACC_X)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].AccY, err = strconv.ParseFloat(data[ACC_Y], 64)
		if err != nil {
			fmt.Println(ACC_Y)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].AccZ, err = strconv.ParseFloat(data[ACC_Z], 64)
		if err != nil {
			fmt.Println(ACC_Z)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].GyroX, err = strconv.ParseFloat(data[GYRO_X], 64)
		if err != nil {
			fmt.Println(GYRO_X)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].GyroY, err = strconv.ParseFloat(data[GYRO_Y], 64)
		if err != nil {
			fmt.Println(GYRO_Y)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].GyroZ, err = strconv.ParseFloat(data[GYRO_Z], 64)
		if err != nil {
			fmt.Println(GYRO_Z)
			log.Fatal(err)
			return
		}
		GLOBAL_BATCH[GLOBAL_COUNTER].DataArray[LOCAL_COUNTER].Time = time.Now()
		//data[TIME] = strings.Replace(data[TIME], "\x00", "",-1)
		//GLOBAL_BATCH[GLOBAL_COUNTER].time[LOCAL_COUNTER], err = strconv.Atoi(data[TIME])
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		LOCAL_COUNTER++
		if LOCAL_COUNTER % 100 == 0 {
			fmt.Println(LOCAL_COUNTER)
		}
		if LOCAL_COUNTER == BATCH_CAPACITY {
			fmt.Println("SENDING DATA TO DB")
			old := GLOBAL_COUNTER
			go SendBatch(&GLOBAL_BATCH[old])
			if err != nil {
				log.Fatal("BATCH SENDING ERROR: ", err)
			}
			LOCAL_COUNTER = 0
			GLOBAL_COUNTER++
			if GLOBAL_COUNTER == PAGE_SIZE {
				GLOBAL_COUNTER = 0
			}
		}
	}
}

func initBatch(name string, capacity int) (d *Batch) {
	d = &Batch {
		Dataname: name,
		DataArray: make([]Data, capacity),
	}
	return d
}

func(d *Batch) mock() {
	for i := 0; i < len(d.DataArray); i++ {
		d.DataArray[i].AccX = rand.Float64()
		d.DataArray[i].AccY = rand.Float64()
		d.DataArray[i].AccZ = rand.Float64()

		d.DataArray[i].GyroX = rand.Float64()
		d.DataArray[i].GyroY = rand.Float64()
		d.DataArray[i].GyroZ = rand.Float64()
	}
}

func (d *Batch) clear() {
	d.DataArray = nil
}

func SendBatch(d *Batch) {
	client := http.Client{}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&d)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(d)
	//fmt.Println(bytes.NewBuffer(b))

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal("SENDING BATCH ERROR", err)
	}
	fmt.Println("SENDING BATCH")
	req.Header.Set("Authorization", "Token " + TOKEN)
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	d.clear()
	return
}

func main()  {
	for i := 0; i < len(GLOBAL_BATCH); i++ {
		GLOBAL_BATCH[i] = *initBatch(NAME_OF_ACTION, BATCH_CAPACITY)
		GLOBAL_BATCH[i].mock()
	}

	//udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", HOST, PORT))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//conn, err := net.ListenUDP("udp", udpAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//keepListening := true
	//for keepListening {
	//		handleUDPConnection(conn)
	//
	//}
	for i:=0; i < PAGE_SIZE; i++ {
		SendBatch(&GLOBAL_BATCH[i])
		time.Sleep(time.Second * 10)
	}
}
