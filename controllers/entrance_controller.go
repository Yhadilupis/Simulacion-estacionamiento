package controllers

import (
	"parking-simulator/models"
	"sync"
)

type EntranceController struct {
	model *models.Entrance
	mu    *sync.Mutex
}

func NewEntranceController(mu *sync.Mutex) *EntranceController {
	return &EntranceController{
		model: models.NewEntrance(),
		mu:    mu,
	}
}