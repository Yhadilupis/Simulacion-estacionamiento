package models

type Entrance struct {
	actualState string
}

func NewEntrance () *Entrance {
	return &Entrance{
		actualState: "Inactivo",
	}
}

func (e *Entrance) GetState() string {
	return e.actualState
}

func (e *Entrance) SetState(state string) {
	e.actualState = state
}