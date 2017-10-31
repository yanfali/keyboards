# Soarer's Converter

@see [Deskthority](https://deskthority.net/workshop-f7/xt-at-ps2-terminal-to-usb-converter-with-nkro-t2510.html) for details about soarer's converter.

## Makefile

default is to run ./scinfo which gives you info about the firmware
running on the converter.

`make` will just run default.

If you want to program the converter, there's a `write` task.  It runs the
scas (assembly tool) against yanfali.sc and creates a binary configuration
file, which is then written to the converter using scwr.

## pre-requisites

  1. You must have built soarer's tools. Right now I have only succeeded on
building them on Linux.
  1. you have root privileges on the Linux box connected to the converter.
