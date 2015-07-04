package main

import (
	"fmt"
	"github.com/jbott/go-mission-control/missioncontrol"
)

func main() {
	mc := missioncontrol.Init()

	// Add the widgets

	mc.Start()
}
