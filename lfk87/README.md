# LFK 87 Project

Purchased a used Coolermaster Quick Fire Rapid from Ebay for the
case and the plate so I can install a custom LFK87 PCB.

### Flashing using DFU

There's an issue if you try to put the keyboard into programming
mode using a USB hub, probably because it wants the full 500mA. I
had to connect it directly to my laptop to get it to stop rebooting.

#### Putting it into programming mode.

 1. `FN` (hold) + `Right Shift` (tap)
 1. enter (tap)

 1. `dfu-programmer at90usb646 erase`
 1. `dfu-programmer at90usb646 flash` &lt;`firmware.hex`&gt;
 1. `dfu-programmer at90usb646 reset`

### Backlight Brightness

 1. `FN` (hold) + `Right Shift` (tap)
 1. `-` (decrease) `+` increase

### Images

![powder coated CM QFR Plate](https://i.imgur.com/sHWNMhU.jpg)
![powder coated CM QFR Plate Front](https://i.imgur.com/L7ItfiS.jpg)
![powder coated CM QFR Plate Close up](https://i.imgur.com/9hiMXo1.jpg)

