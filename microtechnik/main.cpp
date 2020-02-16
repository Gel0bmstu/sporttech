#include <WiFi.h>
#include <WiFiUdp.h>

#include <Wire.h>
#include <TroykaIMU.h>

#include <EEPROM.h>

struct {
  char ssid[20];
  char passwd[20];
} wifi;

struct {
  char ip[16];
  uint port;
} server;

uint WIFI_ADDR = 0;
uint SERVER_ADDR = sizeof(wifi);

void get_settings()
{
  EEPROM.get(WIFI_ADDR, wifi);
  EEPROM.get(SERVER_ADDR, server);
}

void put_settings()
{
  EEPROM.put(WIFI_ADDR, wifi);
  EEPROM.put(SERVER_ADDR, server);
  EEPROM.commit();
}

void print_wifi()
{
  EEPROM.get(WIFI_ADDR,wifi);
  Serial.println("WIFI:");
  Serial.println("ssid: "+String(wifi.ssid));
  Serial.println("password: "+String(wifi.passwd));
}

void print_server()
{
  EEPROM.get(SERVER_ADDR,server);
  Serial.println("SERVER:");
  Serial.println("ip: "+String(server.ip));
  Serial.println("port: "+String(server.port));
}

void print_settings()
{
  print_wifi();
  Serial.println();
  print_server();
}

bool parse_settings(String raw_settings)
{
  char buf[50];
  bool condition=false;
  raw_settings.toCharArray(buf, 50);

  if(sscanf(buf,"wifi '%[^']' '%[^']'",&wifi.ssid,&wifi.passwd))
    EEPROM.put(WIFI_ADDR, wifi);

  if(sscanf(buf,"server %[^:]:%u",&server.ip,&server.port))
    EEPROM.put(SERVER_ADDR, server);

  EEPROM.commit();
  connectToWiFi(wifi.ssid, wifi.passwd);
}

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
          udp.begin(WiFi.localIP(),server.port);
          connected = true;
          break;

      case SYSTEM_EVENT_STA_DISCONNECTED:
          Serial.println("WiFi lost connection");
          connected = false;
          break;

      default: break;
    }
}

void connectToWiFi(const char * ssid, const char * passwd){
  Serial.println("Connecting to WiFi network: " + String(ssid));

  WiFi.disconnect(true);
  WiFi.onEvent(WiFiEvent);
  WiFi.begin(ssid, passwd);

  Serial.println("Waiting for WIFI connection...");
}

void setup(){
  Serial.begin(115200);

  EEPROM.begin(512);
  get_settings();

  connectToWiFi(wifi.ssid, wifi.passwd);

  gyro.begin();
  accel.begin();
}

void loop(){
  if(connected){
    udp.beginPacket(server.ip, server.port);
    udp.printf(zalupa());
    udp.endPacket();
  }

    if(Serial.available())
    {
        parse_settings(Serial.readString());
        print_settings();
    }
}