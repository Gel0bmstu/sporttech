// библиотека для работы I²C
#include <Wire.h>
// библиотека для работы с модулями IMU
#include <TroykaIMU.h>

#include <ArduinoJson.h>

// создаём объект для работы с гироскопом
Gyroscope gyro;
// создаём объект для работы с акселерометром
Accelerometer accel;
// создаём объект для работы с компасом
Compass compass;
// создаём объект для работы с барометром
Barometer barometer;

// калибровочные значения компаса
// полученные в калибровочной матрице из примера «compassCalibrateMatrix»
const double compassCalibrationBias[3] = {
  524.21,
  3352.214,
  -1402.236
};

const double compassCalibrationMatrix[3][3] = {
  {1.757, 0.04, -0.028},
  {0.008, 1.767, -0.016},
  {-0.018, 0.077, 1.782}
};

const char* server = "";
int portNumber = 8080;

StaticJsonBuffer<200> jsonBuffer;
//JsonObject& root = jsonBuffer.createObject();
//JsonObject& data = root.createNestedObject("data");

void setup()
{
  Serial.begin(115200);
  gyro.begin();
  accel.begin();


}

auto zalupa()
{
  JsonObject& root = jsonBuffer.createObject();
  JsonArray& acc = jsonBuffer.createArray();
  acc.add(accel.readAX());
  acc.add(accel.readAY());
  acc.add(accel.readAZ());

  JsonArray& dus = jsonBuffer.createArray();
  dus.add(gyro.readDegPerSecX());
  dus.add(gyro.readDegPerSecY());
  dus.add(gyro.readDegPerSecZ());

  root.set(F("acc"),acc);
  root.set(F("gyro"),dus);
  String
  return
}
void loop()
{
  zalupa();

//  Serial.println();
}