## How to program from linux
1. build firmware hex file
1. push reset
1. type the following
```
avrdude -p atmega32u4 -c avr109 -U flash:w:./mf68_default.hex -P /dev/ttyACM0
```
## Application Notes

### Board Errors

  1. 10/30/2016 v1.1 the space bar LEDs are reversed polarity
  1. there's little or no documentation on the correct resistor Ohms. Use the guidelines from winkeyless.kr based on the color. If you use an online LED resistor calculator it will recommend resistors that will blow the 500mA USB keyboard power budget. For example, I used 910 Ohm resistors with green 1.8mm LEDs (3.0-3.2 fV @ 20mA) and they will only use about 120-150mA to power 69 LEDs, the maximum.
