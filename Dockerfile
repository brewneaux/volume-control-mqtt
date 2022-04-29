FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o volume_control_mqtt

FROM ubuntu:18.04

RUN apt-get update && apt-get install -y alsa-utils pulseaudio-utils
RUN useradd -ms /bin/bash htpc
USER htpc

COPY --from=builder /app/volume_control_mqtt /app/volume_control_mqtt

CMD ["/app/volume_control_mqtt"]
