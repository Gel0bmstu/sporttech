##Настройка ArduinoIDE
__Файл->Настройки->Дополнительные ссылки для менеджера плат:__
https://dl.espressif.com/dl/package_esp32_index.json  

__Менеджер плат:__  
DOIT ESP32 DEVKIT V1

##Библиотеки
microtechnik/libs  

Просто в файле stmhw.h в строчках в самом начале изменить с вот этого:
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
