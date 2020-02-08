#include <Wire.h>
#include <TroykaIMU.h>

Gyroscope gyro;
Accelerometer accel;

void setup() {
    Serial.begin(115200);

    gyro.begin();
    accel.begin();
}

String zalupa()
{
    char buffer[256];

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

void loop() {
    Serial.println(zalupa());
}