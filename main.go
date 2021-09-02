package main

import (
	"time"

	"github.com/retrogamecoder/fallingblocks/model"
)

func run() {
	m := model.New()

	lastTime := time.Now()

	for {
		current := time.Now()
		dt := current.Sub(lastTime)

		// MVC: model, view, controller

		// Check for user input
		m.Update(dt)

		// Render the world
		// Check for exit conditions

		lastTime = current
	}
}

func main() {
}
