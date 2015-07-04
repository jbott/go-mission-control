package missioncontrol

type window struct {
	X int
	Y int
	W int
	H int
}

func NewWindow(x int, y int) *window {
	return &window{x, y, 0, 0}
}
