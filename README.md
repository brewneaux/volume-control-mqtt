# volume-control-mqtt

A simple repo that allows you to set up a Docker container to control Linux volume via MQTT.  

I created this because I recently purchased an [Ikea SYMFONISK
Sound remote](https://www.ikea.com/us/en/p/symfonisk-sound-remote-black-10433847/) that I wanted to use with HomeAssistant to control my HTPC volume, which is also where HomeAssistant runs (in Docker).  Since the rest of my HA infrastructure is Dockerized, I also wanted a way to control the volume from Docker.

This project uses the [eclipse/paho.mqtt.golang](github.com/eclipse/paho.mqtt.golang) to communicate with MQTT, and the [volume-go library by itchyny](github.com/itchyny/volume-go).

# Running it in Docker

This uses PulseAudio over a socket shared into the container

```shell
docker build -t volume-control .
```

```shell
docker run \
  -v "~/.config/pulse:/.config/pulse" \
  -v /run/user/$UID/pulse/native:/run/user/$UID/pulse/native \
  -e VOLUME_INCREMENT=3 \
  -e PULSE_SERVER="unix:/run/user/$UID/pulse/native" \
  -e MQTT_BROKER=localhost \
  -e MQTT_PORT=1883 \
  -e MQTT_TOPIC=/volume-control \
volume-control
```

The VOLUME_INCREMENT is the number it adjusts volume up and down - I found 3 to be the most comfortable with the IKEA control.

The MQTT_BROKER, PORT and TOPIC can all be set to your desire.

# Expected Messages

Right now, I have it hardcoded to just expect an "up" or "down" message.  