package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"image/color"
	"github.com/bennapp/TrafficManager/systems"
)

type Snake struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/Snake.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.

func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	world.AddSystem(new(common.RenderSystem))

	world.AddSystem(&systems.SnakeBuildingSystem{})

	common.SetBackground(color.White)
}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  800,
		Height: 600,
	}
	engo.Run(opts, new(myScene))
}
