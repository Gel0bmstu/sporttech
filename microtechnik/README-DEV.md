#Troyka-INU-master
В файле stmhw.h в строчках в самом начале изменить с вот этого:
Код (C++):
```
if defined(__AVR__) || defined(__SAMD21G18A__) || defined(ESP8266)
define WIRE_IMU Wire

elif defined(__SAM3X8E__) || defined(__SAM3A8C__) || defined(__SAM3A4C__)
define WIRE_IMU Wire1
endif
```
изменить на вот это:
Код (C++):
```
//#if defined(__AVR__) || defined(__SAMD21G18A__) || defined(ESP8266)
#define WIRE_IMU Wire

//#elif defined(__SAM3X8E__) || defined(__SAM3A8C__) || defined(__SAM3A4C__)
//#define WIRE_IMU Wire1
//#endif
```
