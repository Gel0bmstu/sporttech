#include <WiFi.h>
#include <WiFiUdp.h>

#include <Wire.h>
#include <TroykaIMU.h>

const char * networkName = "MGTS_30";
const char * networkPswd = "mgtswifi30";

const char * udpAddress = "192.168.100.240";
const int udpPort = 41234;

boolean connected = false;

WiFiUDP udp;

Gyroscope gyro;
Accelerometer accel;

char buffer[256];
char* zalupa()
{
    sprintf(buffer,
    "%f;%f;%f;%f;%f;%f;%d",
    accel.readAX(),
    accel.readAY(),
    accel.readAZ(),
    gyro.readDegPerSecX(),
    gyro.readDegPerSecY(),
    gyro.readDegPerSecZ(),
    millis());

    return buffer;
}

void WiFiEvent(WiFiEvent_t event){
    switch(event) {

      case SYSTEM_EVENT_STA_GOT_IP:
          Serial.print("WiFi connected! IP address: ");
          Serial.println(WiFi.localIP());
          udp.begin(WiFi.localIP(),udpPort);
          connected = true;
          break;

      case SYSTEM_EVENT_STA_DISCONNECTED:
          Serial.println("WiFi lost connection");
          connected = false;
          break;

      default: break;
    }
}

void connectToWiFi(const char * ssid, const char * pwd){
  Serial.println("Connecting to WiFi network: " + String(ssid));

  WiFi.disconnect(true);
  WiFi.onEvent(WiFiEvent);
  WiFi.begin(ssid, pwd);

  Serial.println("Waiting for WIFI connection...");
}

void setup(){
  Serial.begin(115200);

  connectToWiFi(networkName, networkPswd);

  gyro.begin();
  accel.begin();
}

void loop(){
  if(connected){
    udp.beginPacket(udpAddress,udpPort);
    udp.printf(zalupa());
    udp.endPacket();
  }
}
