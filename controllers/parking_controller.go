package controllers

import (
	"parking-simulator/helpers"
	"parking-simulator/models"
	"sync"
)

type ParkingController struct {
	model *models.Parking
	mu    *sync.Mutex
	sw	   *helpers.SpaceParkingView
}

func NewParkingController(mu *sync.Mutex) *ParkingController {
	return &ParkingController{
		model: models.NewParking(),
		mu:    mu,
		sw: helpers.NewParkingView(),
	}
}

func (pc *ParkingController) Park(chCar *chan models.Car, entranceController *EntranceController, carController *CarController, chEntrance *chan string, chWin chan helpers.ImgCar, chExit chan helpers.ImgCar) {
	go pc.ChangingState(chEntrance, entranceController)
	for {
		select {
		case car, ok := <-*chCar:
			if !ok {
				return
			}
			pos := pc.model.FindSpaces()
			if pos != -1 {
				coo := pc.sw.GetCoordinates(pos)
				sprite := helpers.LoadPicture("./assets/Carro.png")
				if entranceController.model.GetState() == "Inactivo" || entranceController.model.GetState() == "Entrando" {
					go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo, chExit)
				} else {
					*chEntrance <- "Entrando"
					go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo, chExit)
				}
			}
		}
	}
}

func (pc *ParkingController) ChangingState(chEntrance *chan string, entrancecontroller *EntranceController) {
	for {
		select {
		case change, _ := <-*chEntrance:
			entrancecontroller.model.SetState(change)
		}
	}
}