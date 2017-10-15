## How to program from linux
1. build firmware hex file
2. push reset
3. type the following
```
avrdude -p atmega32u4 -c avr109 -U flash:w:./mf68_default.hex -P /dev/ttyACM0
```

