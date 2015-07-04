package main

import (
	"github.com/jbott/go-mission-control/missioncontrol"
)

func main() {
	mc := missioncontrol.Init()

	mc.SetBackgroundColor(20, 20, 20, 255)
	mc.SetDefaultBorderColor(255, 0, 0, 255)

	// Add the widgets
	wind := mc.NewWindow(10, 10)
	wind.W = 200
	wind.H = 400
	mc.Add(wind)

	mc.Start()
}
