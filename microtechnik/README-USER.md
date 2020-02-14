###Начало работы
stty 115200 -F /dev/ttyUSB0 raw -echo  

echo "wifi '<SSID>' '<PASSWORD>'" >> /dev/ttyUSB0
echo "server <IP>:<PORT>" >> /dev/ttyUSB0


