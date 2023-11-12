package controllers

import (
	"parking-simulator/helpers"
	"parking-simulator/models"
	"sync"
)

type CarController struct {
	model *models.Car
	mu    *sync.Mutex
}

func NewCarController(mu *sync.Mutex) *CarController {
	return &CarController{
		model: models.NewCar(),
		mu:    mu,
	}
}

func (cc *CarController) GenerateCars(n int, chCar *chan models.Car) {
	cc.model.GenerateCars(n, *chCar)
}

func (cc *CarController) Unpark(chWin chan helpers.ImgCar,ch chan helpers.ImgCar){
	go models.NewCar().Unpark(chWin, ch)
}