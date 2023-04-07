# Wissh

*An experimental tool for diagnosing offline balena devices.*

*Status:* On early stages of design and development. Hardly usable at the
moment. This is an experiment, not an officially supported balena tool.

## Vision

We are still forming the vision for this tool. As of February 2023, it goes more
or less like this:

* A Field Technician Toolkit for [balena](https://www.balena.io) devices.
* Especially for devices that are offline and therefore cannot be troubleshoot
  through balenaCloud.
* Helps to figure out why a device is offline and bring it back online.

## How can I add a new check?

There are two steps:

1. **Implement check.** In summary, a check is anything that implements the
   `wissh.Check` interface. We keep checks under `pkg/checks`, one check per
   file.
2. **Add your check to the list of checks to execute.** Just add it to the
   `checks.All()` function, in `pkg/checks/check_list.go`.

For an example, take a look at `pkg/checks/ping_api.go` which is a simple one.
For a more complex example, check `pkg/checks/static_ip_config.go`: this one
does a good amount of processing on the data obtained via SSH to check for
issues.

## Building

This is currently being developed under Linux. It should be buildable for all
major operating systems, but I am afraid I currently cannot provide instructions
for anything other than Linux. (If you want to try building for other platforms
and contribute with instructions, please be my guest!)

### Linux

Dependencies:

* A recent Go installation; version 1.19 or above should be fine.
* A C compiler (we use [Fyne](https://fyne.io/) for the GUI, which uses cgo).
* Some tools and libraries we depend upon (see below).

If you are running Ubuntu, installing these packages should suffice:

```sh
apt-get install ca-certificates golang git libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl-dev libxxf86vm-dev
```

To actually build, just issue

```sh
go build ./cmd/wissh-gui
```
