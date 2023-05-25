# gojolt

gojolt is a simple program that keeps your screen awake preventing screen blank and automatic suspension.

It uses [dbus](https://github.com/godbus/dbus) to interact with the D-Bus message system.

It works both on Gnome Shell on Xorg and on Wayland.

## Installation

Binaries are provided under the [Releases](https://github.com/vittoriocandolo/gojolt/releases) section.

Most Linux systems should be able to run them.

### Building from source

**golang 1.20+ is required**

Simply clone this repo and run `go build -o gojolt -ldflags="-s -w" main.go`

gojolt is currently based on dbus v5.1.0.

## Usage

It's a CLI program.

You can place it in your PATH (in `/usr/local/bin/` for example) and then run it with a terminal command like this: `gojolt 20`

The amount of time (in minutes) of screen blank inhibition has to be passed as first and only argument.
