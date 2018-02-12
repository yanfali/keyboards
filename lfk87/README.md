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

#### Backlight Brightness

 1. `FN` (hold) + `Right Shift` (tap)
 1. `-` (decrease) `+` increase
