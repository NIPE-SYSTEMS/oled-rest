# oled-rest — OLED controller with REST-API using Gobot

## Build (for Raspberry Pi)

```bash
# GOARM=6 (Raspberry Pi A, A+, B, B+, Zero)
# GOARM=7 (Raspberry Pi 2, 3)
GOOS=linux GOARCH=arm GOARM=6 go build
```

## Deploy

```bash
scp oled-rest root@$SSH_HOST:/root
ssh -t root@$SSH_HOST /root/oled-rest
```

## API

### Display an image

Use `GET /show` or `POST /show` with JSON data in message body containing an monochrome image to display the image. The JSON data is a two-dimensional array of `boolean`s indicating each pixel's color (on/off).

```go
type color bool // true: on, false: off
type row []color // one row consists of 128 colors
type image []row // one image consists of 64 rows
```

### Set brightness

Use `GET /brightness/%d` where `%d` is an integer between `[0-255]` to set the brightness of the display.

## License

MIT
