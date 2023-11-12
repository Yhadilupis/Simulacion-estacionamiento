package main

import (
	_ "image/png"
	"parking-simulator/controllers"
	"parking-simulator/helpers"
	"parking-simulator/models"
	"sync"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Simulador de estacionamiento",
		Bounds: pixel.R(0, 0, 790, 790),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	carChannel := make(chan models.Car, 100)
	entranceChannel := make(chan string)
	winChannel := make(chan helpers.ImgCar)

	carExit := make(chan helpers.ImgCar)

	mu := &sync.Mutex{}

	parkingController := controllers.NewParkingController(mu)
	entranceController := controllers.NewEntranceController(mu)
	carController := controllers.NewCarController(mu)

	go parkingController.Park(&carChannel, entranceController, carController, &entranceChannel, winChannel, carExit)
	go carController.GenerateCars(100, &carChannel)

	go carController.Unpark(winChannel, carExit)

	spriteParking := helpers.LoadPicture("./assets/Estacionamiento.png")

	var arr []helpers.ImgCar
	for !win.Closed() {
		win.Clear(colornames.Black)

		spriteParking.Draw(win, pixel.IM.Moved(cfg.Bounds.Center()))

		select {
		case val := <-winChannel:

			if val.GetStatus() {
				arr = append(arr, val)
			}else{
				var arrAux []helpers.ImgCar
				for _, value := range arr {
					if value.GetId() != val.GetId() {
						arrAux = append(arrAux, value)
					}
				}
				arr = arr[:0]
				arr = append(arr, arrAux...)
			}
		}

		for _, value := range arr {
			sprite := value.GetSprite()
			pos := value.GetPos()
			sprite.Draw(win, pixel.IM.Moved(pos))
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}