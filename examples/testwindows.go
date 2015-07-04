package main

import (
	"github.com/jbott/go-mission-control/missioncontrol"
)

func main() {
	mc := missioncontrol.Init()

	// Add the widgets
	wind := missioncontrol.NewWindow(10, 10)
	wind.W = 20
	wind.H = 20
	mc.AddWindow(wind)

	mc.Start()
}
