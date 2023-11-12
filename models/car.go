package models

import (
	"fmt"
	"math/rand"
	"parking-simulator/helpers"
	"sync"
	"time"

	"github.com/faiface/pixel"
)

type Car struct {
	ParkingTime int // Tiempo en segundos que estará estacionado
	Id int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano())
	parkingTime := rand.Intn(10) + 15 
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch<- *car
		rand.Seed(time.Now().UnixNano()) 
		newTime := rand.Intn(2) + 1
		time.Sleep(time.Second * time.Duration(newTime))
	}
	fmt.Println("Ya acabaron de generarse los 100 autos")
	close(ch)
}

func (c *Car) Timer(pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan string, sprite *pixel.Sprite, chWin chan helpers.ImgCar, coo pixel.Vec, chExit chan helpers.ImgCar) {
	mu.Lock()
	data := helpers.NewImgCar(sprite, pos, true, coo)
	chWin<-*data
	*chEntrance<-"Entrando"
	pc.nSpaces--
	fmt.Println("Un nuevo automovil acaba de estacionarse y quedan ", pc.nSpaces, " espacios disponibles")
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))
	
	mu.Lock()
	data = helpers.NewImgCar(sprite, pos, false, coo)
	chExit<-*data
	pc.nSpaces = pc.nSpaces + 1
	spaces[pos] = true
	fmt.Println("Un automovil acaba de irse del estacionamiento y quedan ", pc.nSpaces, " espacios disponibles")
	mu.Unlock()

	mu.Lock()
	*chEntrance<-"Saliendo"
	mu.Unlock()
}

func (c *Car) Unpark(ch, chExit chan helpers.ImgCar){
	for {
		select {
		case car, _ := <-chExit:
			ch<-car
		}
	}
}