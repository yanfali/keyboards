# Soarer's Converter Tools

This contains a working build for soarer's converter's tools. It just
adds back in the compatiblity headers that didn't exist in 2012.
If soarer would like me to remove these files I will do so, but since
he's no longer actively maintaining the project anymore and I would like
to use these tools to customize my converter I'm adding them here for
easy access.

## Building

```sh
cd build/linux
make
```

## Usage

The commands require access to rawhid which requires root privileges
by default, so run the commands `scinfo`, `scrd`, `scrd` and `scwr`
using sudo. When I ran scinfo it caused the teensy to stop responding
to keyboard commands. Only unplugging it from usb forced it to reboot.
