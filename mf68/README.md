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
  1. there's little or no documentation on the correct resistor Ohms. Use the guidelines from winkeyless.kr based on the color. If you use an online LED resistor calculator it will recommend resistors that will blow the 500mA USB keyboard power budget. For example, I used 910 Ohm resistors with green 1.8mm LEDs (3.0-3.2 fV @ 20mA) and they will only use about 120-150mA to power 70 LEDs, the maximum.
  
### Assembly notes

  1. the PCB and the Plate have a curious relationship space wise. Be very careful when inserting the switches that they don't unseat their neighbors. Use the Wodan style technique where you solder in a single leg, preferrably the strong one and the carefully hold the plate and PCB together, reheat the leg and push them plate, switch and PCB together so they fit tightly together. Doing so will result in a very tight and solid build.
  1. don't forget to drill the hole in the case for the reset switch, if you have one of the cheaper models without the dip switches, before you reassemble.
  1. the pro micro micro usb connector can be damaged with too much stress, so try to minimize the number of times you plug and unplug it. Also don't do what I did and route the cable in such a way that it stresses the connector. DOH!
