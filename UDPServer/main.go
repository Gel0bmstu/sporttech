package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (

	HOST_IP_ENV = "HOST_IP"


	PORT_ENV = "PORT_ENV"
	URL_ENV = "URL_ENV"
	BATCH_CAP_ENV = "BATCH_CAP_ENV"
	NAME_OF_ACTION_ENV = "NAME_OF_ACTION_ENV"

	//URL = "http://192.168.100.233:8000/series"

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


type Config struct {
	HOST_IP string
	PORT string
	URL string
	BATCH_CAPACITY int
	NAME_OF_ACTION string
}


type Batch struct {
	DataName string `json:"name"`
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
	//fmt.Println("UDP client: ", addr)
	//fmt.Println("Received from UDP client: ", string(buffer))
	//return
	if err != nil {
		log.Fatal(err)
	} else {


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
		if LOCAL_COUNTER == conf.BATCH_CAPACITY {
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
		DataName: name,
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

	req, err := http.NewRequest("POST", conf.URL, bytes.NewBuffer(b))
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
	d.DataArray = make([]Data, conf.BATCH_CAPACITY)
	return
}

var conf Config

func main()  {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Oops: " + err.Error() + "\n")
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				conf.HOST_IP = ipnet.IP.String()
				break
			}
		}
	}
	exist := false

	//conf.HOST_IP, exist = os.LookupEnv(HOST_IP_ENV)
	//if !exist {
	//	log.Fatal("NOT FOUND HOST IP")
	//}
	conf.PORT, exist = os.LookupEnv(PORT_ENV)
	if !exist {
		log.Fatal("NOT FOUND PORT")
	}
	BATCH_CAPACITY, exist := os.LookupEnv(BATCH_CAP_ENV)
	if !exist {
		log.Fatal("NOT FOUND BATCH CAPACITY")
	}
	conf.BATCH_CAPACITY, _ = strconv.Atoi(BATCH_CAPACITY)
	conf.URL, exist = os.LookupEnv(URL_ENV)
	if !exist {
		log.Fatal("NOT FOUND URL")
	}
	conf.NAME_OF_ACTION, exist = os.LookupEnv(NAME_OF_ACTION_ENV)
	if !exist {
		log.Fatal("NOT FOUND NAME_OF_ACTION_ENV")
	}





	//for i := 0; i < len(GLOBAL_BATCH); i++ {
	//	GLOBAL_BATCH[i] = *initBatch(conf.NAME_OF_ACTION, conf.BATCH_CAPACITY)
	//	GLOBAL_BATCH[i].mock()
	//}

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", conf.HOST_IP, conf.PORT))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UDP SERVER IS READY " + fmt.Sprintf("%s:%s", conf.HOST_IP, conf.PORT))
	keepListening := true
	for keepListening {
			handleUDPConnection(conn)

	}
	//for i:=0; i < PAGE_SIZE; i++ {
	//	SendBatch(&GLOBAL_BATCH[i])
	//	time.Sleep(time.Second * 10)
	//}
}
