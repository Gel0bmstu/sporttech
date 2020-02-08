#include <HTTPClient.h>
#include <Wire.h>
#include <TroykaIMU.h>
#include <ArduinoJson.h>

// создаём объект для работы с гироскопом
Gyroscope gyro;
// создаём объект для работы с акселерометром
Accelerometer accel;


StaticJsonBuffer<200> jsonBuffer;
// вводим имя и пароль точки доступа
const char* ssid = "MGTS_30";
const char* password = "mgtswifi30";
   HTTPClient http;
void setup() {
    // иницилизируем монитор порта
    Serial.begin(115200);
    // запас времени на открытие монитора порта — 5 секунд
    delay(5000);
    // подключаемся к Wi-Fi сети
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
        delay(1000);
        Serial.println("Connecting to Wi-Fi..");
    }
    Serial.println("Connected to the Wi-Fi network");

  gyro.begin();
  accel.begin();
          http.begin("http://192.168.100.233:5000/data");
          http.addHeader("Content-Type", "application/json");
}
String zalupa()
{
//  JsonObject& root = jsonBuffer.createObject();
//  JsonArray& acc = jsonBuffer.createArray();
    char buffer[256];
    sprintf(buffer, "%f;%f;%f;%f;%f;%f;%d", accel.readAX(),accel.readAY(),accel.readAZ(),gyro.readDegPerSecX(),gyro.readDegPerSecY(),gyro.readDegPerSecZ(),millis());
//  acc.add(accel.readAX());
//  acc.add(accel.readAY());
//  acc.add(accel.readAZ());
//
//  JsonArray& dus = jsonBuffer.createArray();
//  dus.add(gyro.readDegPerSecX());
//  dus.add(gyro.readDegPerSecY());
//  dus.add(gyro.readDegPerSecZ());
//
//  root.set(F("acc"),acc);
//  root.set(F("gyro"),dus);

//  String  ans;
//  root.printTo(ans);
//  jsonBuffer.clear();
//  Serial.println(telemetry);
  return buffer;
}



void loop() {
//  http.addHeader("Content-Type", "application/json");
  int httpCode = http.POST(zalupa());
}