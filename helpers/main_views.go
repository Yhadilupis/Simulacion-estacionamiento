package helpers

import (
	"image"
	"os"
	"github.com/faiface/pixel"
)

type SpaceParkingView struct {
	spaces [20]pixel.Vec
}

type ImgCar struct {
	sprite *pixel.Sprite
	Id int
	enterating bool
	pos pixel.Vec
}

func NewParkingView() *SpaceParkingView {
	return &SpaceParkingView{
		spaces: [20]pixel.Vec{
			pixel.V(230, 670),
			pixel.V(315, 670),
			pixel.V(400, 670),
			pixel.V(485, 670),
			pixel.V(570, 670),

			pixel.V(230, 520),
			pixel.V(315, 520),
			pixel.V(400, 520),
			pixel.V(485, 520),
			pixel.V(570, 520),

			pixel.V(230, 275),
			pixel.V(315, 275),
			pixel.V(400, 275),
			pixel.V(485, 275),
			pixel.V(570, 275),

			pixel.V(230, 125),
			pixel.V(315, 125),
			pixel.V(400, 125),
			pixel.V(485, 125),
			pixel.V(570, 125),
		},
	}
}

func NewImgCar(sprite *pixel.Sprite, Id int, state bool, pos pixel.Vec) *ImgCar {
	return &ImgCar{
		sprite: sprite,
		Id: Id,
		enterating: state,
		pos: pos,
	}
}

func (ic *ImgCar) GetSprite() *pixel.Sprite {
	return ic.sprite
}

func (ic *ImgCar) GetPos() pixel.Vec {
	return ic.pos
}

func (ic *ImgCar) GetId() int {
	return ic.Id
}

func (ic *ImgCar) GetStatus() bool {
	return ic.enterating
}

func (spw *SpaceParkingView) GetCoordinates(n int) pixel.Vec {
	return spw.spaces[n]
}

func LoadPicture(path string) (*pixel.Sprite) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pic, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	imgPP := pixel.PictureDataFromImage(pic)

	return pixel.NewSprite(imgPP, imgPP.Bounds())
}