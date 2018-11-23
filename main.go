package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type image [][]bool

func main() {
	raspberry := raspi.NewAdaptor()
	raspberry.Connect()
	oled := i2c.NewSSD1306Driver(raspberry)
	oled.Start()
	oled.Clear()
	oled.Display()

	router := mux.NewRouter()
	router.HandleFunc("/show", func(response http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var i image
		err := decoder.Decode(&i)
		if err != nil {
			log.Println("error:", err)
			response.WriteHeader(400)
			response.Write([]byte(err.Error()))
			return
		}
		if len(i) != 64 {
			log.Println("dimension wrong:", len(i), "expected 64")
			response.WriteHeader(400)
			response.Write([]byte("dimension wrong, expected height = 64"))
			return
		}
		for y := 0; y < len(i); y++ {
			if len(i[y]) != 128 {
				log.Println("dimension wrong:", len(i[y]), "expected 128")
				response.WriteHeader(400)
				response.Write([]byte("dimension wrong, expected height = 128"))
				return
			}
		}
		response.WriteHeader(200)
		log.Println("Displaying new image ...")
		oled.Clear()
		for y, w, h := 0, oled.DisplayWidth, oled.DisplayHeight; y < h; y++ {
			for x := 0; x < w; x++ {
				if i[y][x] {
					oled.Set(x, y, 1)
				} else {
					oled.Set(x, y, 0)
				}
			}
		}
		oled.Display()
	})
	router.HandleFunc("/brightness/{brightness:[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]}", func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		brightness, _ := strconv.Atoi(vars["brightness"])
		log.Println("Setting brightness to", brightness, "...")
		oled.SetContrast(byte(brightness))
		response.WriteHeader(200)
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
