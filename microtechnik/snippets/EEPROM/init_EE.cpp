 #include <EEPROM.h>

struct {
  char ssid[20] = "HUAWEI";
  char passwd[20] = "1234567890";
} wifi;

struct {
  char ip[16] = "0.0.0.0";
  uint port = 8080;
} server;

uint WIFI_ADDR = 0;
uint SERVER_ADDR = sizeof(wifi);

void setup()
{
  Serial.begin(115200);

  EEPROM.begin(512); //???????

  // replace values in EEPROM
  EEPROM.put(WIFI_ADDR, wifi);
  EEPROM.put(SERVER_ADDR, server);
  EEPROM.commit();

  // reload data for EEPROM, see the change
  EEPROM.get(WIFI_ADDR,wifi);
  Serial.println("WIFI:");
  Serial.println("ssid: "+String(wifi.ssid));
  Serial.println("password: "+String(wifi.passwd));

  Serial.println();

  EEPROM.get(SERVER_ADDR,server);
  Serial.println("SERVER:");
  Serial.println("ssid: "+String(server.ip));
  Serial.println("port: "+String(server.port));
}

void loop()
{
  delay(1000);
}