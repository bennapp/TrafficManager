package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"engo.io/engo"
		)

var Spritesheet *common.Spritesheet

var snakeParts = map[string]int {
	"head": 0,
}

type Snake struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}


type SnakeBuildingSystem struct {
	world *ecs.World
}

// New is the initialisation of the System
func (sb *SnakeBuildingSystem) New(w *ecs.World) {
	sb.world = w
	Spritesheet = common.NewSpritesheetWithBorderFromFile("textures/Snake.png", 16, 16, 1, 1)
	sb.generateSnake()
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (sb *SnakeBuildingSystem) Update(dt float32) {
	//fmt.Println(dt)
}

// Remove is called whenever an Entity is removed from the scene, and thus from this system
func (*SnakeBuildingSystem) Remove(ecs.BasicEntity) {}

// generateCity randomly generates a city in a random location on the map
func (sb *SnakeBuildingSystem) generateSnake() {
	snake := &Snake{BasicEntity: ecs.NewBasic()}

	snake.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    10,
		Height:   10,
	}

	snake.RenderComponent.Drawable = Spritesheet.Cell(snakeParts["head"])

	for _, system := range sb.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&snake.BasicEntity, &snake.RenderComponent, &snake.SpaceComponent)
		}
	}
}
