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

  if(sscanf(buf,"server '%[^']' %u",&server.ip,&server.port))
    EEPROM.put(SERVER_ADDR, server);

  EEPROM.commit();
}

void setup()
{
  Serial.begin(115200);
  EEPROM.begin(512);
  get_settings();
}

void loop()
{
  if(Serial.available())
  {
    parse_settings(Serial.readString());
  }
  Serial.println("---");
  print_settings();
  Serial.println("---");
  delay(1000);
}